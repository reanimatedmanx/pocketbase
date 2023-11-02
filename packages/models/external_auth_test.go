package models_test

import (
	"testing"

	"github.com/pocketbase/pocketbase/packages/models"
)

func TestExternalAuthTableName(t *testing.T) {
	m := models.ExternalAuth{}
	if m.TableName() != "_externalAuths" {
		t.Fatalf("Unexpected table name, got %q", m.TableName())
	}
}
