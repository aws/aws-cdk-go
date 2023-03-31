package awsssm


// Configuration options for sending command output to Amazon CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchOutputConfigProperty := &cloudWatchOutputConfigProperty{
//   	cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   	cloudWatchOutputEnabled: jsii.Boolean(false),
//   }
//
type CfnMaintenanceWindowTask_CloudWatchOutputConfigProperty struct {
	// The name of the CloudWatch Logs log group where you want to send command output.
	//
	// If you don't specify a group name, AWS Systems Manager automatically creates a log group for you. The log group uses the following naming format:
	//
	// `aws/ssm/ *SystemsManagerDocumentName*`.
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// Enables Systems Manager to send command output to CloudWatch Logs.
	CloudWatchOutputEnabled interface{} `field:"optional" json:"cloudWatchOutputEnabled" yaml:"cloudWatchOutputEnabled"`
}

