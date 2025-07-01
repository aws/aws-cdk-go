package awsopensearchserverless


// Properties for defining a `CfnSecurityPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityPolicyProps := &CfnSecurityPolicyProps{
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html
//
type CfnSecurityPolicyProps struct {
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The JSON policy document without any whitespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-policy
	//
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The type of security policy.
	//
	// Can be either `encryption` or `network` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the security policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html#cfn-opensearchserverless-securitypolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

