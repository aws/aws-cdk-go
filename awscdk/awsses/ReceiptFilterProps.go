package awsses


// Construction properties for a ReceiptFilter.
//
// Example:
//   ses.NewReceiptFilter(this, jsii.String("Filter"), &ReceiptFilterProps{
//   	Ip: jsii.String("1.2.3.4/16"),
//   })
//
type ReceiptFilterProps struct {
	// The ip address or range to filter.
	// Default: 0.0.0.0/0
	//
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The policy for the filter.
	// Default: Block.
	//
	Policy ReceiptFilterPolicy `field:"optional" json:"policy" yaml:"policy"`
	// The name for the receipt filter.
	// Default: a CloudFormation generated name.
	//
	ReceiptFilterName *string `field:"optional" json:"receiptFilterName" yaml:"receiptFilterName"`
}

