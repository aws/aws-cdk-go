package awssecurityhub


// Properties for defining a `CfnPolicyAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyAssociationProps := &CfnPolicyAssociationProps{
//   	ConfigurationPolicyId: jsii.String("configurationPolicyId"),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html
//
type CfnPolicyAssociationProps struct {
	// The universally unique identifier (UUID) of the configuration policy.
	//
	// A self-managed configuration has no UUID. The identifier of a self-managed configuration is `SELF_MANAGED_SECURITY_HUB` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html#cfn-securityhub-policyassociation-configurationpolicyid
	//
	ConfigurationPolicyId *string `field:"required" json:"configurationPolicyId" yaml:"configurationPolicyId"`
	// The identifier of the target account, organizational unit, or the root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html#cfn-securityhub-policyassociation-targetid
	//
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Specifies whether the target is an AWS account , organizational unit, or the root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html#cfn-securityhub-policyassociation-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

