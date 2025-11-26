package previewawsgroundstationmixins


// Provides information about how AWS Ground Station should save downlink data to S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3RecordingConfigProperty := &S3RecordingConfigProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	Prefix: jsii.String("prefix"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-s3recordingconfig.html
//
type CfnConfigPropsMixin_S3RecordingConfigProperty struct {
	// S3 Bucket where the data is written.
	//
	// The name of the S3 Bucket provided must begin with `aws-groundstation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-s3recordingconfig.html#cfn-groundstation-config-s3recordingconfig-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The prefix of the S3 data object.
	//
	// If you choose to use any optional keys for substitution, these values will be replaced with the corresponding information from your contact details. For example, a prefix of `{satellite_id}/{year}/{month}/{day}/` will replaced with `fake_satellite_id/2021/01/10/`
	//
	// *Optional keys for substitution* : `{satellite_id}` | `{config-name}` | `{config-id}` | `{year}` | `{month}` | `{day}`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-s3recordingconfig.html#cfn-groundstation-config-s3recordingconfig-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Defines the ARN of the role assumed for putting archives to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-s3recordingconfig.html#cfn-groundstation-config-s3recordingconfig-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

