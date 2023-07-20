package awsopensearchserverless


// Properties for defining a `CfnSecurityConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityConfigProps := &CfnSecurityConfigProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	SamlOptions: &SamlConfigOptionsProperty{
//   		Metadata: jsii.String("metadata"),
//
//   		// the properties below are optional
//   		GroupAttribute: jsii.String("groupAttribute"),
//   		SessionTimeout: jsii.Number(123),
//   		UserAttribute: jsii.String("userAttribute"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securityconfig.html
//
type CfnSecurityConfigProps struct {
	// The description of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securityconfig.html#cfn-opensearchserverless-securityconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securityconfig.html#cfn-opensearchserverless-securityconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// SAML options for the security configuration in the form of a key-value map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securityconfig.html#cfn-opensearchserverless-securityconfig-samloptions
	//
	SamlOptions interface{} `field:"optional" json:"samlOptions" yaml:"samlOptions"`
	// The type of security configuration.
	//
	// Currently the only option is `saml` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securityconfig.html#cfn-opensearchserverless-securityconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

