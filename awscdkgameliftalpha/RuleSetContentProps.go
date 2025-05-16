package awscdkgameliftalpha


// Properties for a new matchmaking ruleSet content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   var ruleSetBody iRuleSetBody
//
//   ruleSetContentProps := &RuleSetContentProps{
//   	Content: ruleSetBody,
//   }
//
// Experimental.
type RuleSetContentProps struct {
	// RuleSet body content.
	// Default: use a default empty RuleSet body.
	//
	// Experimental.
	Content IRuleSetBody `field:"optional" json:"content" yaml:"content"`
}

