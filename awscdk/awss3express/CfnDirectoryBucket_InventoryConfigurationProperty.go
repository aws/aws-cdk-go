package awss3express


// Specifies an inventory configuration for an Amazon S3 Express bucket.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html
//
type CfnDirectoryBucket_InventoryConfigurationProperty struct {
	// Specifies information about where to publish inventory reports for an Amazon S3 Express bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-includedobjectversions
	//
	IncludedObjectVersions *string `field:"required" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Specifies the schedule for generating inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-schedulefrequency
	//
	ScheduleFrequency *string `field:"required" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Contains the optional fields that are included in the inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-optionalfields
	//
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// The prefix that is prepended to all inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

