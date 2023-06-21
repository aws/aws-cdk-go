package awss3


// All supported inventory frequencies.
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
type InventoryFrequency string

const (
	// A report is generated every day.
	InventoryFrequency_DAILY InventoryFrequency = "DAILY"
	// A report is generated every Sunday (UTC timezone) after the initial report.
	InventoryFrequency_WEEKLY InventoryFrequency = "WEEKLY"
)

