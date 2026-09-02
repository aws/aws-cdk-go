package awsiot


// Configuration for pre-signed S3 URLs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   presignedUrlConfigProperty := &PresignedUrlConfigProperty{
//   	ExpiresInSec: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-presignedurlconfig.html
//
type CfnJobPropsMixin_PresignedUrlConfigProperty struct {
	// How long (in seconds) pre-signed URLs are valid.
	//
	// Valid values are 60 - 3600, the default value is 3600 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-presignedurlconfig.html#cfn-iot-job-presignedurlconfig-expiresinsec
	//
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
	// The ARN of an IAM role that grants permission to download files from the S3 bucket where the job data/updates are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-job-presignedurlconfig.html#cfn-iot-job-presignedurlconfig-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

