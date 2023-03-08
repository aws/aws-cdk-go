package awsopensearchserverless


// Properties for defining a `CfnAccessPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPolicyProps := &CfnAccessPolicyProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//   }
//
type CfnAccessPolicyProps struct {
	// The description of the policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the policy.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The JSON policy document without any whitespaces.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The type of access policy.
	//
	// Currently the only option is `data` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

