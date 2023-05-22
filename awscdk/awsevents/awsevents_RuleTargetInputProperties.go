package awsevents


// The input properties for an event target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleTargetInputProperties := &RuleTargetInputProperties{
//   	Input: jsii.String("input"),
//   	InputPath: jsii.String("inputPath"),
//   	InputPathsMap: map[string]*string{
//   		"inputPathsMapKey": jsii.String("inputPathsMap"),
//   	},
//   	InputTemplate: jsii.String("inputTemplate"),
//   }
//
// Experimental.
type RuleTargetInputProperties struct {
	// Literal input to the target service (must be valid JSON).
	// Experimental.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// JsonPath to take input from the input event.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// Paths map to extract values from event and insert into `inputTemplate`.
	// Experimental.
	InputPathsMap *map[string]*string `field:"optional" json:"inputPathsMap" yaml:"inputPathsMap"`
	// Input template to insert paths map into.
	// Experimental.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

