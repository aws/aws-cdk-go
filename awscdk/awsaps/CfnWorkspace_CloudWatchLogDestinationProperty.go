package awsaps


// Configuration details for logging to CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogDestinationProperty := &CloudWatchLogDestinationProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-cloudwatchlogdestination.html
//
type CfnWorkspace_CloudWatchLogDestinationProperty struct {
	// The ARN of the CloudWatch log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-cloudwatchlogdestination.html#cfn-aps-workspace-cloudwatchlogdestination-loggrouparn
	//
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

