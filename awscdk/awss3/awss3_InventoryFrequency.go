package awss3


// All supported inventory frequencies.
//
// Example:
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
// Experimental.
type InventoryFrequency string

const (
	// A report is generated every day.
	// Experimental.
	InventoryFrequency_DAILY InventoryFrequency = "DAILY"
	// A report is generated every Sunday (UTC timezone) after the initial report.
	// Experimental.
	InventoryFrequency_WEEKLY InventoryFrequency = "WEEKLY"
)

