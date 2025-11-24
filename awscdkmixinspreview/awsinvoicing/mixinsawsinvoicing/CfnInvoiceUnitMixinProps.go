package mixinsawsinvoicing


// Properties for CfnInvoiceUnitPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInvoiceUnitMixinProps := &CfnInvoiceUnitMixinProps{
//   	Description: jsii.String("description"),
//   	InvoiceReceiver: jsii.String("invoiceReceiver"),
//   	Name: jsii.String("name"),
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Rule: &RuleProperty{
//   		LinkedAccounts: []*string{
//   			jsii.String("linkedAccounts"),
//   		},
//   	},
//   	TaxInheritanceDisabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html
//
type CfnInvoiceUnitMixinProps struct {
	// The assigned description for an invoice unit.
	//
	// This information can't be modified or deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The account that receives invoices related to the invoice unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-invoicereceiver
	//
	InvoiceReceiver *string `field:"optional" json:"invoiceReceiver" yaml:"invoiceReceiver"`
	// A unique name that is distinctive within your AWS .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tag structure that contains a tag key and value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-resourcetags
	//
	ResourceTags *[]*CfnInvoiceUnitPropsMixin_ResourceTagProperty `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// An `InvoiceUnitRule` object used the categorize invoice units.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-rule
	//
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// Whether the invoice unit based tax inheritance is/ should be enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-invoicing-invoiceunit.html#cfn-invoicing-invoiceunit-taxinheritancedisabled
	//
	TaxInheritanceDisabled interface{} `field:"optional" json:"taxInheritanceDisabled" yaml:"taxInheritanceDisabled"`
}

