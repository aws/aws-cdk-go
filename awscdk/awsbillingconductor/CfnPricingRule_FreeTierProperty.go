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
type CfnPricingRule_FreeTierProperty struct {
	// Activate or deactivate AWS Free Tier.
	Activated interface{} `field:"required" json:"activated" yaml:"activated"`
}

