package interfacesawsarczonalshift

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_arczonalshift.AutoshiftObserverNotificationStatusReference",
		reflect.TypeOf((*AutoshiftObserverNotificationStatusReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_arczonalshift.IAutoshiftObserverNotificationStatusRef",
		reflect.TypeOf((*IAutoshiftObserverNotificationStatusRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoshiftObserverNotificationStatusRef", GoGetter: "AutoshiftObserverNotificationStatusRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IAutoshiftObserverNotificationStatusRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_arczonalshift.IZonalAutoshiftConfigurationRef",
		reflect.TypeOf((*IZonalAutoshiftConfigurationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "zonalAutoshiftConfigurationRef", GoGetter: "ZonalAutoshiftConfigurationRef"},
		},
		func() interface{} {
			j := jsiiProxy_IZonalAutoshiftConfigurationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_arczonalshift.ZonalAutoshiftConfigurationReference",
		reflect.TypeOf((*ZonalAutoshiftConfigurationReference)(nil)).Elem(),
	)
}
