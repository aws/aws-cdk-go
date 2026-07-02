package awsbedrockagentcore


// Configuration for connecting to a private resource using a self-managed VPC Lattice resource configuration.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-selfmanagedlatticeresource.html
//
type CfnHarnessPropsMixin_SelfManagedLatticeResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-selfmanagedlatticeresource.html#cfn-bedrockagentcore-harness-selfmanagedlatticeresource-resourceconfigurationidentifier
	//
	ResourceConfigurationIdentifier *string `field:"optional" json:"resourceConfigurationIdentifier" yaml:"resourceConfigurationIdentifier"`
}

