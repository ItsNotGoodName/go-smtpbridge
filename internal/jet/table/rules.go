//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Rules = newRulesTable("", "rules", "")

type rulesTable struct {
	sqlite.Table

	// Columns
	ID         sqlite.ColumnInteger
	Internal   sqlite.ColumnBool
	InternalID sqlite.ColumnString
	Name       sqlite.ColumnString
	Expression sqlite.ColumnString
	Enable     sqlite.ColumnBool
	UpdatedAt  sqlite.ColumnTimestamp
	CreatedAt  sqlite.ColumnTimestamp

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type RulesTable struct {
	rulesTable

	EXCLUDED rulesTable
}

// AS creates new RulesTable with assigned alias
func (a RulesTable) AS(alias string) *RulesTable {
	return newRulesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RulesTable with assigned schema name
func (a RulesTable) FromSchema(schemaName string) *RulesTable {
	return newRulesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RulesTable with assigned table prefix
func (a RulesTable) WithPrefix(prefix string) *RulesTable {
	return newRulesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RulesTable with assigned table suffix
func (a RulesTable) WithSuffix(suffix string) *RulesTable {
	return newRulesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRulesTable(schemaName, tableName, alias string) *RulesTable {
	return &RulesTable{
		rulesTable: newRulesTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newRulesTableImpl("", "excluded", ""),
	}
}

func newRulesTableImpl(schemaName, tableName, alias string) rulesTable {
	var (
		IDColumn         = sqlite.IntegerColumn("id")
		InternalColumn   = sqlite.BoolColumn("internal")
		InternalIDColumn = sqlite.StringColumn("internal_id")
		NameColumn       = sqlite.StringColumn("name")
		ExpressionColumn = sqlite.StringColumn("expression")
		EnableColumn     = sqlite.BoolColumn("enable")
		UpdatedAtColumn  = sqlite.TimestampColumn("updated_at")
		CreatedAtColumn  = sqlite.TimestampColumn("created_at")
		allColumns       = sqlite.ColumnList{IDColumn, InternalColumn, InternalIDColumn, NameColumn, ExpressionColumn, EnableColumn, UpdatedAtColumn, CreatedAtColumn}
		mutableColumns   = sqlite.ColumnList{InternalColumn, InternalIDColumn, NameColumn, ExpressionColumn, EnableColumn, UpdatedAtColumn, CreatedAtColumn}
	)

	return rulesTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		Internal:   InternalColumn,
		InternalID: InternalIDColumn,
		Name:       NameColumn,
		Expression: ExpressionColumn,
		Enable:     EnableColumn,
		UpdatedAt:  UpdatedAtColumn,
		CreatedAt:  CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}