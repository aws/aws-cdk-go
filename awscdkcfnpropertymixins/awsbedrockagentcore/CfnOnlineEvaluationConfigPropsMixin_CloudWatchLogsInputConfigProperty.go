package awsbedrockagentcore


// The configuration for reading agent traces from CloudWatch logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cloudWatchLogsInputConfigProperty := &CloudWatchLogsInputConfigProperty{
//   	LogGroupNames: []*string{
//   		jsii.String("logGroupNames"),
//   	},
//   	ServiceNames: []*string{
//   		jsii.String("serviceNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html
//
type CfnOnlineEvaluationConfigPropsMixin_CloudWatchLogsInputConfigProperty struct {
	// The list of CloudWatch log group names to monitor for agent traces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig-loggroupnames
	//
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
	// The list of service names to filter traces within the specified log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig-servicenames
	//
	ServiceNames *[]*string `field:"optional" json:"serviceNames" yaml:"serviceNames"`
}

