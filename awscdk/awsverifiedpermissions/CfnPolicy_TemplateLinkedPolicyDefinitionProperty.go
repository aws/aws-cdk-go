package awsverifiedpermissions


// A structure that describes a policy created by instantiating a policy template.
//
// > You can't directly update a template-linked policy. You must update the associated policy template instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateLinkedPolicyDefinitionProperty := &TemplateLinkedPolicyDefinitionProperty{
//   	PolicyTemplateId: jsii.String("policyTemplateId"),
//
//   	// the properties below are optional
//   	Principal: &EntityIdentifierProperty{
//   		EntityId: jsii.String("entityId"),
//   		EntityType: jsii.String("entityType"),
//   	},
//   	Resource: &EntityIdentifierProperty{
//   		EntityId: jsii.String("entityId"),
//   		EntityType: jsii.String("entityType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-templatelinkedpolicydefinition.html
//
type CfnPolicy_TemplateLinkedPolicyDefinitionProperty struct {
	// The unique identifier of the policy template used to create this policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-templatelinkedpolicydefinition.html#cfn-verifiedpermissions-policy-templatelinkedpolicydefinition-policytemplateid
	//
	PolicyTemplateId *string `field:"required" json:"policyTemplateId" yaml:"policyTemplateId"`
	// The principal associated with this template-linked policy.
	//
	// Verified Permissions substitutes this principal for the `?principal` placeholder in the policy template when it evaluates an authorization request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-templatelinkedpolicydefinition.html#cfn-verifiedpermissions-policy-templatelinkedpolicydefinition-principal
	//
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// The resource associated with this template-linked policy.
	//
	// Verified Permissions substitutes this resource for the `?resource` placeholder in the policy template when it evaluates an authorization request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-templatelinkedpolicydefinition.html#cfn-verifiedpermissions-policy-templatelinkedpolicydefinition-resource
	//
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}

