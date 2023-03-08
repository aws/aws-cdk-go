package awslookoutmetrics


// Contains information about the SNS topic to which you want to send your alerts and the IAM role that has access to that topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sNSConfigurationProperty := &sNSConfigurationProperty{
//   	roleArn: jsii.String("roleArn"),
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type CfnAlert_SNSConfigurationProperty struct {
	// The ARN of the IAM role that has access to the target SNS topic.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the target SNS topic.
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

