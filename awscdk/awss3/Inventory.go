package awss3


// Specifies the inventory configuration of an S3 Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   inventory := &Inventory{
//   	Destination: &InventoryDestination{
//   		Bucket: bucket,
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   	Format: awscdk.Aws_s3.InventoryFormat_CSV,
//   	Frequency: awscdk.*Aws_s3.InventoryFrequency_DAILY,
//   	IncludeObjectVersions: awscdk.*Aws_s3.InventoryObjectVersion_ALL,
//   	InventoryId: jsii.String("inventoryId"),
//   	ObjectsPrefix: jsii.String("objectsPrefix"),
//   	OptionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
//
type Inventory struct {
	// The destination of the inventory.
	Destination *InventoryDestination `field:"required" json:"destination" yaml:"destination"`
	// Whether the inventory is enabled or disabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The format of the inventory.
	// Default: InventoryFormat.CSV
	//
	Format InventoryFormat `field:"optional" json:"format" yaml:"format"`
	// Frequency at which the inventory should be generated.
	// Default: InventoryFrequency.WEEKLY
	//
	Frequency InventoryFrequency `field:"optional" json:"frequency" yaml:"frequency"`
	// If the inventory should contain all the object versions or only the current one.
	// Default: InventoryObjectVersion.ALL
	//
	IncludeObjectVersions InventoryObjectVersion `field:"optional" json:"includeObjectVersions" yaml:"includeObjectVersions"`
	// The inventory configuration ID.
	//
	// Should be limited to 64 characters and can only contain letters, numbers, periods, dashes, and underscores.
	// Default: - generated ID.
	//
	InventoryId *string `field:"optional" json:"inventoryId" yaml:"inventoryId"`
	// The inventory will only include objects that meet the prefix filter criteria.
	// Default: - No objects prefix.
	//
	ObjectsPrefix *string `field:"optional" json:"objectsPrefix" yaml:"objectsPrefix"`
	// A list of optional fields to be included in the inventory result.
	// Default: - No optional fields.
	//
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
}

