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
	// The description of the security configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the security configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// SAML options for the security configuration in the form of a key-value map.
	SamlOptions interface{} `field:"optional" json:"samlOptions" yaml:"samlOptions"`
	// The type of security configuration.
	//
	// Currently the only option is `saml` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

