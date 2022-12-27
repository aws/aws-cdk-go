package awsopensearchserverless


// Properties for defining a `CfnAccessPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPolicyProps := &cfnAccessPolicyProps{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	policy: jsii.String("policy"),
//   	type: jsii.String("type"),
//   }
//
type CfnAccessPolicyProps struct {
	// `AWS::OpenSearchServerless::AccessPolicy.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::OpenSearchServerless::AccessPolicy.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::OpenSearchServerless::AccessPolicy.Policy`.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// `AWS::OpenSearchServerless::AccessPolicy.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

