package awsverifiedpermissions


// Properties for defining a `CfnPolicyTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyTemplateProps := &CfnPolicyTemplateProps{
//   	Statement: jsii.String("statement"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html
//
type CfnPolicyTemplateProps struct {
	// Specifies the content that you want to use for the new , written in the policy language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html#cfn-verifiedpermissions-policytemplate-statement
	//
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// The description to attach to the new or updated policy template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html#cfn-verifiedpermissions-policytemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the policy store that contains the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policytemplate.html#cfn-verifiedpermissions-policytemplate-policystoreid
	//
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
}

