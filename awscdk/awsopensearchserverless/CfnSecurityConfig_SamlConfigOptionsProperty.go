package awsopensearchserverless


// Describes SAML options for an OpenSearch Serverless security configuration in the form of a key-value map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samlConfigOptionsProperty := &SamlConfigOptionsProperty{
//   	Metadata: jsii.String("metadata"),
//
//   	// the properties below are optional
//   	GroupAttribute: jsii.String("groupAttribute"),
//   	SessionTimeout: jsii.Number(123),
//   	UserAttribute: jsii.String("userAttribute"),
//   }
//
type CfnSecurityConfig_SamlConfigOptionsProperty struct {
	// The XML IdP metadata file generated from your identity provider.
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// The group attribute for this SAML integration.
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// The session timeout, in minutes.
	//
	// Default is 60 minutes (12 hours).
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// A user attribute for this SAML integration.
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

