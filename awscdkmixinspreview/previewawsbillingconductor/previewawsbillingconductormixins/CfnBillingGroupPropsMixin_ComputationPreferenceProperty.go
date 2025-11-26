package previewawsbillingconductormixins


// The preferences and settings that will be used to compute the AWS charges for a billing group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   computationPreferenceProperty := &ComputationPreferenceProperty{
//   	PricingPlanArn: jsii.String("pricingPlanArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-billinggroup-computationpreference.html
//
type CfnBillingGroupPropsMixin_ComputationPreferenceProperty struct {
	// The Amazon Resource Name (ARN) of the pricing plan used to compute the AWS charges for a billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-billinggroup-computationpreference.html#cfn-billingconductor-billinggroup-computationpreference-pricingplanarn
	//
	PricingPlanArn *string `field:"optional" json:"pricingPlanArn" yaml:"pricingPlanArn"`
}

