package awss3


// Inventory version support.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))
//
//   dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &bucketProps{
//   	inventories: []inventory{
//   		&inventory{
//   			frequency: s3.inventoryFrequency_DAILY,
//   			includeObjectVersions: s3.inventoryObjectVersion_CURRENT,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   			},
//   		},
//   		&inventory{
//   			frequency: s3.*inventoryFrequency_WEEKLY,
//   			includeObjectVersions: s3.*inventoryObjectVersion_ALL,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   				prefix: jsii.String("with-all-versions"),
//   			},
//   		},
//   	},
//   })
//
type InventoryObjectVersion string

const (
	// Includes all versions of each object in the report.
	InventoryObjectVersion_ALL InventoryObjectVersion = "ALL"
	// Includes only the current version of each object in the report.
	InventoryObjectVersion_CURRENT InventoryObjectVersion = "CURRENT"
)

