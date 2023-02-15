package awscdkapprunneralpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.AssetProps",
		reflect.TypeOf((*AssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
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
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.CodeConfiguration",
		reflect.TypeOf((*CodeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.CodeConfigurationValues",
		reflect.TypeOf((*CodeConfigurationValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.CodeRepositoryProps",
		reflect.TypeOf((*CodeRepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apprunner-alpha.ConfigurationSourceType",
		reflect.TypeOf((*ConfigurationSourceType)(nil)).Elem(),
		map[string]interface{}{
			"REPOSITORY": ConfigurationSourceType_REPOSITORY,
			"API": ConfigurationSourceType_API,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		reflect.TypeOf((*Cpu)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			return &jsiiProxy_Cpu{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.EcrProps",
		reflect.TypeOf((*EcrProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicProps",
		reflect.TypeOf((*EcrPublicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
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
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
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
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		reflect.TypeOf((*GitHubConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectionArn", GoGetter: "ConnectionArn"},
		},
		func() interface{} {
			return &jsiiProxy_GitHubConnection{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.GithubRepositoryProps",
		reflect.TypeOf((*GithubRepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
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
		"@aws-cdk/aws-apprunner-alpha.IService",
		reflect.TypeOf((*IService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
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
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apprunner-alpha.IVpcConnector",
		reflect.TypeOf((*IVpcConnector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectorArn", GoGetter: "VpcConnectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectorName", GoGetter: "VpcConnectorName"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectorRevision", GoGetter: "VpcConnectorRevision"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcConnector{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.ImageConfiguration",
		reflect.TypeOf((*ImageConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.ImageRepository",
		reflect.TypeOf((*ImageRepository)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apprunner-alpha.ImageRepositoryType",
		reflect.TypeOf((*ImageRepositoryType)(nil)).Elem(),
		map[string]interface{}{
			"ECR_PUBLIC": ImageRepositoryType_ECR_PUBLIC,
			"ECR": ImageRepositoryType_ECR,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.Memory",
		reflect.TypeOf((*Memory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
		},
		func() interface{} {
			return &jsiiProxy_Memory{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "hasField", GoGetter: "HasField"},
		},
		func() interface{} {
			return &jsiiProxy_Secret{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.SecretVersionInfo",
		reflect.TypeOf((*SecretVersionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnvironmentVariable", GoMethod: "AddEnvironmentVariable"},
			_jsii_.MemberMethod{JsiiMethod: "addSecret", GoMethod: "AddSecret"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceArn", GoGetter: "ServiceArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceId", GoGetter: "ServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceStatus", GoGetter: "ServiceStatus"},
			_jsii_.MemberProperty{JsiiProperty: "serviceUrl", GoGetter: "ServiceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Service{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.ServiceAttributes",
		reflect.TypeOf((*ServiceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.Source",
		reflect.TypeOf((*Source)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Source{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.SourceCodeVersion",
		reflect.TypeOf((*SourceCodeVersion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.SourceConfig",
		reflect.TypeOf((*SourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		reflect.TypeOf((*VpcConnector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectorArn", GoGetter: "VpcConnectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectorName", GoGetter: "VpcConnectorName"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConnectorRevision", GoGetter: "VpcConnectorRevision"},
		},
		func() interface{} {
			j := jsiiProxy_VpcConnector{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcConnector)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.VpcConnectorAttributes",
		reflect.TypeOf((*VpcConnectorAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.VpcConnectorProps",
		reflect.TypeOf((*VpcConnectorProps)(nil)).Elem(),
	)
}
