package awsverifiedpermissions


// A structure that defines a static policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticPolicyDefinitionProperty := &StaticPolicyDefinitionProperty{
//   	Statement: jsii.String("statement"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type CfnPolicy_StaticPolicyDefinitionProperty struct {
	// The policy content of the static policy, written in the Cedar policy language.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// The description of the static policy.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

