package awsstepfunctions


// Properties for defining a Choice state.
//
// Example:
//   choice := sfn.NewChoice(this, jsii.String("What color is it?"), &ChoiceProps{
//   	Comment: jsii.String("color comment"),
//   })
//   handleBlueItem := sfn.NewPass(this, jsii.String("HandleBlueItem"))
//   handleOtherItemColor := sfn.NewPass(this, jsii.String("HanldeOtherItemColor"))
//   choice.When(sfn.Condition_StringEquals(jsii.String("$.color"), jsii.String("BLUE")), handleBlueItem, &ChoiceTransitionOptions{
//   	Comment: jsii.String("blue item comment"),
//   })
//   choice.Otherwise(handleOtherItemColor)
//
type ChoiceProps struct {
	// An optional description for this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
}

