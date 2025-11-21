package awsbillingconductor


// The set of accounts that will be under the billing group.
//
// The set of accounts resemble the linked accounts in a consolidated billing family.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountGroupingProperty := &AccountGroupingProperty{
//   	AutoAssociate: jsii.Boolean(false),
//   	LinkedAccountIds: []*string{
//   		jsii.String("linkedAccountIds"),
//   	},
//   	ResponsibilityTransferArn: jsii.String("responsibilityTransferArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-billinggroup-accountgrouping.html
//
type CfnBillingGroup_AccountGroupingProperty struct {
	// Specifies if this billing group will automatically associate newly added AWS accounts that join your consolidated billing family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-billinggroup-accountgrouping.html#cfn-billingconductor-billinggroup-accountgrouping-autoassociate
	//
	AutoAssociate interface{} `field:"optional" json:"autoAssociate" yaml:"autoAssociate"`
	// The account IDs that make up the billing group.
	//
	// Account IDs must be a part of the consolidated billing family, and not associated with another billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-billinggroup-accountgrouping.html#cfn-billingconductor-billinggroup-accountgrouping-linkedaccountids
	//
	LinkedAccountIds *[]*string `field:"optional" json:"linkedAccountIds" yaml:"linkedAccountIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-billinggroup-accountgrouping.html#cfn-billingconductor-billinggroup-accountgrouping-responsibilitytransferarn
	//
	ResponsibilityTransferArn *string `field:"optional" json:"responsibilityTransferArn" yaml:"responsibilityTransferArn"`
}

