package awscognitoidentitypool

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_cognito_identitypool.IIdentityPool",
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
		"aws-cdk-lib.aws_cognito_identitypool.IUserPoolAuthenticationProvider",
		reflect.TypeOf((*IUserPoolAuthenticationProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IUserPoolAuthenticationProvider{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPool",
		reflect.TypeOf((*IdentityPool)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "roleAttachment", GoGetter: "RoleAttachment"},
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
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolAmazonLoginProvider",
		reflect.TypeOf((*IdentityPoolAmazonLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolAppleLoginProvider",
		reflect.TypeOf((*IdentityPoolAppleLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolAuthenticationProviders",
		reflect.TypeOf((*IdentityPoolAuthenticationProviders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolFacebookLoginProvider",
		reflect.TypeOf((*IdentityPoolFacebookLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolGoogleLoginProvider",
		reflect.TypeOf((*IdentityPoolGoogleLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolProps",
		reflect.TypeOf((*IdentityPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolProviderType",
		reflect.TypeOf((*IdentityPoolProviderType)(nil)).Elem(),
		map[string]interface{}{
			"FACEBOOK": IdentityPoolProviderType_FACEBOOK,
			"GOOGLE": IdentityPoolProviderType_GOOGLE,
			"AMAZON": IdentityPoolProviderType_AMAZON,
			"APPLE": IdentityPoolProviderType_APPLE,
			"TWITTER": IdentityPoolProviderType_TWITTER,
			"OPEN_ID": IdentityPoolProviderType_OPEN_ID,
			"SAML": IdentityPoolProviderType_SAML,
			"USER_POOL": IdentityPoolProviderType_USER_POOL,
			"CUSTOM": IdentityPoolProviderType_CUSTOM,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolProviderUrl",
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
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolRoleMapping",
		reflect.TypeOf((*IdentityPoolRoleMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.IdentityPoolTwitterLoginProvider",
		reflect.TypeOf((*IdentityPoolTwitterLoginProvider)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cognito_identitypool.RoleMappingMatchType",
		reflect.TypeOf((*RoleMappingMatchType)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": RoleMappingMatchType_EQUALS,
			"CONTAINS": RoleMappingMatchType_CONTAINS,
			"STARTS_WITH": RoleMappingMatchType_STARTS_WITH,
			"NOTEQUAL": RoleMappingMatchType_NOTEQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.RoleMappingRule",
		reflect.TypeOf((*RoleMappingRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cognito_identitypool.UserPoolAuthenticationProvider",
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
		"aws-cdk-lib.aws_cognito_identitypool.UserPoolAuthenticationProviderBindConfig",
		reflect.TypeOf((*UserPoolAuthenticationProviderBindConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.UserPoolAuthenticationProviderBindOptions",
		reflect.TypeOf((*UserPoolAuthenticationProviderBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cognito_identitypool.UserPoolAuthenticationProviderProps",
		reflect.TypeOf((*UserPoolAuthenticationProviderProps)(nil)).Elem(),
	)
}
