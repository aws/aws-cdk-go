package awsses


// Construction properties for a WhiteListReceiptFilter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   whiteListReceiptFilterProps := &WhiteListReceiptFilterProps{
//   	Ips: []*string{
//   		jsii.String("ips"),
//   	},
//   }
//
// Deprecated: use `AllowListReceiptFilterProps`.
type WhiteListReceiptFilterProps struct {
	// A list of ip addresses or ranges to allow list.
	// Deprecated: use `AllowListReceiptFilterProps`.
	Ips *[]*string `field:"required" json:"ips" yaml:"ips"`
}

