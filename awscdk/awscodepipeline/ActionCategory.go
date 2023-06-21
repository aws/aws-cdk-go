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
type ActionCategory string

const (
	ActionCategory_SOURCE ActionCategory = "SOURCE"
	ActionCategory_BUILD ActionCategory = "BUILD"
	ActionCategory_TEST ActionCategory = "TEST"
	ActionCategory_APPROVAL ActionCategory = "APPROVAL"
	ActionCategory_DEPLOY ActionCategory = "DEPLOY"
	ActionCategory_INVOKE ActionCategory = "INVOKE"
)

