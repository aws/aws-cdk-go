package interfacesawsbedrockagentcore


// A reference to a PaymentManager resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paymentManagerReference := &PaymentManagerReference{
//   	PaymentManagerArn: jsii.String("paymentManagerArn"),
//   }
//
type PaymentManagerReference struct {
	// The PaymentManagerArn of the PaymentManager resource.
	PaymentManagerArn *string `field:"required" json:"paymentManagerArn" yaml:"paymentManagerArn"`
}

