package previewawscodepipelinemixins


// The retry configuration specifies automatic retry for a failed stage, along with the configured retry mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retryConfigurationProperty := &RetryConfigurationProperty{
//   	RetryMode: jsii.String("retryMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-retryconfiguration.html
//
type CfnPipelinePropsMixin_RetryConfigurationProperty struct {
	// The method that you want to configure for automatic stage retry on stage failure.
	//
	// You can specify to retry only failed action in the stage or all actions in the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-retryconfiguration.html#cfn-codepipeline-pipeline-retryconfiguration-retrymode
	//
	RetryMode *string `field:"optional" json:"retryMode" yaml:"retryMode"`
}

