package awsbillingconductor


// The preferences and settings that will be used to compute the AWS charges for a billing group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computationPreferenceProperty := &computationPreferenceProperty{
//   	pricingPlanArn: jsii.String("pricingPlanArn"),
//   }
//
type CfnBillingGroup_ComputationPreferenceProperty struct {
	// The Amazon Resource Name (ARN) of the pricing plan used to compute the AWS charges for a billing group.
	PricingPlanArn *string `field:"required" json:"pricingPlanArn" yaml:"pricingPlanArn"`
}

