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
	// `AWS::OpenSearchServerless::SecurityPolicy.Policy`.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// `AWS::OpenSearchServerless::SecurityPolicy.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::OpenSearchServerless::SecurityPolicy.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::OpenSearchServerless::SecurityPolicy.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

