package awseventschemas


// Properties for defining a `CfnRegistryPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnRegistryPolicyProps := &cfnRegistryPolicyProps{
//   	policy: policy,
//   	registryName: jsii.String("registryName"),
//
//   	// the properties below are optional
//   	revisionId: jsii.String("revisionId"),
//   }
//
type CfnRegistryPolicyProps struct {
	// A resource-based policy.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The name of the registry.
	RegistryName *string `field:"required" json:"registryName" yaml:"registryName"`
	// The revision ID of the policy.
	RevisionId *string `field:"optional" json:"revisionId" yaml:"revisionId"`
}

