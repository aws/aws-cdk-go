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
//   userQuickConnectConfigProperty := &UserQuickConnectConfigProperty{
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   	UserArn: jsii.String("userArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html
//
type CfnQuickConnect_UserQuickConnectConfigProperty struct {
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html#cfn-connect-quickconnect-userquickconnectconfig-contactflowarn
	//
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The Amazon Resource Name (ARN) of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-userquickconnectconfig.html#cfn-connect-quickconnect-userquickconnectconfig-userarn
	//
	UserArn *string `field:"required" json:"userArn" yaml:"userArn"`
}

