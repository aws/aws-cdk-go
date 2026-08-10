package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   selfManagedLatticeResourceProperty := &SelfManagedLatticeResourceProperty{
//   	ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-selfmanagedlatticeresource.html
//
type CfnGatewayPropsMixin_SelfManagedLatticeResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-selfmanagedlatticeresource.html#cfn-bedrockagentcore-gateway-selfmanagedlatticeresource-resourceconfigurationidentifier
	//
	ResourceConfigurationIdentifier *string `field:"optional" json:"resourceConfigurationIdentifier" yaml:"resourceConfigurationIdentifier"`
}

