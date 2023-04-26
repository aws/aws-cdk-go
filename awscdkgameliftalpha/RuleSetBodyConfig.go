// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha


// Interface to represent output result of a RuleSetContent binding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   ruleSetBodyConfig := &RuleSetBodyConfig{
//   	RuleSetBody: jsii.String("ruleSetBody"),
//   }
//
// Experimental.
type RuleSetBodyConfig struct {
	// Inline ruleSet body.
	// Experimental.
	RuleSetBody *string `field:"required" json:"ruleSetBody" yaml:"ruleSetBody"`
}

