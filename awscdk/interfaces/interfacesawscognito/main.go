package interfacesawscognito

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IIdentityPoolPrincipalTagRef",
		reflect.TypeOf((*IIdentityPoolPrincipalTagRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolPrincipalTagRef", GoGetter: "IdentityPoolPrincipalTagRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentityPoolPrincipalTagRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IIdentityPoolRef",
		reflect.TypeOf((*IIdentityPoolRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolRef", GoGetter: "IdentityPoolRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentityPoolRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IIdentityPoolRoleAttachmentRef",
		reflect.TypeOf((*IIdentityPoolRoleAttachmentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolRoleAttachmentRef", GoGetter: "IdentityPoolRoleAttachmentRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentityPoolRoleAttachmentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.ILogDeliveryConfigurationRef",
		reflect.TypeOf((*ILogDeliveryConfigurationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "logDeliveryConfigurationRef", GoGetter: "LogDeliveryConfigurationRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ILogDeliveryConfigurationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IManagedLoginBrandingRef",
		reflect.TypeOf((*IManagedLoginBrandingRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "managedLoginBrandingRef", GoGetter: "ManagedLoginBrandingRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IManagedLoginBrandingRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.ITermsRef",
		reflect.TypeOf((*ITermsRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "termsRef", GoGetter: "TermsRef"},
		},
		func() interface{} {
			j := jsiiProxy_ITermsRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolClientRef",
		reflect.TypeOf((*IUserPoolClientRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolClientRef", GoGetter: "UserPoolClientRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolClientRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolDomainRef",
		reflect.TypeOf((*IUserPoolDomainRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolDomainRef", GoGetter: "UserPoolDomainRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolDomainRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolGroupRef",
		reflect.TypeOf((*IUserPoolGroupRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolGroupRef", GoGetter: "UserPoolGroupRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolGroupRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolIdentityProviderRef",
		reflect.TypeOf((*IUserPoolIdentityProviderRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdentityProviderRef", GoGetter: "UserPoolIdentityProviderRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolIdentityProviderRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolRef",
		reflect.TypeOf((*IUserPoolRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolRef", GoGetter: "UserPoolRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolResourceServerRef",
		reflect.TypeOf((*IUserPoolResourceServerRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolResourceServerRef", GoGetter: "UserPoolResourceServerRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolResourceServerRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolRiskConfigurationAttachmentRef",
		reflect.TypeOf((*IUserPoolRiskConfigurationAttachmentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolRiskConfigurationAttachmentRef", GoGetter: "UserPoolRiskConfigurationAttachmentRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolRiskConfigurationAttachmentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolUICustomizationAttachmentRef",
		reflect.TypeOf((*IUserPoolUICustomizationAttachmentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolUiCustomizationAttachmentRef", GoGetter: "UserPoolUiCustomizationAttachmentRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolUICustomizationAttachmentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolUserRef",
		reflect.TypeOf((*IUserPoolUserRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolUserRef", GoGetter: "UserPoolUserRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolUserRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.interfaces.aws_cognito.IUserPoolUserToGroupAttachmentRef",
		reflect.TypeOf((*IUserPoolUserToGroupAttachmentRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolUserToGroupAttachmentRef", GoGetter: "UserPoolUserToGroupAttachmentRef"},
		},
		func() interface{} {
			j := jsiiProxy_IUserPoolUserToGroupAttachmentRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__interfacesIEnvironmentAware)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.IdentityPoolPrincipalTagReference",
		reflect.TypeOf((*IdentityPoolPrincipalTagReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.IdentityPoolReference",
		reflect.TypeOf((*IdentityPoolReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.IdentityPoolRoleAttachmentReference",
		reflect.TypeOf((*IdentityPoolRoleAttachmentReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.LogDeliveryConfigurationReference",
		reflect.TypeOf((*LogDeliveryConfigurationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.ManagedLoginBrandingReference",
		reflect.TypeOf((*ManagedLoginBrandingReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.TermsReference",
		reflect.TypeOf((*TermsReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolClientReference",
		reflect.TypeOf((*UserPoolClientReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolDomainReference",
		reflect.TypeOf((*UserPoolDomainReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolGroupReference",
		reflect.TypeOf((*UserPoolGroupReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolIdentityProviderReference",
		reflect.TypeOf((*UserPoolIdentityProviderReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolReference",
		reflect.TypeOf((*UserPoolReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolResourceServerReference",
		reflect.TypeOf((*UserPoolResourceServerReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolRiskConfigurationAttachmentReference",
		reflect.TypeOf((*UserPoolRiskConfigurationAttachmentReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolUICustomizationAttachmentReference",
		reflect.TypeOf((*UserPoolUICustomizationAttachmentReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolUserReference",
		reflect.TypeOf((*UserPoolUserReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.interfaces.aws_cognito.UserPoolUserToGroupAttachmentReference",
		reflect.TypeOf((*UserPoolUserToGroupAttachmentReference)(nil)).Elem(),
	)
}
