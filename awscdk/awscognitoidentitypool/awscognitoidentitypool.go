package awscognitoidentitypool

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"monocdk.aws_cognito_identitypool.IIdentityPool",
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
		"monocdk.aws_cognito_identitypool.IIdentityPoolRoleAttachment",
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
		"monocdk.aws_cognito_identitypool.IUserPoolAuthenticationProvider",
		reflect.TypeOf((*IUserPoolAuthenticationProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IUserPoolAuthenticationProvider{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_cognito_identitypool.IdentityPool",
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
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unauthenticatedRole", GoGetter: "UnauthenticatedRole"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_IdentityPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityPool)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolAmazonLoginProvider",
		reflect.TypeOf((*IdentityPoolAmazonLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolAppleLoginProvider",
		reflect.TypeOf((*IdentityPoolAppleLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolAuthenticationProviders",
		reflect.TypeOf((*IdentityPoolAuthenticationProviders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolDigitsLoginProvider",
		reflect.TypeOf((*IdentityPoolDigitsLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolFacebookLoginProvider",
		reflect.TypeOf((*IdentityPoolFacebookLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolGoogleLoginProvider",
		reflect.TypeOf((*IdentityPoolGoogleLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolProps",
		reflect.TypeOf((*IdentityPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderType",
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
		"monocdk.aws_cognito_identitypool.IdentityPoolProviderUrl",
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
		"monocdk.aws_cognito_identitypool.IdentityPoolProviders",
		reflect.TypeOf((*IdentityPoolProviders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleAttachment",
		reflect.TypeOf((*IdentityPoolRoleAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_IdentityPoolRoleAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityPoolRoleAttachment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleAttachmentProps",
		reflect.TypeOf((*IdentityPoolRoleAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolRoleMapping",
		reflect.TypeOf((*IdentityPoolRoleMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.IdentityPoolTwitterLoginProvider",
		reflect.TypeOf((*IdentityPoolTwitterLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_cognito_identitypool.RoleMappingMatchType",
		reflect.TypeOf((*RoleMappingMatchType)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": RoleMappingMatchType_EQUALS,
			"CONTAINS": RoleMappingMatchType_CONTAINS,
			"STARTS_WITH": RoleMappingMatchType_STARTS_WITH,
			"NOTEQUAL": RoleMappingMatchType_NOTEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.RoleMappingRule",
		reflect.TypeOf((*RoleMappingRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProvider",
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
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProviderBindConfig",
		reflect.TypeOf((*UserPoolAuthenticationProviderBindConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProviderBindOptions",
		reflect.TypeOf((*UserPoolAuthenticationProviderBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_cognito_identitypool.UserPoolAuthenticationProviderProps",
		reflect.TypeOf((*UserPoolAuthenticationProviderProps)(nil)).Elem(),
	)
}
