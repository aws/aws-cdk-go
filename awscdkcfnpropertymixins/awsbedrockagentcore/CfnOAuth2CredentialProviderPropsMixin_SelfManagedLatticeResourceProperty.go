package awsbedrockagentcore


// Configuration for a self-managed VPC Lattice resource.
//
// You create and manage the VPC Lattice resource gateway and resource configuration, then provide the resource configuration identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   selfManagedLatticeResourceProperty := &SelfManagedLatticeResourceProperty{
//   	ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-selfmanagedlatticeresource.html
//
type CfnOAuth2CredentialProviderPropsMixin_SelfManagedLatticeResourceProperty struct {
	// The ARN or ID of the VPC Lattice resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-selfmanagedlatticeresource.html#cfn-bedrockagentcore-oauth2credentialprovider-selfmanagedlatticeresource-resourceconfigurationidentifier
	//
	ResourceConfigurationIdentifier *string `field:"optional" json:"resourceConfigurationIdentifier" yaml:"resourceConfigurationIdentifier"`
}

