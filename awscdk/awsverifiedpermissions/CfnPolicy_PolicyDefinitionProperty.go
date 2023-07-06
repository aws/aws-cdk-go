package awsverifiedpermissions


// A structure that defines a Cedar policy.
//
// It includes the policy type, a description, and a policy body. This is a top level data type used to create a policy.
//
// This data type is used as a request parameter for the [CreatePolicy](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html) operation. This structure must always have either an `Static` or a `TemplateLinked` element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDefinitionProperty := &PolicyDefinitionProperty{
//   	Static: &StaticPolicyDefinitionProperty{
//   		Statement: jsii.String("statement"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	TemplateLinked: &TemplateLinkedPolicyDefinitionProperty{
//   		PolicyTemplateId: jsii.String("policyTemplateId"),
//
//   		// the properties below are optional
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
type CfnPolicy_PolicyDefinitionProperty struct {
	// A structure that describes a static policy.
	//
	// An static policy doesn't use a template or allow placeholders for entities.
	Static interface{} `field:"optional" json:"static" yaml:"static"`
	// A structure that describes a policy that was instantiated from a template.
	//
	// The template can specify placeholders for `principal` and `resource` . When you use [CreatePolicy](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html) to create a policy from a template, you specify the exact principal and resource to use for the instantiated policy.
	TemplateLinked interface{} `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

