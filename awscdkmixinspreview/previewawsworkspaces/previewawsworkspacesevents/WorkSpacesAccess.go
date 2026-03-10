package previewawsworkspacesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.workspaces@WorkSpacesAccess.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workSpacesAccess := awscdkmixinspreview.Events.NewWorkSpacesAccess()
//
// Experimental.
type WorkSpacesAccess interface {
}

// The jsii proxy struct for WorkSpacesAccess
type jsiiProxy_WorkSpacesAccess struct {
	_ byte // padding
}

// Experimental.
func NewWorkSpacesAccess() WorkSpacesAccess {
	_init_.Initialize()

	j := jsiiProxy_WorkSpacesAccess{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkSpacesAccess",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkSpacesAccess_Override(w WorkSpacesAccess) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkSpacesAccess",
		nil, // no parameters
		w,
	)
}

// EventBridge event pattern for WorkSpaces Access.
// Experimental.
func WorkSpacesAccess_WorkSpacesAccessPattern(options *WorkSpacesAccess_WorkSpacesAccessProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateWorkSpacesAccess_WorkSpacesAccessPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkSpacesAccess",
		"workSpacesAccessPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

