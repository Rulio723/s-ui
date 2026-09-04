package sub

import (
	"encoding/json"
	"testing"

	"github.com/alireza0/s-ui/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetClientBySubscriptionID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:subscription-id?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Client{}); err != nil {
		t.Fatal(err)
	}

	clients := []model.Client{
		{Name: "legacy", Enable: true, Config: json.RawMessage(`{}`), Inbounds: json.RawMessage(`[]`), Links: json.RawMessage(`[]`)},
		{Name: "customer@example.com", Email: "customer@example.com", SubID: "private-token", Enable: true, Config: json.RawMessage(`{}`), Inbounds: json.RawMessage(`[]`), Links: json.RawMessage(`[]`)},
	}
	if err := db.Create(&clients).Error; err != nil {
		t.Fatal(err)
	}

	client, err := getClientBySubscriptionID(db, "private-token")
	if err != nil || client.Name != "customer@example.com" {
		t.Fatalf("sub ID lookup failed: client=%v err=%v", client, err)
	}
	client, err = getClientBySubscriptionID(db, "legacy")
	if err != nil || client.Name != "legacy" {
		t.Fatalf("legacy name lookup failed: client=%v err=%v", client, err)
	}
}
