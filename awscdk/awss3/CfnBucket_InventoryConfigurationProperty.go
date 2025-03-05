package awss3


// Specifies the inventory configuration for an Amazon S3 bucket.
//
// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inventoryConfigurationProperty := &InventoryConfigurationProperty{
//   	Destination: &DestinationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		Format: jsii.String("format"),
//
//   		// the properties below are optional
//   		BucketAccountId: jsii.String("bucketAccountId"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	Id: jsii.String("id"),
//   	IncludedObjectVersions: jsii.String("includedObjectVersions"),
//   	ScheduleFrequency: jsii.String("scheduleFrequency"),
//
//   	// the properties below are optional
//   	OptionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html
//
type CfnBucket_InventoryConfigurationProperty struct {
	// Contains information about where to publish the inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// If set to `True` , an inventory list is generated. If set to `False` , no inventory list is generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// If set to `All` , the list includes all the object versions, which adds the version-related fields `VersionId` , `IsLatest` , and `DeleteMarker` to the list. If set to `Current` , the list does not contain these version-related fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-includedobjectversions
	//
	IncludedObjectVersions *string `field:"required" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Specifies the schedule for generating inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-schedulefrequency
	//
	ScheduleFrequency *string `field:"required" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Contains the optional fields that are included in the inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-optionalfields
	//
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// Specifies the inventory filter prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-inventoryconfiguration.html#cfn-s3-bucket-inventoryconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

