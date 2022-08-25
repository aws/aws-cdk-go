package awsiot


// Describes an action that updates a CloudWatch log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudwatchLogsActionProperty := &cloudwatchLogsActionProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnTopicRule_CloudwatchLogsActionProperty struct {
	// The CloudWatch log name.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The IAM role that allows access to the CloudWatch log.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

