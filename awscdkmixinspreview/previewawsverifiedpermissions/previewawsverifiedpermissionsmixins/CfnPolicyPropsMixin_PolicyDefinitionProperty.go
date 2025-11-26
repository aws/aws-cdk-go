package previewawsverifiedpermissionsmixins


// A structure that defines a Cedar policy.
//
// It includes the policy type, a description, and a policy body. This is a top level data type used to create a policy.
//
// This data type is used as a request parameter for the [CreatePolicy](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html) operation. This structure must always have either an `Static` or a `TemplateLinked` element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDefinitionProperty := &PolicyDefinitionProperty{
//   	Static: &StaticPolicyDefinitionProperty{
//   		Description: jsii.String("description"),
//   		Statement: jsii.String("statement"),
//   	},
//   	TemplateLinked: &TemplateLinkedPolicyDefinitionProperty{
//   		PolicyTemplateId: jsii.String("policyTemplateId"),
//   		Principal: &EntityIdentifierProperty{
//   			EntityId: jsii.String("entityId"),
//   			EntityType: jsii.String("entityType"),
//   		},
//   		Resource: &EntityIdentifierProperty{
//   			EntityId: jsii.String("entityId"),
//   			EntityType: jsii.String("entityType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-policydefinition.html
//
type CfnPolicyPropsMixin_PolicyDefinitionProperty struct {
	// A structure that describes a static policy.
	//
	// An static policy doesn't use a template or allow placeholders for entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-policydefinition.html#cfn-verifiedpermissions-policy-policydefinition-static
	//
	Static interface{} `field:"optional" json:"static" yaml:"static"`
	// A structure that describes a policy that was instantiated from a template.
	//
	// The template can specify placeholders for `principal` and `resource` . When you use [CreatePolicy](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html) to create a policy from a template, you specify the exact principal and resource to use for the instantiated policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-policydefinition.html#cfn-verifiedpermissions-policy-policydefinition-templatelinked
	//
	TemplateLinked interface{} `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

