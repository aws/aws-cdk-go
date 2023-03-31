package awsfms


// Properties for defining a `CfnNotificationChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNotificationChannelProps := &cfnNotificationChannelProps{
//   	snsRoleName: jsii.String("snsRoleName"),
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type CfnNotificationChannelProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that allows Amazon SNS to record AWS Firewall Manager activity.
	SnsRoleName *string `field:"required" json:"snsRoleName" yaml:"snsRoleName"`
	// The Amazon Resource Name (ARN) of the SNS topic that collects notifications from AWS Firewall Manager .
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

