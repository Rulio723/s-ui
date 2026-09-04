package service

import (
	"encoding/json"
	"testing"

	"github.com/alireza0/s-ui/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSaveCalibratesClientTraffic(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:client-traffic?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Client{}); err != nil {
		t.Fatal(err)
	}

	client := model.Client{Name: "traffic-client", Up: 100, Down: 200, TotalUp: 300, TotalDown: 400}
	if err := db.Create(&client).Error; err != nil {
		t.Fatal(err)
	}

	payload, err := json.Marshal(map[string]interface{}{"id": client.Id, "up": int64(5_000), "down": int64(7_000)})
	if err != nil {
		t.Fatal(err)
	}
	ids, err := (&ClientService{}).Save(db, "traffic", payload, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != 0 {
		t.Fatalf("traffic calibration unexpectedly changed inbounds: %v", ids)
	}
	if err := db.First(&client, client.Id).Error; err != nil {
		t.Fatal(err)
	}
	if client.Up != 5_000 || client.Down != 7_000 {
		t.Fatalf("traffic mismatch: up=%d down=%d", client.Up, client.Down)
	}
	if client.TotalUp != 300 || client.TotalDown != 400 {
		t.Fatalf("historical traffic changed: totalUp=%d totalDown=%d", client.TotalUp, client.TotalDown)
	}
}

func TestSaveRejectsInvalidClientTraffic(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:invalid-client-traffic?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Client{}); err != nil {
		t.Fatal(err)
	}

	for name, payload := range map[string]json.RawMessage{
		"missing client": json.RawMessage(`{"id":99,"up":1,"down":2}`),
		"negative value": json.RawMessage(`{"id":1,"up":-1,"down":2}`),
		"missing id":     json.RawMessage(`{"up":1,"down":2}`),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := (&ClientService{}).Save(db, "traffic", payload, ""); err == nil {
				t.Fatal("expected traffic calibration to fail")
			}
		})
	}
}
