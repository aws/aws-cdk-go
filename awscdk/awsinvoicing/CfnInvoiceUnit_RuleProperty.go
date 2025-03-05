package awsinvoicing


// The `InvoiceUnitRule` object used to update invoice units.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	LinkedAccounts: []*string{
//   		jsii.String("linkedAccounts"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-invoiceunit-rule.html
//
type CfnInvoiceUnit_RuleProperty struct {
	// The list of `LINKED_ACCOUNT` IDs where charges are included within the invoice unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-invoicing-invoiceunit-rule.html#cfn-invoicing-invoiceunit-rule-linkedaccounts
	//
	LinkedAccounts *[]*string `field:"required" json:"linkedAccounts" yaml:"linkedAccounts"`
}

