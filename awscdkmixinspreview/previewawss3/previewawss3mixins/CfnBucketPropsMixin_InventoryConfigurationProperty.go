package previewawss3mixins


// Specifies the S3 Inventory configuration for an Amazon S3 bucket.
//
// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inventoryConfigurationProperty := &InventoryConfigurationProperty{
//   	Destination: &DestinationProperty{
//   		BucketAccountId: jsii.String("bucketAccountId"),
//   		BucketArn: jsii.String("bucketArn"),
//   		Format: jsii.String("format"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	Id: jsii.String("id"),
//   	IncludedObjectVersions: jsii.String("includedObjectVersions"),
//   	OptionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   	Prefix: jsii.String("prefix"),
//   	ScheduleFrequency: jsii.String("scheduleFrequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html
//
type CfnBucketPropsMixin_InventoryConfigurationProperty struct {
	// Contains information about where to publish the inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// If set to `True` , an inventory list is generated. If set to `False` , no inventory list is generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// If set to `All` , the list includes all the object versions, which adds the version-related fields `VersionId` , `IsLatest` , and `DeleteMarker` to the list. If set to `Current` , the list does not contain these version-related fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-includedobjectversions
	//
	IncludedObjectVersions *string `field:"optional" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Contains the optional fields that are included in the inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-optionalfields
	//
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// Specifies the inventory filter prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies the schedule for generating inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-schedulefrequency
	//
	ScheduleFrequency *string `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
}

