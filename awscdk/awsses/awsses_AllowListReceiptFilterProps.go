package awsses


// Construction properties for am AllowListReceiptFilter.
//
// Example:
//   ses.NewAllowListReceiptFilter(this, jsii.String("AllowList"), &allowListReceiptFilterProps{
//   	ips: []*string{
//   		jsii.String("10.0.0.0/16"),
//   		jsii.String("1.2.3.4/16"),
//   	},
//   })
//
type AllowListReceiptFilterProps struct {
	// A list of ip addresses or ranges to allow list.
	Ips *[]*string `field:"required" json:"ips" yaml:"ips"`
}

