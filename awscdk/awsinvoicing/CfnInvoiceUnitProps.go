package awsinvoicing


// Properties for defining a `CfnInvoiceUnit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInvoiceUnitProps := &CfnInvoiceUnitProps{
//   	InvoiceReceiver: jsii.String("invoiceReceiver"),
//   	Name: jsii.String("name"),
//   	Rule: &RuleProperty{
//   		LinkedAccounts: []*string{
//   			jsii.String("linkedAccounts"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaxInheritanceDisabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html
//
type CfnInvoiceUnitProps struct {
	// The account that receives invoices related to the invoice unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-invoicereceiver
	//
	InvoiceReceiver *string `field:"required" json:"invoiceReceiver" yaml:"invoiceReceiver"`
	// A unique name that is distinctive within your AWS .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An `InvoiceUnitRule` object used the categorize invoice units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-rule
	//
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// The assigned description for an invoice unit.
	//
	// This information can't be modified or deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tag structure that contains a tag key and value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-resourcetags
	//
	ResourceTags *[]*CfnInvoiceUnit_ResourceTagProperty `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Whether the invoice unit based tax inheritance is/ should be enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-taxinheritancedisabled
	//
	TaxInheritanceDisabled interface{} `field:"optional" json:"taxInheritanceDisabled" yaml:"taxInheritanceDisabled"`
}

