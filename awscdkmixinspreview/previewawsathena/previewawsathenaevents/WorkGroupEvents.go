package previewawsathenaevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsathena"
)

// EventBridge event patterns for WorkGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var workGroupRef IWorkGroupRef
//
//   workGroupEvents := awscdkmixinspreview.Events.WorkGroupEvents_FromWorkGroup(workGroupRef)
//
// Experimental.
type WorkGroupEvents interface {
	// EventBridge event pattern for WorkGroup Athena Query State Change.
	// Experimental.
	AthenaQueryStateChangePattern(options *WorkGroupEvents_AthenaQueryStateChange_AthenaQueryStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for WorkGroupEvents
type jsiiProxy_WorkGroupEvents struct {
	_ byte // padding
}

// Create WorkGroupEvents from a WorkGroup reference.
// Experimental.
func WorkGroupEvents_FromWorkGroup(workGroupRef interfacesawsathena.IWorkGroupRef) WorkGroupEvents {
	_init_.Initialize()

	if err := validateWorkGroupEvents_FromWorkGroupParameters(workGroupRef); err != nil {
		panic(err)
	}
	var returns WorkGroupEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_athena.events.WorkGroupEvents",
		"fromWorkGroup",
		[]interface{}{workGroupRef},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkGroupEvents) AthenaQueryStateChangePattern(options *WorkGroupEvents_AthenaQueryStateChange_AthenaQueryStateChangeProps) *awsevents.EventPattern {
	if err := w.validateAthenaQueryStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		w,
		"athenaQueryStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

