package appdelivery

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.app_delivery.PipelineDeployStackAction",
		reflect.TypeOf((*PipelineDeployStackAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "addToDeploymentRolePolicy", GoMethod: "AddToDeploymentRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRole", GoGetter: "DeploymentRole"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
		},
		func() interface{} {
			j := jsiiProxy_PipelineDeployStackAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.app_delivery.PipelineDeployStackActionProps",
		reflect.TypeOf((*PipelineDeployStackActionProps)(nil)).Elem(),
	)
}
