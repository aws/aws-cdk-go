package interfacesawswafregional


// A reference to a RateBasedRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rateBasedRuleReference := &RateBasedRuleReference{
//   	RateBasedRuleId: jsii.String("rateBasedRuleId"),
//   }
//
type RateBasedRuleReference struct {
	// The Id of the RateBasedRule resource.
	RateBasedRuleId *string `field:"required" json:"rateBasedRuleId" yaml:"rateBasedRuleId"`
}

