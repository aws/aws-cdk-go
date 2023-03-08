package awsiot


// A structure containing the alert target ARN and the role ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alertTargetProperty := &alertTargetProperty{
//   	alertTargetArn: jsii.String("alertTargetArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnSecurityProfile_AlertTargetProperty struct {
	// The Amazon Resource Name (ARN) of the notification target to which alerts are sent.
	AlertTargetArn *string `field:"required" json:"alertTargetArn" yaml:"alertTargetArn"`
	// The ARN of the role that grants permission to send alerts to the notification target.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

