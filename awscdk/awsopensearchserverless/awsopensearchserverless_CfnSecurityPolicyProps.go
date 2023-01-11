package awsopensearchserverless


// Properties for defining a `CfnSecurityPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityPolicyProps := &cfnSecurityPolicyProps{
//   	policy: jsii.String("policy"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnSecurityPolicyProps struct {
	// The JSON policy document without any whitespaces.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The description of the security policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of security policy.
	//
	// Can be either `encryption` or `network` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

