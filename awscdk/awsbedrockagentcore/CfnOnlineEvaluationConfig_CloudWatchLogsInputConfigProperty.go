package awsbedrockagentcore


// The configuration for reading agent traces from CloudWatch logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsInputConfigProperty := &CloudWatchLogsInputConfigProperty{
//   	ServiceNames: []*string{
//   		jsii.String("serviceNames"),
//   	},
//
//   	// the properties below are optional
//   	LogGroupNamePrefixes: []*string{
//   		jsii.String("logGroupNamePrefixes"),
//   	},
//   	LogGroupNames: []*string{
//   		jsii.String("logGroupNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html
//
type CfnOnlineEvaluationConfig_CloudWatchLogsInputConfigProperty struct {
	// The list of service names to filter traces within the specified log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig-servicenames
	//
	ServiceNames *[]*string `field:"required" json:"serviceNames" yaml:"serviceNames"`
	// The list of CloudWatch log group name prefixes to monitor for agent traces.
	//
	// Mutually exclusive with LogGroupNames; specify exactly one of the two selectors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig-loggroupnameprefixes
	//
	LogGroupNamePrefixes *[]*string `field:"optional" json:"logGroupNamePrefixes" yaml:"logGroupNamePrefixes"`
	// The list of CloudWatch log group names to monitor for agent traces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig-loggroupnames
	//
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
}

