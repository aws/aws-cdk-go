package awss3


// All supported inventory list formats.
type InventoryFormat string

const (
	// Generate the inventory list as CSV.
	InventoryFormat_CSV InventoryFormat = "CSV"
	// Generate the inventory list as Parquet.
	InventoryFormat_PARQUET InventoryFormat = "PARQUET"
	// Generate the inventory list as ORC.
	InventoryFormat_ORC InventoryFormat = "ORC"
)

