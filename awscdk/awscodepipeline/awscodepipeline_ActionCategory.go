package awscodepipeline


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
// Experimental.
type ActionCategory string

const (
	// Experimental.
	ActionCategory_SOURCE ActionCategory = "SOURCE"
	// Experimental.
	ActionCategory_BUILD ActionCategory = "BUILD"
	// Experimental.
	ActionCategory_TEST ActionCategory = "TEST"
	// Experimental.
	ActionCategory_APPROVAL ActionCategory = "APPROVAL"
	// Experimental.
	ActionCategory_DEPLOY ActionCategory = "DEPLOY"
	// Experimental.
	ActionCategory_INVOKE ActionCategory = "INVOKE"
)

