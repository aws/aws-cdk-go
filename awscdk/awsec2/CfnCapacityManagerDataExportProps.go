package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCapacityManagerDataExport`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityManagerDataExportProps := &CfnCapacityManagerDataExportProps{
//   	OutputFormat: jsii.String("outputFormat"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   	Schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	S3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacitymanagerdataexport.html
//
type CfnCapacityManagerDataExportProps struct {
	// The file format of the exported data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacitymanagerdataexport.html#cfn-ec2-capacitymanagerdataexport-outputformat
	//
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The name of the S3 bucket where export files are delivered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacitymanagerdataexport.html#cfn-ec2-capacitymanagerdataexport-s3bucketname
	//
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The frequency at which data exports are generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacitymanagerdataexport.html#cfn-ec2-capacitymanagerdataexport-schedule
	//
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// The S3 key prefix used for organizing export files within the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacitymanagerdataexport.html#cfn-ec2-capacitymanagerdataexport-s3bucketprefix
	//
	S3BucketPrefix *string `field:"optional" json:"s3BucketPrefix" yaml:"s3BucketPrefix"`
	// The tags associated with the data export configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacitymanagerdataexport.html#cfn-ec2-capacitymanagerdataexport-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

