package awss3


// Inventory version support.
//
// Example:
//   inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))
//
//   dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &BucketProps{
//   	Inventories: []inventory{
//   		&inventory{
//   			Frequency: s3.InventoryFrequency_DAILY,
//   			IncludeObjectVersions: s3.InventoryObjectVersion_CURRENT,
//   			Destination: &InventoryDestination{
//   				Bucket: inventoryBucket,
//   			},
//   		},
//   		&inventory{
//   			Frequency: s3.InventoryFrequency_WEEKLY,
//   			IncludeObjectVersions: s3.InventoryObjectVersion_ALL,
//   			Destination: &InventoryDestination{
//   				Bucket: inventoryBucket,
//   				Prefix: jsii.String("with-all-versions"),
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

