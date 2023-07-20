package awsssm


// Configuration options for sending command output to Amazon CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchOutputConfigProperty := &CloudWatchOutputConfigProperty{
//   	CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	CloudWatchOutputEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-cloudwatchoutputconfig.html
//
type CfnMaintenanceWindowTask_CloudWatchOutputConfigProperty struct {
	// The name of the CloudWatch Logs log group where you want to send command output.
	//
	// If you don't specify a group name, AWS Systems Manager automatically creates a log group for you. The log group uses the following naming format:
	//
	// `aws/ssm/ *SystemsManagerDocumentName*`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-cloudwatchoutputconfig.html#cfn-ssm-maintenancewindowtask-cloudwatchoutputconfig-cloudwatchloggroupname
	//
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// Enables Systems Manager to send command output to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-cloudwatchoutputconfig.html#cfn-ssm-maintenancewindowtask-cloudwatchoutputconfig-cloudwatchoutputenabled
	//
	CloudWatchOutputEnabled interface{} `field:"optional" json:"cloudWatchOutputEnabled" yaml:"cloudWatchOutputEnabled"`
}

