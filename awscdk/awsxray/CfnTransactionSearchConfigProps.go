package awsxray


// Properties for defining a `CfnTransactionSearchConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransactionSearchConfigProps := &CfnTransactionSearchConfigProps{
//   	CloudWatchLogsConfiguration: &CloudWatchLogsConfigurationProperty{
//   		SpansLogGroupKmsKeyId: jsii.String("spansLogGroupKmsKeyId"),
//   	},
//   	IndexingPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-transactionsearchconfig.html
//
type CfnTransactionSearchConfigProps struct {
	// Configuration for the CloudWatch Logs log group that X-Ray creates when this resource enables Transaction Search.
	//
	// It is used only at that time and is never returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-transactionsearchconfig.html#cfn-xray-transactionsearchconfig-cloudwatchlogsconfiguration
	//
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// Determines the percentage of traces indexed from CloudWatch Logs to X-Ray.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-xray-transactionsearchconfig.html#cfn-xray-transactionsearchconfig-indexingpercentage
	//
	IndexingPercentage *float64 `field:"optional" json:"indexingPercentage" yaml:"indexingPercentage"`
}

