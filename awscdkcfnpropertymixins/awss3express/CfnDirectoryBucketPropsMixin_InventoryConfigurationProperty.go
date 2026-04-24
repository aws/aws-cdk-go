package awss3express


// Specifies an inventory configuration for an Amazon S3 Express bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html
//
type CfnDirectoryBucketPropsMixin_InventoryConfigurationProperty struct {
	// Specifies information about where to publish inventory reports for an Amazon S3 Express bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-includedobjectversions
	//
	IncludedObjectVersions *string `field:"optional" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Contains the optional fields that are included in the inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-optionalfields
	//
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// The prefix that is prepended to all inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies the schedule for generating inventory results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-inventoryconfiguration.html#cfn-s3express-directorybucket-inventoryconfiguration-schedulefrequency
	//
	ScheduleFrequency *string `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
}

