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
type CfnPolicyTemplateProps struct {
	// Specifies the content that you want to use for the new policy template, written in the Cedar policy language.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// The description to attach to the new or updated policy template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the policy store that contains the template.
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
}

