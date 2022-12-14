package awsopensearchserverless


// Properties for defining a `CfnSecurityConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityConfigProps := &cfnSecurityConfigProps{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	samlOptions: &samlConfigOptionsProperty{
//   		metadata: jsii.String("metadata"),
//
//   		// the properties below are optional
//   		groupAttribute: jsii.String("groupAttribute"),
//   		sessionTimeout: jsii.Number(123),
//   		userAttribute: jsii.String("userAttribute"),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnSecurityConfigProps struct {
	// `AWS::OpenSearchServerless::SecurityConfig.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::OpenSearchServerless::SecurityConfig.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::OpenSearchServerless::SecurityConfig.SamlOptions`.
	SamlOptions interface{} `field:"optional" json:"samlOptions" yaml:"samlOptions"`
	// `AWS::OpenSearchServerless::SecurityConfig.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

