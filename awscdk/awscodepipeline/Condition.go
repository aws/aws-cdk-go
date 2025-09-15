package awscodepipeline


// The condition for the stage.
//
// A condition is made up of the rules and the result for the condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var rule rule
//
//   condition := &Condition{
//   	Result: awscdk.Aws_codepipeline.Result_ROLLBACK,
//   	Rules: []*rule{
//   		rule,
//   	},
//   }
//
type Condition struct {
	// The action to be done when the condition is met.
	// Default: - No result action is taken.
	//
	Result Result `field:"optional" json:"result" yaml:"result"`
	// The rules that make up the condition.
	// Default: - No rules are applied.
	//
	Rules *[]Rule `field:"optional" json:"rules" yaml:"rules"`
}

