package awsverifiedpermissions


// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &CfnPolicyProps{
//   	Definition: &PolicyDefinitionProperty{
//   		Static: &StaticPolicyDefinitionProperty{
//   			Statement: jsii.String("statement"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   		TemplateLinked: &TemplateLinkedPolicyDefinitionProperty{
//   			PolicyTemplateId: jsii.String("policyTemplateId"),
//
//   			// the properties below are optional
//   			Principal: &EntityIdentifierProperty{
//   				EntityId: jsii.String("entityId"),
//   				EntityType: jsii.String("entityType"),
//   			},
//   			Resource: &EntityIdentifierProperty{
//   				EntityId: jsii.String("entityId"),
//   				EntityType: jsii.String("entityType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policy.html
//
type CfnPolicyProps struct {
	// Specifies the policy type and content to use for the new or updated policy.
	//
	// The definition structure must include either a `Static` or a `TemplateLinked` element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policy.html#cfn-verifiedpermissions-policy-definition
	//
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// Specifies the `PolicyStoreId` of the policy store you want to store the policy in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policy.html#cfn-verifiedpermissions-policy-policystoreid
	//
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
}

