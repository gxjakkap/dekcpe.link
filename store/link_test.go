package store

import (
	"strings"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DryRun builds SQL without touching a database; the callback captures it.
func TestGetLinkBySlugExcludesSoftDeleted(t *testing.T) {
	db, err := gorm.Open(postgres.Open("host=localhost"), &gorm.Config{DryRun: true, DisableAutomaticPing: true})
	if err != nil {
		t.Fatal(err)
	}
	var sql string
	db.Callback().Query().After("gorm:query").Register("capture", func(tx *gorm.DB) {
		sql = tx.Statement.SQL.String()
	})

	if _, err := NewLinkStore(db).GetLinkBySlug("abc"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(sql, `"links"."deleted_at" IS NULL`) {
		t.Fatalf("query does not filter soft-deleted links: %s", sql)
	}
}
