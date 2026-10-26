package awscodepipeline


// Example:
//   // MyAction is some action type that produces variables, like EcrSourceAction
//   myAction := newMyAction(&myActionProps{
//   	// ...
//   	actionName: jsii.String("myAction"),
//   })
//   newOtherAction(&otherActionProps{
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
	ActionCategory_COMPUTE ActionCategory = "COMPUTE"
)

