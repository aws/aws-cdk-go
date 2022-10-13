package awscodepipeline


// Specifies the constraints on the number of input and output artifacts an action can have.
//
// The constraints for each action type are documented on the
// {@link https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html Pipeline Structure Reference} page.
//
// Example:
//   // MyAction is some action type that produces variables, like EcrSourceAction
//   myAction := NewMyAction(&myActionProps{
//   	// ...
//   	actionName: jsii.String("myAction"),
//   })
//   NewOtherAction(&otherActionProps{
//   	// ...
//   	config: myAction.variables.myVariable,
//   	actionName: jsii.String("otherAction"),
//   })
//
type ActionArtifactBounds struct {
	MaxInputs *float64 `field:"required" json:"maxInputs" yaml:"maxInputs"`
	MaxOutputs *float64 `field:"required" json:"maxOutputs" yaml:"maxOutputs"`
	MinInputs *float64 `field:"required" json:"minInputs" yaml:"minInputs"`
	MinOutputs *float64 `field:"required" json:"minOutputs" yaml:"minOutputs"`
}

