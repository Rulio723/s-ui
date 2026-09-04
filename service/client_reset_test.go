package service

import (
	"encoding/json"
	"testing"

	"github.com/alireza0/s-ui/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestResetClientsRenewsExpiryWhenEnabled(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:reset-renew?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Client{}, &model.Changes{}); err != nil {
		t.Fatal(err)
	}

	const now int64 = 1_800_000_000
	client := model.Client{
		Name:      "renewing-client",
		Enable:    false,
		Inbounds:  json.RawMessage(`[7]`),
		Links:     json.RawMessage(`[]`),
		Config:    json.RawMessage(`{"vless":{"name":"renewing-client","uuid":"00000000-0000-4000-8000-000000000001"}}`),
		Expiry:    now - 60,
		Up:        100,
		Down:      200,
		AutoReset: true,
		AutoRenew: true,
		ResetDays: 30,
		NextReset: now - 1,
	}
	if err := db.Create(&client).Error; err != nil {
		t.Fatal(err)
	}

	ids, err := (&ClientService{}).ResetClients(db, now)
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 1 || ids[0] != 7 {
		t.Fatalf("unexpected changed inbound IDs: %v", ids)
	}
	if err := db.First(&client, client.Id).Error; err != nil {
		t.Fatal(err)
	}
	wantReset := now + 30*86400
	if !client.Enable || client.NextReset != wantReset || client.Expiry != wantReset {
		t.Fatalf("renewal state mismatch: enable=%v next=%d expiry=%d", client.Enable, client.NextReset, client.Expiry)
	}
	if client.Up != 0 || client.Down != 0 || client.TotalUp != 100 || client.TotalDown != 200 {
		t.Fatalf("traffic reset mismatch: up=%d down=%d totalUp=%d totalDown=%d", client.Up, client.Down, client.TotalUp, client.TotalDown)
	}
}

func TestResetClientsLeavesExpiryWhenAutoRenewDisabled(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:reset-no-renew?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Client{}, &model.Changes{}); err != nil {
		t.Fatal(err)
	}

	const now int64 = 1_800_000_000
	const expiry int64 = 1_900_000_000
	client := model.Client{
		Name:      "reset-only-client",
		Enable:    true,
		Inbounds:  json.RawMessage(`[]`),
		Links:     json.RawMessage(`[]`),
		Config:    json.RawMessage(`{}`),
		Expiry:    expiry,
		AutoReset: true,
		ResetDays: 30,
		NextReset: now - 1,
	}
	if err := db.Create(&client).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := (&ClientService{}).ResetClients(db, now); err != nil {
		t.Fatal(err)
	}
	if err := db.First(&client, client.Id).Error; err != nil {
		t.Fatal(err)
	}
	if client.Expiry != expiry {
		t.Fatalf("expiry changed without auto-renew: got %d want %d", client.Expiry, expiry)
	}
}
