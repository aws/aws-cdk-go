package previewawsbedrockagentcoremixins


// The recording configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordingConfigProperty := &RecordingConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	S3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-recordingconfig.html
//
type CfnBrowserCustomPropsMixin_RecordingConfigProperty struct {
	// The recording configuration for a browser.
	//
	// This structure defines how browser sessions are recorded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-recordingconfig.html#cfn-bedrockagentcore-browsercustom-recordingconfig-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-recordingconfig.html#cfn-bedrockagentcore-browsercustom-recordingconfig-s3location
	//
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

