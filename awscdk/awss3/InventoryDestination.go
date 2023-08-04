package awss3


// The destination of the inventory.
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
type InventoryDestination struct {
	// Bucket where all inventories will be saved in.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The account ID that owns the destination S3 bucket.
	//
	// If no account ID is provided, the owner is not validated before exporting data.
	// It's recommended to set an account ID to prevent problems if the destination bucket ownership changes.
	// Default: - No account ID.
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The prefix to be used when saving the inventory.
	// Default: - No prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

