package awspricingplanmanager


// Properties for CfnSubscriptionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSubscriptionMixinProps := &CfnSubscriptionMixinProps{
//   	PlanFamily: jsii.String("planFamily"),
//   	PlanTier: jsii.String("planTier"),
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	UsageLevel: jsii.String("usageLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pricingplanmanager-subscription.html
//
type CfnSubscriptionMixinProps struct {
	// The name of the pricing plan family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pricingplanmanager-subscription.html#cfn-pricingplanmanager-subscription-planfamily
	//
	PlanFamily *string `field:"optional" json:"planFamily" yaml:"planFamily"`
	// The tier of the pricing plan.
	//
	// Upgrades take effect immediately. However, rolling back an upgrade does not revert billing instantly; it schedules a downgrade to the end of the current billing period, and the higher-tier charge applies for the remainder of that month. While a downgrade is scheduled, the CurrentPlanTier property reports the tier currently being billed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pricingplanmanager-subscription.html#cfn-pricingplanmanager-subscription-plantier
	//
	PlanTier *string `field:"optional" json:"planTier" yaml:"planTier"`
	// The ARNs of resources associated with the subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pricingplanmanager-subscription.html#cfn-pricingplanmanager-subscription-resourcearns
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pricingplanmanager-subscription.html#cfn-pricingplanmanager-subscription-usagelevel
	//
	UsageLevel *string `field:"optional" json:"usageLevel" yaml:"usageLevel"`
}

