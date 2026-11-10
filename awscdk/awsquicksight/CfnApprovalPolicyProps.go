package awsquicksight


// Properties for defining a `CfnApprovalPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApprovalPolicyProps := &CfnApprovalPolicyProps{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ApplicableTo: &ApplicableToProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		GroupArns: []*string{
//   			jsii.String("groupArns"),
//   		},
//   	},
//   	ApprovalGroups: []*string{
//   		jsii.String("approvalGroups"),
//   	},
//   	AssetTypes: []*string{
//   		jsii.String("assetTypes"),
//   	},
//   	Name: jsii.String("name"),
//   	PolicyId: jsii.String("policyId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html
//
type CfnApprovalPolicyProps struct {
	// List of governed actions a policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Scoping: who the policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-applicableto
	//
	ApplicableTo interface{} `field:"required" json:"applicableTo" yaml:"applicableTo"`
	// List of approval group ARNs (e.g. QuickSight group ARNs). At least one approval group is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-approvalgroups
	//
	ApprovalGroups *[]*string `field:"required" json:"approvalGroups" yaml:"approvalGroups"`
	// List of asset types a policy applies to.
	//
	// At least one asset type is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-assettypes
	//
	AssetTypes *[]*string `field:"required" json:"assetTypes" yaml:"assetTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-policyid
	//
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-approvalpolicy.html#cfn-quicksight-approvalpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

