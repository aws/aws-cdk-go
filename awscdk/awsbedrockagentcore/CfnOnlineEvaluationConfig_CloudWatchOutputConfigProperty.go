package awsbedrockagentcore


// The CloudWatch configuration for writing evaluation results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchOutputConfigProperty := &CloudWatchOutputConfigProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   	MetricsNamespace: jsii.String("metricsNamespace"),
//   	ResultDestination: jsii.String("resultDestination"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig.html
//
type CfnOnlineEvaluationConfig_CloudWatchOutputConfigProperty struct {
	// The CloudWatch log group name for evaluation results.
	//
	// Omit to use the service-managed default log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The CloudWatch metrics namespace for evaluation result metrics.
	//
	// Omit to use the service-managed default namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig-metricsnamespace
	//
	MetricsNamespace *string `field:"optional" json:"metricsNamespace" yaml:"metricsNamespace"`
	// Where evaluation results are written.
	//
	// DEDICATED_LOG_GROUP, the default when omitted, writes to a dedicated result log group. SOURCE_LOG_GROUP writes results back to the trace source log group; LogGroupName must not be specified with SOURCE_LOG_GROUP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchoutputconfig-resultdestination
	//
	ResultDestination *string `field:"optional" json:"resultDestination" yaml:"resultDestination"`
}

