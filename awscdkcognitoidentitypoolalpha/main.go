package awscdkcognitoidentitypoolalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-cognito-identitypool-alpha.IIdentityPool",
		reflect.TypeOf((*IIdentityPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolArn", GoGetter: "IdentityPoolArn"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolName", GoGetter: "IdentityPoolName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentityPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-cognito-identitypool-alpha.IIdentityPoolRoleAttachment",
		reflect.TypeOf((*IIdentityPoolRoleAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentityPoolRoleAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-cognito-identitypool-alpha.IUserPoolAuthenticationProvider",
		reflect.TypeOf((*IUserPoolAuthenticationProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IUserPoolAuthenticationProvider{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPool",
		reflect.TypeOf((*IdentityPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoleMappings", GoMethod: "AddRoleMappings"},
			_jsii_.MemberMethod{JsiiMethod: "addUserPoolAuthentication", GoMethod: "AddUserPoolAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticatedRole", GoGetter: "AuthenticatedRole"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolArn", GoGetter: "IdentityPoolArn"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolName", GoGetter: "IdentityPoolName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unauthenticatedRole", GoGetter: "UnauthenticatedRole"},
		},
		func() interface{} {
			j := jsiiProxy_IdentityPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityPool)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolAmazonLoginProvider",
		reflect.TypeOf((*IdentityPoolAmazonLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolAppleLoginProvider",
		reflect.TypeOf((*IdentityPoolAppleLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolAuthenticationProviders",
		reflect.TypeOf((*IdentityPoolAuthenticationProviders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolDigitsLoginProvider",
		reflect.TypeOf((*IdentityPoolDigitsLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolFacebookLoginProvider",
		reflect.TypeOf((*IdentityPoolFacebookLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolGoogleLoginProvider",
		reflect.TypeOf((*IdentityPoolGoogleLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProps",
		reflect.TypeOf((*IdentityPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderType",
		reflect.TypeOf((*IdentityPoolProviderType)(nil)).Elem(),
		map[string]interface{}{
			"FACEBOOK": IdentityPoolProviderType_FACEBOOK,
			"GOOGLE": IdentityPoolProviderType_GOOGLE,
			"AMAZON": IdentityPoolProviderType_AMAZON,
			"APPLE": IdentityPoolProviderType_APPLE,
			"TWITTER": IdentityPoolProviderType_TWITTER,
			"DIGITS": IdentityPoolProviderType_DIGITS,
			"OPEN_ID": IdentityPoolProviderType_OPEN_ID,
			"SAML": IdentityPoolProviderType_SAML,
			"USER_POOL": IdentityPoolProviderType_USER_POOL,
			"CUSTOM": IdentityPoolProviderType_CUSTOM,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviderUrl",
		reflect.TypeOf((*IdentityPoolProviderUrl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IdentityPoolProviderUrl{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolProviders",
		reflect.TypeOf((*IdentityPoolProviders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachment",
		reflect.TypeOf((*IdentityPoolRoleAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IdentityPoolRoleAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityPoolRoleAttachment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachmentProps",
		reflect.TypeOf((*IdentityPoolRoleAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleMapping",
		reflect.TypeOf((*IdentityPoolRoleMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolTwitterLoginProvider",
		reflect.TypeOf((*IdentityPoolTwitterLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-cognito-identitypool-alpha.RoleMappingMatchType",
		reflect.TypeOf((*RoleMappingMatchType)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": RoleMappingMatchType_EQUALS,
			"CONTAINS": RoleMappingMatchType_CONTAINS,
			"STARTS_WITH": RoleMappingMatchType_STARTS_WITH,
			"NOTEQUAL": RoleMappingMatchType_NOTEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.RoleMappingRule",
		reflect.TypeOf((*RoleMappingRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-cognito-identitypool-alpha.UserPoolAuthenticationProvider",
		reflect.TypeOf((*UserPoolAuthenticationProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_UserPoolAuthenticationProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUserPoolAuthenticationProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.UserPoolAuthenticationProviderBindConfig",
		reflect.TypeOf((*UserPoolAuthenticationProviderBindConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.UserPoolAuthenticationProviderBindOptions",
		reflect.TypeOf((*UserPoolAuthenticationProviderBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-cognito-identitypool-alpha.UserPoolAuthenticationProviderProps",
		reflect.TypeOf((*UserPoolAuthenticationProviderProps)(nil)).Elem(),
	)
}
