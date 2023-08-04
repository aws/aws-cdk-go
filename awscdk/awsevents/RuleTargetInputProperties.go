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
type RuleTargetInputProperties struct {
	// Literal input to the target service (must be valid JSON).
	// Default: - input for the event target. If the input contains a paths map
	// values wil be extracted from event and inserted into the `inputTemplate`.
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// JsonPath to take input from the input event.
	// Default: - None. The entire matched event is passed as input
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// Paths map to extract values from event and insert into `inputTemplate`.
	// Default: - No values extracted from event.
	//
	InputPathsMap *map[string]*string `field:"optional" json:"inputPathsMap" yaml:"inputPathsMap"`
	// Input template to insert paths map into.
	// Default: - None.
	//
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

