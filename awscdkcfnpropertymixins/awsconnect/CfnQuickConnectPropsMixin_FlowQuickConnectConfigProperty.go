package awsconnect


// The flow configuration.
//
// This is required only if QuickConnectType is FLOW.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   flowQuickConnectConfigProperty := &FlowQuickConnectConfigProperty{
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-flowquickconnectconfig.html
//
type CfnQuickConnectPropsMixin_FlowQuickConnectConfigProperty struct {
	// The identifier of the contact flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-flowquickconnectconfig.html#cfn-connect-quickconnect-flowquickconnectconfig-contactflowarn
	//
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
}

