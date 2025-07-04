package awsbillingconductor


// The possible AWS Free Tier configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeTierProperty := &FreeTierProperty{
//   	Activated: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-freetier.html
//
type CfnPricingRule_FreeTierProperty struct {
	// Activate or deactivate AWS Free Tier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-pricingrule-freetier.html#cfn-billingconductor-pricingrule-freetier-activated
	//
	Activated interface{} `field:"required" json:"activated" yaml:"activated"`
}

