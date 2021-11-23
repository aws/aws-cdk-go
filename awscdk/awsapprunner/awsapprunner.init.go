package awsapprunner

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.AssetProps",
		reflect.TypeOf((*AssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.AssetSource",
		reflect.TypeOf((*AssetSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AssetSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Source)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.CfnService",
		reflect.TypeOf((*CfnService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceArn", GoGetter: "AttrServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceId", GoGetter: "AttrServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceUrl", GoGetter: "AttrServiceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationArn", GoGetter: "AutoScalingConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfiguration", GoGetter: "EncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckConfiguration", GoGetter: "HealthCheckConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceConfiguration", GoGetter: "InstanceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConfiguration", GoGetter: "SourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnService_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.CodeConfigurationProperty",
		reflect.TypeOf((*CfnService_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.CodeConfigurationValuesProperty",
		reflect.TypeOf((*CfnService_CodeConfigurationValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.CodeRepositoryProperty",
		reflect.TypeOf((*CfnService_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnService_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.HealthCheckConfigurationProperty",
		reflect.TypeOf((*CfnService_HealthCheckConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.ImageConfigurationProperty",
		reflect.TypeOf((*CfnService_ImageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.ImageRepositoryProperty",
		reflect.TypeOf((*CfnService_ImageRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.InstanceConfigurationProperty",
		reflect.TypeOf((*CfnService_InstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.KeyValuePairProperty",
		reflect.TypeOf((*CfnService_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.SourceCodeVersionProperty",
		reflect.TypeOf((*CfnService_SourceCodeVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnService.SourceConfigurationProperty",
		reflect.TypeOf((*CfnService_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CfnServiceProps",
		reflect.TypeOf((*CfnServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CodeConfiguration",
		reflect.TypeOf((*CodeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CodeConfigurationValues",
		reflect.TypeOf((*CodeConfigurationValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.CodeRepositoryProps",
		reflect.TypeOf((*CodeRepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_apprunner.ConfigurationSourceType",
		reflect.TypeOf((*ConfigurationSourceType)(nil)).Elem(),
		map[string]interface{}{
			"API": ConfigurationSourceType_API,
			"REPOSITORY": ConfigurationSourceType_REPOSITORY,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.Cpu",
		reflect.TypeOf((*Cpu)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			return &jsiiProxy_Cpu{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.EcrProps",
		reflect.TypeOf((*EcrProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.EcrPublicProps",
		reflect.TypeOf((*EcrPublicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.EcrPublicSource",
		reflect.TypeOf((*EcrPublicSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_EcrPublicSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Source)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.EcrSource",
		reflect.TypeOf((*EcrSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_EcrSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Source)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.GitHubConnection",
		reflect.TypeOf((*GitHubConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectionArn", GoGetter: "ConnectionArn"},
		},
		func() interface{} {
			return &jsiiProxy_GitHubConnection{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.GithubRepositoryProps",
		reflect.TypeOf((*GithubRepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.GithubSource",
		reflect.TypeOf((*GithubSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_GithubSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Source)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_apprunner.IService",
		reflect.TypeOf((*IService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceArn", GoGetter: "ServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.ImageConfiguration",
		reflect.TypeOf((*ImageConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.ImageRepository",
		reflect.TypeOf((*ImageRepository)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_apprunner.ImageRepositoryType",
		reflect.TypeOf((*ImageRepositoryType)(nil)).Elem(),
		map[string]interface{}{
			"ECR": ImageRepositoryType_ECR,
			"ECR_PUBLIC": ImageRepositoryType_ECR_PUBLIC,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.Memory",
		reflect.TypeOf((*Memory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			return &jsiiProxy_Memory{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
			_jsii_.MemberProperty{JsiiProperty: "serviceArn", GoGetter: "ServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceId", GoGetter: "ServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceStatus", GoGetter: "ServiceStatus"},
			_jsii_.MemberProperty{JsiiProperty: "serviceUrl", GoGetter: "ServiceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Service{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.ServiceAttributes",
		reflect.TypeOf((*ServiceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_apprunner.Source",
		reflect.TypeOf((*Source)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Source{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.SourceCodeVersion",
		reflect.TypeOf((*SourceCodeVersion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_apprunner.SourceConfig",
		reflect.TypeOf((*SourceConfig)(nil)).Elem(),
	)
}
