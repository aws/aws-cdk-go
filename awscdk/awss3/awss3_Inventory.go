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
//   inventory := &inventory{
//   	destination: &inventoryDestination{
//   		bucket: bucket,
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	format: awscdk.Aws_s3.inventoryFormat_CSV,
//   	frequency: awscdk.*Aws_s3.inventoryFrequency_DAILY,
//   	includeObjectVersions: awscdk.*Aws_s3.inventoryObjectVersion_ALL,
//   	inventoryId: jsii.String("inventoryId"),
//   	objectsPrefix: jsii.String("objectsPrefix"),
//   	optionalFields: []*string{
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
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The format of the inventory.
	Format InventoryFormat `field:"optional" json:"format" yaml:"format"`
	// Frequency at which the inventory should be generated.
	Frequency InventoryFrequency `field:"optional" json:"frequency" yaml:"frequency"`
	// If the inventory should contain all the object versions or only the current one.
	IncludeObjectVersions InventoryObjectVersion `field:"optional" json:"includeObjectVersions" yaml:"includeObjectVersions"`
	// The inventory configuration ID.
	InventoryId *string `field:"optional" json:"inventoryId" yaml:"inventoryId"`
	// The inventory will only include objects that meet the prefix filter criteria.
	ObjectsPrefix *string `field:"optional" json:"objectsPrefix" yaml:"objectsPrefix"`
	// A list of optional fields to be included in the inventory result.
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
}

