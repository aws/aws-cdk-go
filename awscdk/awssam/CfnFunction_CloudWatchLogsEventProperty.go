package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsEventProperty := &CloudWatchLogsEventProperty{
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatchlogsevent.html
//
type CfnFunction_CloudWatchLogsEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatchlogsevent.html#cfn-serverless-function-cloudwatchlogsevent-filterpattern
	//
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-cloudwatchlogsevent.html#cfn-serverless-function-cloudwatchlogsevent-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

