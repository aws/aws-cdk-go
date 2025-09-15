package awsdatabrew


// A reference to a Ruleset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rulesetReference := &RulesetReference{
//   	RulesetName: jsii.String("rulesetName"),
//   }
//
type RulesetReference struct {
	// The Name of the Ruleset resource.
	RulesetName *string `field:"required" json:"rulesetName" yaml:"rulesetName"`
}

