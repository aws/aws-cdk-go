package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The CodePipeline variables that are global, not bound to a specific action.
//
// This class defines a bunch of static fields that represent the different variables.
// These can be used can be used in any action configuration.
//
// Example:
//   // OtherAction is some action type that produces variables, like EcrSourceAction
//   // OtherAction is some action type that produces variables, like EcrSourceAction
//   NewOtherAction(&otherActionProps{
//   	// ...
//   	config: codepipeline.GlobalVariables_ExecutionId(),
//   	actionName: jsii.String("otherAction"),
//   })
//
type GlobalVariables interface {
}

// The jsii proxy struct for GlobalVariables
type jsiiProxy_GlobalVariables struct {
	_ byte // padding
}

func NewGlobalVariables() GlobalVariables {
	_init_.Initialize()

	j := jsiiProxy_GlobalVariables{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewGlobalVariables_Override(g GlobalVariables) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
		nil, // no parameters
		g,
	)
}

func GlobalVariables_ExecutionId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codepipeline.GlobalVariables",
		"executionId",
		&returns,
	)
	return returns
}

