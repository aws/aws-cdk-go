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
//   	OpenSearchServerlessEntityId: jsii.String("openSearchServerlessEntityId"),
//   	SessionTimeout: jsii.Number(123),
//   	UserAttribute: jsii.String("userAttribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html
//
type CfnSecurityConfig_SamlConfigOptionsProperty struct {
	// The XML IdP metadata file generated from your identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html#cfn-opensearchserverless-securityconfig-samlconfigoptions-metadata
	//
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// The group attribute for this SAML integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html#cfn-opensearchserverless-securityconfig-samlconfigoptions-groupattribute
	//
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// Custom entity ID attribute to override the default entity ID for this SAML integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html#cfn-opensearchserverless-securityconfig-samlconfigoptions-opensearchserverlessentityid
	//
	OpenSearchServerlessEntityId *string `field:"optional" json:"openSearchServerlessEntityId" yaml:"openSearchServerlessEntityId"`
	// The session timeout, in minutes.
	//
	// Default is 60 minutes (12 hours).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html#cfn-opensearchserverless-securityconfig-samlconfigoptions-sessiontimeout
	//
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// A user attribute for this SAML integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html#cfn-opensearchserverless-securityconfig-samlconfigoptions-userattribute
	//
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

