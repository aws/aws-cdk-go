package awsbillingconductor


// A reference to a CustomLineItem resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLineItemReference := &CustomLineItemReference{
//   	CustomLineItemArn: jsii.String("customLineItemArn"),
//   }
//
type CustomLineItemReference struct {
	// The Arn of the CustomLineItem resource.
	CustomLineItemArn *string `field:"required" json:"customLineItemArn" yaml:"customLineItemArn"`
}

