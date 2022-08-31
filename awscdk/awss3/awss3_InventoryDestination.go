package awss3


// The destination of the inventory.
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
type InventoryDestination struct {
	// Bucket where all inventories will be saved in.
	// Experimental.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The account ID that owns the destination S3 bucket.
	//
	// If no account ID is provided, the owner is not validated before exporting data.
	// It's recommended to set an account ID to prevent problems if the destination bucket ownership changes.
	// Experimental.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The prefix to be used when saving the inventory.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

