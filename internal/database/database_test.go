package database

import (
	"testing"
	"time"
)

type TestModel struct {
	Id        uint
	Test      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TestMustConnect(t *testing.T) {
	MustConnect()
	Close()
}

func TestConnect(t *testing.T) {
	if err := Connect(); err != nil {
		t.Error(err)
	}
	Close()
}

func TestDuplicateConnect(t *testing.T) {
	MustConnect()
	MustConnect()
	Close()
}

func TestDuplicateClose(t *testing.T) {
	Close()
	Close()
}

func TestMigrate(t *testing.T) {
	MustConnect()
	if err := Migrate(TestModel{}); err != nil {
		t.Error(err)
	}
	if err := DB.Migrator().DropTable(TestModel{}); err != nil {
		t.Errorf("테스트 리소스가 제대로 정리되지 않음 %v", err)
	}
	Close()
}

func TestMigrateWithNoConnection(t *testing.T) {
	if err := Migrate(TestModel{}); err == nil {
		t.Error("잘못된 에러 처리")
	}
}
