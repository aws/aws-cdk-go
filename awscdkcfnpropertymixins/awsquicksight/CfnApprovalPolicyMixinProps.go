package awsquicksight


// Properties for CfnApprovalPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var actions interface{}
//   var assetTypes interface{}
//   var type interface{}
//
//   cfnApprovalPolicyMixinProps := &CfnApprovalPolicyMixinProps{
//   	Actions: []interface{}{
//   		actions,
//   	},
//   	ApplicableTo: &ApplicableToProperty{
//   		GroupArns: []*string{
//   			jsii.String("groupArns"),
//   		},
//   		Type: type,
//   	},
//   	ApprovalGroups: []*string{
//   		jsii.String("approvalGroups"),
//   	},
//   	AssetTypes: []interface{}{
//   		assetTypes,
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PolicyId: jsii.String("policyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html
//
type CfnApprovalPolicyMixinProps struct {
	// List of governed actions a policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Scoping: who the policy applies to.
	//
	// GROUP: `groupArns` required (one or more group ARNs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-applicableto
	//
	ApplicableTo interface{} `field:"optional" json:"applicableTo" yaml:"applicableTo"`
	// List of approval group ARNs (e.g. QuickSight group ARNs). At least one approval group is required; the upper bound is enforced per-account at the service layer via the configurable approver-group limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-approvalgroups
	//
	ApprovalGroups *[]*string `field:"optional" json:"approvalGroups" yaml:"approvalGroups"`
	// List of asset types a policy applies to.
	//
	// At least one asset type is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-assettypes
	//
	AssetTypes interface{} `field:"optional" json:"assetTypes" yaml:"assetTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-policyid
	//
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
}

