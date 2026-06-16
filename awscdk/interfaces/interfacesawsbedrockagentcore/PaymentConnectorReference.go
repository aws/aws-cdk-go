package interfacesawsbedrockagentcore


// A reference to a PaymentConnector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paymentConnectorReference := &PaymentConnectorReference{
//   	PaymentConnectorArn: jsii.String("paymentConnectorArn"),
//   }
//
type PaymentConnectorReference struct {
	// The PaymentConnectorArn of the PaymentConnector resource.
	PaymentConnectorArn *string `field:"required" json:"paymentConnectorArn" yaml:"paymentConnectorArn"`
}

