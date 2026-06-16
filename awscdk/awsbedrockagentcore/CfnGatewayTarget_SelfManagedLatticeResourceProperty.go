package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedLatticeResourceProperty := &SelfManagedLatticeResourceProperty{
//   	ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-selfmanagedlatticeresource.html
//
type CfnGatewayTarget_SelfManagedLatticeResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-selfmanagedlatticeresource.html#cfn-bedrockagentcore-gatewaytarget-selfmanagedlatticeresource-resourceconfigurationidentifier
	//
	ResourceConfigurationIdentifier *string `field:"required" json:"resourceConfigurationIdentifier" yaml:"resourceConfigurationIdentifier"`
}

