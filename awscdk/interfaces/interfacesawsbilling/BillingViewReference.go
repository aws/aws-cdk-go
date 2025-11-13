package interfacesawsbilling


// A reference to a BillingView resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billingViewReference := &BillingViewReference{
//   	BillingViewArn: jsii.String("billingViewArn"),
//   }
//
type BillingViewReference struct {
	// The Arn of the BillingView resource.
	BillingViewArn *string `field:"required" json:"billingViewArn" yaml:"billingViewArn"`
}

