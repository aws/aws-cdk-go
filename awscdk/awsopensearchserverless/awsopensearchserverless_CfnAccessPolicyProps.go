package awsopensearchserverless


// Properties for defining a `CfnAccessPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPolicyProps := &CfnAccessPolicyProps{
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type CfnAccessPolicyProps struct {
	// The name of the policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The JSON policy document without any whitespaces.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The type of access policy.
	//
	// Currently the only option is `data` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

