package awsconnect


// Contains information about the quick connect configuration settings for a user.
//
// The contact flow must be of type Transfer to Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userQuickConnectConfigProperty := &userQuickConnectConfigProperty{
//   	contactFlowArn: jsii.String("contactFlowArn"),
//   	userArn: jsii.String("userArn"),
//   }
//
type CfnQuickConnect_UserQuickConnectConfigProperty struct {
	// The Amazon Resource Name (ARN) of the contact flow.
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The Amazon Resource Name (ARN) of the user.
	UserArn *string `field:"required" json:"userArn" yaml:"userArn"`
}

