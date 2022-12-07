package awsamplify

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.App",
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
			j := jsiiProxy_App{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApp)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.AppProps",
		reflect.TypeOf((*AppProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.AutoBranchCreation",
		reflect.TypeOf((*AutoBranchCreation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.BasicAuth",
		reflect.TypeOf((*BasicAuth)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_BasicAuth{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.BasicAuthConfig",
		reflect.TypeOf((*BasicAuthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.BasicAuthProps",
		reflect.TypeOf((*BasicAuthProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.Branch",
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
			j := jsiiProxy_Branch{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBranch)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.BranchOptions",
		reflect.TypeOf((*BranchOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.BranchProps",
		reflect.TypeOf((*BranchProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.CfnApp",
		reflect.TypeOf((*CfnApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessToken", GoGetter: "AccessToken"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAppId", GoGetter: "AttrAppId"},
			_jsii_.MemberProperty{JsiiProperty: "attrAppName", GoGetter: "AttrAppName"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultDomain", GoGetter: "AttrDefaultDomain"},
			_jsii_.MemberProperty{JsiiProperty: "autoBranchCreationConfig", GoGetter: "AutoBranchCreationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "basicAuthConfig", GoGetter: "BasicAuthConfig"},
			_jsii_.MemberProperty{JsiiProperty: "buildSpec", GoGetter: "BuildSpec"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customHeaders", GoGetter: "CustomHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "customRules", GoGetter: "CustomRules"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enableBranchAutoDeletion", GoGetter: "EnableBranchAutoDeletion"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVariables", GoGetter: "EnvironmentVariables"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "iamServiceRole", GoGetter: "IamServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oauthToken", GoGetter: "OauthToken"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApp{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnApp.AutoBranchCreationConfigProperty",
		reflect.TypeOf((*CfnApp_AutoBranchCreationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnApp.BasicAuthConfigProperty",
		reflect.TypeOf((*CfnApp_BasicAuthConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnApp.CustomRuleProperty",
		reflect.TypeOf((*CfnApp_CustomRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnApp.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnApp_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnAppProps",
		reflect.TypeOf((*CfnAppProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.CfnBranch",
		reflect.TypeOf((*CfnBranch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appId", GoGetter: "AppId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrBranchName", GoGetter: "AttrBranchName"},
			_jsii_.MemberProperty{JsiiProperty: "basicAuthConfig", GoGetter: "BasicAuthConfig"},
			_jsii_.MemberProperty{JsiiProperty: "branchName", GoGetter: "BranchName"},
			_jsii_.MemberProperty{JsiiProperty: "buildSpec", GoGetter: "BuildSpec"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutoBuild", GoGetter: "EnableAutoBuild"},
			_jsii_.MemberProperty{JsiiProperty: "enablePerformanceMode", GoGetter: "EnablePerformanceMode"},
			_jsii_.MemberProperty{JsiiProperty: "enablePullRequestPreview", GoGetter: "EnablePullRequestPreview"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVariables", GoGetter: "EnvironmentVariables"},
			_jsii_.MemberProperty{JsiiProperty: "framework", GoGetter: "Framework"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "pullRequestEnvironmentName", GoGetter: "PullRequestEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBranch{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnBranch.BasicAuthConfigProperty",
		reflect.TypeOf((*CfnBranch_BasicAuthConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnBranch.EnvironmentVariableProperty",
		reflect.TypeOf((*CfnBranch_EnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnBranchProps",
		reflect.TypeOf((*CfnBranchProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.CfnDomain",
		reflect.TypeOf((*CfnDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appId", GoGetter: "AppId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAutoSubDomainCreationPatterns", GoGetter: "AttrAutoSubDomainCreationPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "attrAutoSubDomainIamRole", GoGetter: "AttrAutoSubDomainIamRole"},
			_jsii_.MemberProperty{JsiiProperty: "attrCertificateRecord", GoGetter: "AttrCertificateRecord"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainName", GoGetter: "AttrDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainStatus", GoGetter: "AttrDomainStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnableAutoSubDomain", GoGetter: "AttrEnableAutoSubDomain"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusReason", GoGetter: "AttrStatusReason"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubDomainCreationPatterns", GoGetter: "AutoSubDomainCreationPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubDomainIamRole", GoGetter: "AutoSubDomainIamRole"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutoSubDomain", GoGetter: "EnableAutoSubDomain"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subDomainSettings", GoGetter: "SubDomainSettings"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnDomain.SubDomainSettingProperty",
		reflect.TypeOf((*CfnDomain_SubDomainSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CfnDomainProps",
		reflect.TypeOf((*CfnDomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.CodeCommitSourceCodeProvider",
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
		"monocdk.aws_amplify.CodeCommitSourceCodeProviderProps",
		reflect.TypeOf((*CodeCommitSourceCodeProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.CustomResponseHeader",
		reflect.TypeOf((*CustomResponseHeader)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.CustomRule",
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
		"monocdk.aws_amplify.CustomRuleOptions",
		reflect.TypeOf((*CustomRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.Domain",
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
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Domain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.DomainOptions",
		reflect.TypeOf((*DomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.DomainProps",
		reflect.TypeOf((*DomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.GitHubSourceCodeProvider",
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
		"monocdk.aws_amplify.GitHubSourceCodeProviderProps",
		reflect.TypeOf((*GitHubSourceCodeProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_amplify.GitLabSourceCodeProvider",
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
		"monocdk.aws_amplify.GitLabSourceCodeProviderProps",
		reflect.TypeOf((*GitLabSourceCodeProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_amplify.IApp",
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
		"monocdk.aws_amplify.IBranch",
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
		"monocdk.aws_amplify.ISourceCodeProvider",
		reflect.TypeOf((*ISourceCodeProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ISourceCodeProvider{}
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_amplify.RedirectStatus",
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
		"monocdk.aws_amplify.SourceCodeProviderConfig",
		reflect.TypeOf((*SourceCodeProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_amplify.SubDomain",
		reflect.TypeOf((*SubDomain)(nil)).Elem(),
	)
}
