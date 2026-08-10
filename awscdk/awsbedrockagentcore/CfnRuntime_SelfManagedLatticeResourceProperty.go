package awsbedrockagentcore


// Self-managed VPC Lattice resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedLatticeResourceProperty := &SelfManagedLatticeResourceProperty{
//   	ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-selfmanagedlatticeresource.html
//
type CfnRuntime_SelfManagedLatticeResourceProperty struct {
	// The identifier of the VPC Lattice resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-selfmanagedlatticeresource.html#cfn-bedrockagentcore-runtime-selfmanagedlatticeresource-resourceconfigurationidentifier
	//
	ResourceConfigurationIdentifier *string `field:"required" json:"resourceConfigurationIdentifier" yaml:"resourceConfigurationIdentifier"`
}

