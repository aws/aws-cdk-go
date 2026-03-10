package previewawsbedrockagentcoremixins


// The configuration that specifies where evaluation results should be written.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputConfigProperty := &OutputConfigProperty{
//   	CloudWatchConfig: &CloudWatchOutputConfigProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-outputconfig.html
//
type CfnOnlineEvaluationConfigPropsMixin_OutputConfigProperty struct {
	// The CloudWatch configuration for writing evaluation results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-outputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-outputconfig-cloudwatchconfig
	//
	CloudWatchConfig interface{} `field:"optional" json:"cloudWatchConfig" yaml:"cloudWatchConfig"`
}

