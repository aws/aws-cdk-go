package awsbillingconductor


// A reference to a PricingPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pricingPlanReference := &PricingPlanReference{
//   	PricingPlanArn: jsii.String("pricingPlanArn"),
//   }
//
type PricingPlanReference struct {
	// The Arn of the PricingPlan resource.
	PricingPlanArn *string `field:"required" json:"pricingPlanArn" yaml:"pricingPlanArn"`
}

