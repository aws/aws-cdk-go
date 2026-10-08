package interfacesawsram

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_ram.IPermissionAssociationRef",
		reflect.TypeOf((*IPermissionAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionAssociationRef", GoGetter: "PermissionAssociationRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPermissionAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_ram.IPermissionRef",
		reflect.TypeOf((*IPermissionRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionRef", GoGetter: "PermissionRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPermissionRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_ram.IPrincipalAssociationRef",
		reflect.TypeOf((*IPrincipalAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "principalAssociationRef", GoGetter: "PrincipalAssociationRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPrincipalAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_ram.IResourceAssociationRef",
		reflect.TypeOf((*IResourceAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAssociationRef", GoGetter: "ResourceAssociationRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IResourceAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_ram.IResourceShareRef",
		reflect.TypeOf((*IResourceShareRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "resourceShareRef", GoGetter: "ResourceShareRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IResourceShareRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_ram.ISourceAssociationRef",
		reflect.TypeOf((*ISourceAssociationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sourceAssociationRef", GoGetter: "SourceAssociationRef"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ISourceAssociationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_ram.PermissionAssociationReference",
		reflect.TypeOf((*PermissionAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_ram.PermissionReference",
		reflect.TypeOf((*PermissionReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_ram.PrincipalAssociationReference",
		reflect.TypeOf((*PrincipalAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_ram.ResourceAssociationReference",
		reflect.TypeOf((*ResourceAssociationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_ram.ResourceShareReference",
		reflect.TypeOf((*ResourceShareReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_ram.SourceAssociationReference",
		reflect.TypeOf((*SourceAssociationReference)(nil)).Elem(),
	)
}
