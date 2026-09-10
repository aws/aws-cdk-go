package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   systemManagedBlockProperty := &SystemManagedBlockProperty{
//   	ManagedBy: jsii.String("managedBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-systemmanagedblock.html
//
type CfnGatewayRulePropsMixin_SystemManagedBlockProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-systemmanagedblock.html#cfn-bedrockagentcore-gatewayrule-systemmanagedblock-managedby
	//
	ManagedBy *string `field:"optional" json:"managedBy" yaml:"managedBy"`
}

