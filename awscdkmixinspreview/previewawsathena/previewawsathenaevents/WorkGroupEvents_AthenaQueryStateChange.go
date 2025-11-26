package previewawsathenaevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.athena@AthenaQueryStateChange event types for WorkGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   athenaQueryStateChange := #error#.NewAthenaQueryStateChange()
//
// Experimental.
type WorkGroupEvents_AthenaQueryStateChange interface {
}

// The jsii proxy struct for WorkGroupEvents_AthenaQueryStateChange
type jsiiProxy_WorkGroupEvents_AthenaQueryStateChange struct {
	_ byte // padding
}

// Experimental.
func NewWorkGroupEvents_AthenaQueryStateChange() WorkGroupEvents_AthenaQueryStateChange {
	_init_.Initialize()

	j := jsiiProxy_WorkGroupEvents_AthenaQueryStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.events.WorkGroupEvents.AthenaQueryStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkGroupEvents_AthenaQueryStateChange_Override(w WorkGroupEvents_AthenaQueryStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.events.WorkGroupEvents.AthenaQueryStateChange",
		nil, // no parameters
		w,
	)
}

