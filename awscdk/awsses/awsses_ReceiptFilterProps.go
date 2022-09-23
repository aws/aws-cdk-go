package awsses


// Construction properties for a ReceiptFilter.
//
// Example:
//   ses.NewReceiptFilter(this, jsii.String("Filter"), &receiptFilterProps{
//   	ip: jsii.String("1.2.3.4/16"),
//   })
//
type ReceiptFilterProps struct {
	// The ip address or range to filter.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The policy for the filter.
	Policy ReceiptFilterPolicy `field:"optional" json:"policy" yaml:"policy"`
	// The name for the receipt filter.
	ReceiptFilterName *string `field:"optional" json:"receiptFilterName" yaml:"receiptFilterName"`
}

