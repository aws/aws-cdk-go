package mixinsawsopensearchserverless


// Describes IAM federation options for an OpenSearch Serverless security configuration in the form of a key-value map.
//
// These options define how OpenSearch Serverless integrates with external identity providers using federation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iamFederationConfigOptionsProperty := &IamFederationConfigOptionsProperty{
//   	GroupAttribute: jsii.String("groupAttribute"),
//   	UserAttribute: jsii.String("userAttribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamfederationconfigoptions.html
//
type CfnSecurityConfigPropsMixin_IamFederationConfigOptionsProperty struct {
	// The group attribute for this IAM federation integration.
	//
	// This attribute is used to map identity provider groups to OpenSearch Serverless permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamfederationconfigoptions.html#cfn-opensearchserverless-securityconfig-iamfederationconfigoptions-groupattribute
	//
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// The user attribute for this IAM federation integration.
	//
	// This attribute is used to identify users in the federated authentication process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-securityconfig-iamfederationconfigoptions.html#cfn-opensearchserverless-securityconfig-iamfederationconfigoptions-userattribute
	//
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

