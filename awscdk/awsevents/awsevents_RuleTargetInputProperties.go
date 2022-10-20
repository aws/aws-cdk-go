package awsevents


// The input properties for an event target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleTargetInputProperties := &ruleTargetInputProperties{
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   	inputPathsMap: map[string]*string{
//   		"inputPathsMapKey": jsii.String("inputPathsMap"),
//   	},
//   	inputTemplate: jsii.String("inputTemplate"),
//   }
//
type RuleTargetInputProperties struct {
	// Literal input to the target service (must be valid JSON).
	Input *string `field:"optional" json:"input" yaml:"input"`
	// JsonPath to take input from the input event.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// Paths map to extract values from event and insert into `inputTemplate`.
	InputPathsMap *map[string]*string `field:"optional" json:"inputPathsMap" yaml:"inputPathsMap"`
	// Input template to insert paths map into.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

