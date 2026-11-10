package interfacesawsbcmpricingcalculator


// A reference to a BillEstimate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   billEstimateReference := &BillEstimateReference{
//   	BillEstimateArn: jsii.String("billEstimateArn"),
//   }
//
type BillEstimateReference struct {
	// The Arn of the BillEstimate resource.
	BillEstimateArn *string `field:"required" json:"billEstimateArn" yaml:"billEstimateArn"`
}

