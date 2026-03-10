package previewawsworkspacesevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkSpacesAccess",
		reflect.TypeOf((*WorkSpacesAccess)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WorkSpacesAccess{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkSpacesAccess.WorkSpacesAccessProps",
		reflect.TypeOf((*WorkSpacesAccess_WorkSpacesAccessProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_workspaces.events.WorkspaceEvents",
		reflect.TypeOf((*WorkspaceEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "workSpacesAccessPattern", GoMethod: "WorkSpacesAccessPattern"},
		},
		func() interface{} {
			return &jsiiProxy_WorkspaceEvents{}
		},
	)
}
