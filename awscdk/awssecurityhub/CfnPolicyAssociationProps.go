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
	// The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html#cfn-securityhub-policyassociation-configurationpolicyid
	//
	ConfigurationPolicyId *string `field:"required" json:"configurationPolicyId" yaml:"configurationPolicyId"`
	// The identifier of the target account, organizational unit, or the root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html#cfn-securityhub-policyassociation-targetid
	//
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Indicates whether the target is an AWS account, organizational unit, or the organization root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-policyassociation.html#cfn-securityhub-policyassociation-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

