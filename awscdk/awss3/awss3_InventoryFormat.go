package awss3


// All supported inventory list formats.
// Experimental.
type InventoryFormat string

const (
	// Generate the inventory list as CSV.
	// Experimental.
	InventoryFormat_CSV InventoryFormat = "CSV"
	// Generate the inventory list as Parquet.
	// Experimental.
	InventoryFormat_PARQUET InventoryFormat = "PARQUET"
	// Generate the inventory list as ORC.
	// Experimental.
	InventoryFormat_ORC InventoryFormat = "ORC"
)

