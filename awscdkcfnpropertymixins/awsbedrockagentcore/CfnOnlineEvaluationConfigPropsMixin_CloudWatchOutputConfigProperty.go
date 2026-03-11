package awsbedrockagentcore


// The CloudWatch configuration for writing evaluation results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cloudWatchOutputConfigProperty := &CloudWatchOutputConfigProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig.html
//
type CfnOnlineEvaluationConfigPropsMixin_CloudWatchOutputConfigProperty struct {
	// The CloudWatch log group name for evaluation results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

