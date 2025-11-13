package interfacesawsxray


// A reference to a SamplingRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samplingRuleReference := &SamplingRuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type SamplingRuleReference struct {
	// The RuleARN of the SamplingRule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

