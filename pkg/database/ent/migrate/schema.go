// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AlertsColumns holds the columns for the "alerts" table.
	AlertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "scenario", Type: field.TypeString},
		{Name: "bucket_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "events_count", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "stopped_at", Type: field.TypeTime, Nullable: true},
		{Name: "source_ip", Type: field.TypeString, Nullable: true},
		{Name: "source_range", Type: field.TypeString, Nullable: true},
		{Name: "source_as_number", Type: field.TypeString, Nullable: true},
		{Name: "source_as_name", Type: field.TypeString, Nullable: true},
		{Name: "source_country", Type: field.TypeString, Nullable: true},
		{Name: "source_latitude", Type: field.TypeFloat32, Nullable: true},
		{Name: "source_longitude", Type: field.TypeFloat32, Nullable: true},
		{Name: "source_scope", Type: field.TypeString, Nullable: true},
		{Name: "source_value", Type: field.TypeString, Nullable: true},
		{Name: "capacity", Type: field.TypeInt32, Nullable: true},
		{Name: "leak_speed", Type: field.TypeString, Nullable: true},
		{Name: "scenario_version", Type: field.TypeString, Nullable: true},
		{Name: "scenario_hash", Type: field.TypeString, Nullable: true},
		{Name: "simulated", Type: field.TypeBool, Default: false},
		{Name: "uuid", Type: field.TypeString, Nullable: true},
		{Name: "remediation", Type: field.TypeBool, Nullable: true},
		{Name: "machine_alerts", Type: field.TypeInt, Nullable: true},
	}
	// AlertsTable holds the schema information for the "alerts" table.
	AlertsTable = &schema.Table{
		Name:       "alerts",
		Columns:    AlertsColumns,
		PrimaryKey: []*schema.Column{AlertsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "alerts_machines_alerts",
				Columns:    []*schema.Column{AlertsColumns[25]},
				RefColumns: []*schema.Column{MachinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "alert_id",
				Unique:  false,
				Columns: []*schema.Column{AlertsColumns[0]},
			},
		},
	}
	// AllowListsColumns holds the columns for the "allow_lists" table.
	AllowListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "from_console", Type: field.TypeBool},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "allowlist_id", Type: field.TypeString, Nullable: true},
	}
	// AllowListsTable holds the schema information for the "allow_lists" table.
	AllowListsTable = &schema.Table{
		Name:       "allow_lists",
		Columns:    AllowListsColumns,
		PrimaryKey: []*schema.Column{AllowListsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "allowlist_id",
				Unique:  true,
				Columns: []*schema.Column{AllowListsColumns[0]},
			},
			{
				Name:    "allowlist_name",
				Unique:  true,
				Columns: []*schema.Column{AllowListsColumns[3]},
			},
		},
	}
	// AllowListItemsColumns holds the columns for the "allow_list_items" table.
	AllowListItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "comment", Type: field.TypeString, Nullable: true},
		{Name: "value", Type: field.TypeString},
		{Name: "start_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "end_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "start_suffix", Type: field.TypeInt64, Nullable: true},
		{Name: "end_suffix", Type: field.TypeInt64, Nullable: true},
		{Name: "ip_size", Type: field.TypeInt64, Nullable: true},
	}
	// AllowListItemsTable holds the schema information for the "allow_list_items" table.
	AllowListItemsTable = &schema.Table{
		Name:       "allow_list_items",
		Columns:    AllowListItemsColumns,
		PrimaryKey: []*schema.Column{AllowListItemsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "allowlistitem_id",
				Unique:  false,
				Columns: []*schema.Column{AllowListItemsColumns[0]},
			},
			{
				Name:    "allowlistitem_start_ip_end_ip",
				Unique:  false,
				Columns: []*schema.Column{AllowListItemsColumns[6], AllowListItemsColumns[7]},
			},
		},
	}
	// BouncersColumns holds the columns for the "bouncers" table.
	BouncersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "api_key", Type: field.TypeString},
		{Name: "revoked", Type: field.TypeBool},
		{Name: "ip_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "last_pull", Type: field.TypeTime, Nullable: true},
		{Name: "auth_type", Type: field.TypeString, Default: "api-key"},
		{Name: "osname", Type: field.TypeString, Nullable: true},
		{Name: "osversion", Type: field.TypeString, Nullable: true},
		{Name: "featureflags", Type: field.TypeString, Nullable: true},
		{Name: "auto_created", Type: field.TypeBool, Default: false},
	}
	// BouncersTable holds the schema information for the "bouncers" table.
	BouncersTable = &schema.Table{
		Name:       "bouncers",
		Columns:    BouncersColumns,
		PrimaryKey: []*schema.Column{BouncersColumns[0]},
	}
	// ConfigItemsColumns holds the columns for the "config_items" table.
	ConfigItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "value", Type: field.TypeString, SchemaType: map[string]string{"mysql": "longtext", "postgres": "text"}},
	}
	// ConfigItemsTable holds the schema information for the "config_items" table.
	ConfigItemsTable = &schema.Table{
		Name:       "config_items",
		Columns:    ConfigItemsColumns,
		PrimaryKey: []*schema.Column{ConfigItemsColumns[0]},
	}
	// DecisionsColumns holds the columns for the "decisions" table.
	DecisionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "until", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "scenario", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "start_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "end_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "start_suffix", Type: field.TypeInt64, Nullable: true},
		{Name: "end_suffix", Type: field.TypeInt64, Nullable: true},
		{Name: "ip_size", Type: field.TypeInt64, Nullable: true},
		{Name: "scope", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "simulated", Type: field.TypeBool, Default: false},
		{Name: "uuid", Type: field.TypeString, Nullable: true},
		{Name: "alert_decisions", Type: field.TypeInt, Nullable: true},
	}
	// DecisionsTable holds the schema information for the "decisions" table.
	DecisionsTable = &schema.Table{
		Name:       "decisions",
		Columns:    DecisionsColumns,
		PrimaryKey: []*schema.Column{DecisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "decisions_alerts_decisions",
				Columns:    []*schema.Column{DecisionsColumns[16]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "decision_start_ip_end_ip",
				Unique:  false,
				Columns: []*schema.Column{DecisionsColumns[6], DecisionsColumns[7]},
			},
			{
				Name:    "decision_value",
				Unique:  false,
				Columns: []*schema.Column{DecisionsColumns[12]},
			},
			{
				Name:    "decision_until",
				Unique:  false,
				Columns: []*schema.Column{DecisionsColumns[3]},
			},
			{
				Name:    "decision_alert_decisions",
				Unique:  false,
				Columns: []*schema.Column{DecisionsColumns[16]},
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "time", Type: field.TypeTime},
		{Name: "serialized", Type: field.TypeString, Size: 8191},
		{Name: "alert_events", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_alerts_events",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "event_alert_events",
				Unique:  false,
				Columns: []*schema.Column{EventsColumns[5]},
			},
		},
	}
	// LocksColumns holds the columns for the "locks" table.
	LocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// LocksTable holds the schema information for the "locks" table.
	LocksTable = &schema.Table{
		Name:       "locks",
		Columns:    LocksColumns,
		PrimaryKey: []*schema.Column{LocksColumns[0]},
	}
	// MachinesColumns holds the columns for the "machines" table.
	MachinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_push", Type: field.TypeTime, Nullable: true},
		{Name: "last_heartbeat", Type: field.TypeTime, Nullable: true},
		{Name: "machine_id", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "scenarios", Type: field.TypeString, Nullable: true, Size: 100000},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "is_validated", Type: field.TypeBool, Default: false},
		{Name: "auth_type", Type: field.TypeString, Default: "password"},
		{Name: "osname", Type: field.TypeString, Nullable: true},
		{Name: "osversion", Type: field.TypeString, Nullable: true},
		{Name: "featureflags", Type: field.TypeString, Nullable: true},
		{Name: "hubstate", Type: field.TypeJSON, Nullable: true},
		{Name: "datasources", Type: field.TypeJSON, Nullable: true},
	}
	// MachinesTable holds the schema information for the "machines" table.
	MachinesTable = &schema.Table{
		Name:       "machines",
		Columns:    MachinesColumns,
		PrimaryKey: []*schema.Column{MachinesColumns[0]},
	}
	// MetaColumns holds the columns for the "meta" table.
	MetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString, Size: 4095},
		{Name: "alert_metas", Type: field.TypeInt, Nullable: true},
	}
	// MetaTable holds the schema information for the "meta" table.
	MetaTable = &schema.Table{
		Name:       "meta",
		Columns:    MetaColumns,
		PrimaryKey: []*schema.Column{MetaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "meta_alerts_metas",
				Columns:    []*schema.Column{MetaColumns[5]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "meta_alert_metas",
				Unique:  false,
				Columns: []*schema.Column{MetaColumns[5]},
			},
		},
	}
	// MetricsColumns holds the columns for the "metrics" table.
	MetricsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "generated_type", Type: field.TypeEnum, Enums: []string{"LP", "RC"}},
		{Name: "generated_by", Type: field.TypeString},
		{Name: "received_at", Type: field.TypeTime},
		{Name: "pushed_at", Type: field.TypeTime, Nullable: true},
		{Name: "payload", Type: field.TypeString, Size: 2147483647},
	}
	// MetricsTable holds the schema information for the "metrics" table.
	MetricsTable = &schema.Table{
		Name:       "metrics",
		Columns:    MetricsColumns,
		PrimaryKey: []*schema.Column{MetricsColumns[0]},
	}
	// AllowListAllowlistItemsColumns holds the columns for the "allow_list_allowlist_items" table.
	AllowListAllowlistItemsColumns = []*schema.Column{
		{Name: "allow_list_id", Type: field.TypeInt},
		{Name: "allow_list_item_id", Type: field.TypeInt},
	}
	// AllowListAllowlistItemsTable holds the schema information for the "allow_list_allowlist_items" table.
	AllowListAllowlistItemsTable = &schema.Table{
		Name:       "allow_list_allowlist_items",
		Columns:    AllowListAllowlistItemsColumns,
		PrimaryKey: []*schema.Column{AllowListAllowlistItemsColumns[0], AllowListAllowlistItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "allow_list_allowlist_items_allow_list_id",
				Columns:    []*schema.Column{AllowListAllowlistItemsColumns[0]},
				RefColumns: []*schema.Column{AllowListsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "allow_list_allowlist_items_allow_list_item_id",
				Columns:    []*schema.Column{AllowListAllowlistItemsColumns[1]},
				RefColumns: []*schema.Column{AllowListItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlertsTable,
		AllowListsTable,
		AllowListItemsTable,
		BouncersTable,
		ConfigItemsTable,
		DecisionsTable,
		EventsTable,
		LocksTable,
		MachinesTable,
		MetaTable,
		MetricsTable,
		AllowListAllowlistItemsTable,
	}
)

func init() {
	AlertsTable.ForeignKeys[0].RefTable = MachinesTable
	DecisionsTable.ForeignKeys[0].RefTable = AlertsTable
	EventsTable.ForeignKeys[0].RefTable = AlertsTable
	MetaTable.ForeignKeys[0].RefTable = AlertsTable
	AllowListAllowlistItemsTable.ForeignKeys[0].RefTable = AllowListsTable
	AllowListAllowlistItemsTable.ForeignKeys[1].RefTable = AllowListItemsTable
}
