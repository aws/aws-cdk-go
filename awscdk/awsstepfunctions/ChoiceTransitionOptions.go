package awsstepfunctions


// Options for Choice Transition.
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
type ChoiceTransitionOptions struct {
	// An optional description for the choice transition.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

