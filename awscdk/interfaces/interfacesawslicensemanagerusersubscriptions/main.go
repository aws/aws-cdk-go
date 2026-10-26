package interfacesawslicensemanagerusersubscriptions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_licensemanagerusersubscriptions.ILicenseServerEndpointRef",
		reflect.TypeOf((*ILicenseServerEndpointRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "licenseServerEndpointRef", GoGetter: "LicenseServerEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ILicenseServerEndpointRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_licensemanagerusersubscriptions.LicenseServerEndpointReference",
		reflect.TypeOf((*LicenseServerEndpointReference)(nil)).Elem(),
	)
}
