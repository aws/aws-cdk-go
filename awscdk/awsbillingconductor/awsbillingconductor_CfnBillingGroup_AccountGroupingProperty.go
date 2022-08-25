package awsbillingconductor


// The set of accounts that will be under the billing group.
//
// The set of accounts resemble the linked accounts in a consolidated family.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountGroupingProperty := &accountGroupingProperty{
//   	linkedAccountIds: []*string{
//   		jsii.String("linkedAccountIds"),
//   	},
//   }
//
type CfnBillingGroup_AccountGroupingProperty struct {
	// The account IDs that make up the billing group.
	//
	// Account IDs must be a part of the consolidated billing family, and not associated with another billing group.
	LinkedAccountIds *[]*string `field:"required" json:"linkedAccountIds" yaml:"linkedAccountIds"`
}

