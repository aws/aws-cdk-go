// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.App",
		reflect.TypeOf((*App)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAutoBranchEnvironment", GoMethod: "AddAutoBranchEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCustomRule", GoMethod: "AddCustomRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDomain", GoMethod: "AddDomain"},
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "appId", GoGetter: "AppId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "appName", GoGetter: "AppName"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDomain", GoGetter: "DefaultDomain"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_App{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApp)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.AppProps",
		reflect.TypeOf((*AppProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.AutoBranchCreation",
		reflect.TypeOf((*AutoBranchCreation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.BasicAuth",
		reflect.TypeOf((*BasicAuth)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_BasicAuth{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.BasicAuthConfig",
		reflect.TypeOf((*BasicAuthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.BasicAuthProps",
		reflect.TypeOf((*BasicAuthProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.Branch",
		reflect.TypeOf((*Branch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "branchName", GoGetter: "BranchName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Branch{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBranch)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.BranchOptions",
		reflect.TypeOf((*BranchOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.BranchProps",
		reflect.TypeOf((*BranchProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-amplify-alpha.CacheConfigType",
		reflect.TypeOf((*CacheConfigType)(nil)).Elem(),
		map[string]interface{}{
			"AMPLIFY_MANAGED": CacheConfigType_AMPLIFY_MANAGED,
			"AMPLIFY_MANAGED_NO_COOKIES": CacheConfigType_AMPLIFY_MANAGED_NO_COOKIES,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.CodeCommitSourceCodeProvider",
		reflect.TypeOf((*CodeCommitSourceCodeProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CodeCommitSourceCodeProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISourceCodeProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.CodeCommitSourceCodeProviderProps",
		reflect.TypeOf((*CodeCommitSourceCodeProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.CustomResponseHeader",
		reflect.TypeOf((*CustomResponseHeader)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.CustomRule",
		reflect.TypeOf((*CustomRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
		},
		func() interface{} {
			return &jsiiProxy_CustomRule{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.CustomRuleOptions",
		reflect.TypeOf((*CustomRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.Domain",
		reflect.TypeOf((*Domain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "certificateRecord", GoGetter: "CertificateRecord"},
			_jsii_.MemberProperty{JsiiProperty: "domainAutoSubDomainCreationPatterns", GoGetter: "DomainAutoSubDomainCreationPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "domainAutoSubDomainIamRole", GoGetter: "DomainAutoSubDomainIamRole"},
			_jsii_.MemberProperty{JsiiProperty: "domainEnableAutoSubDomain", GoGetter: "DomainEnableAutoSubDomain"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainStatus", GoGetter: "DomainStatus"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "mapRoot", GoMethod: "MapRoot"},
			_jsii_.MemberMethod{JsiiMethod: "mapSubDomain", GoMethod: "MapSubDomain"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Domain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.DomainOptions",
		reflect.TypeOf((*DomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.DomainProps",
		reflect.TypeOf((*DomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.GitHubSourceCodeProvider",
		reflect.TypeOf((*GitHubSourceCodeProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubSourceCodeProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISourceCodeProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.GitHubSourceCodeProviderProps",
		reflect.TypeOf((*GitHubSourceCodeProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-amplify-alpha.GitLabSourceCodeProvider",
		reflect.TypeOf((*GitLabSourceCodeProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_GitLabSourceCodeProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISourceCodeProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.GitLabSourceCodeProviderProps",
		reflect.TypeOf((*GitLabSourceCodeProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-amplify-alpha.IApp",
		reflect.TypeOf((*IApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appId", GoGetter: "AppId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApp{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-amplify-alpha.IBranch",
		reflect.TypeOf((*IBranch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "branchName", GoGetter: "BranchName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBranch{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-amplify-alpha.ISourceCodeProvider",
		reflect.TypeOf((*ISourceCodeProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISourceCodeProvider{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-amplify-alpha.Platform",
		reflect.TypeOf((*Platform)(nil)).Elem(),
		map[string]interface{}{
			"WEB": Platform_WEB,
			"WEB_COMPUTE": Platform_WEB_COMPUTE,
			"WEB_DYNAMIC": Platform_WEB_DYNAMIC,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-amplify-alpha.RedirectStatus",
		reflect.TypeOf((*RedirectStatus)(nil)).Elem(),
		map[string]interface{}{
			"REWRITE": RedirectStatus_REWRITE,
			"PERMANENT_REDIRECT": RedirectStatus_PERMANENT_REDIRECT,
			"TEMPORARY_REDIRECT": RedirectStatus_TEMPORARY_REDIRECT,
			"NOT_FOUND": RedirectStatus_NOT_FOUND,
			"NOT_FOUND_REWRITE": RedirectStatus_NOT_FOUND_REWRITE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.SourceCodeProviderConfig",
		reflect.TypeOf((*SourceCodeProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-amplify-alpha.SubDomain",
		reflect.TypeOf((*SubDomain)(nil)).Elem(),
	)
}
