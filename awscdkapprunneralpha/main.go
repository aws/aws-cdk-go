// The CDK Construct Library for AWS::AppRunner
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
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.AutoScalingConfiguration",
		reflect.TypeOf((*AutoScalingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationArn", GoGetter: "AutoScalingConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationName", GoGetter: "AutoScalingConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationRevision", GoGetter: "AutoScalingConfigurationRevision"},
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
			j := jsiiProxy_AutoScalingConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAutoScalingConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.AutoScalingConfigurationAttributes",
		reflect.TypeOf((*AutoScalingConfigurationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.AutoScalingConfigurationProps",
		reflect.TypeOf((*AutoScalingConfigurationProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@aws-cdk/aws-apprunner-alpha.HealthCheck",
		reflect.TypeOf((*HealthCheck)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckProtocolType", GoGetter: "HealthCheckProtocolType"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThreshold", GoGetter: "HealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThreshold", GoGetter: "UnhealthyThreshold"},
		},
		func() interface{} {
			return &jsiiProxy_HealthCheck{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apprunner-alpha.HealthCheckProtocolType",
		reflect.TypeOf((*HealthCheckProtocolType)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": HealthCheckProtocolType_HTTP,
			"TCP": HealthCheckProtocolType_TCP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.HttpHealthCheckOptions",
		reflect.TypeOf((*HttpHealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apprunner-alpha.IAutoScalingConfiguration",
		reflect.TypeOf((*IAutoScalingConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationArn", GoGetter: "AutoScalingConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationName", GoGetter: "AutoScalingConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingConfigurationRevision", GoGetter: "AutoScalingConfigurationRevision"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAutoScalingConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-apprunner-alpha.IObservabilityConfiguration",
		reflect.TypeOf((*IObservabilityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "observabilityConfigurationArn", GoGetter: "ObservabilityConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "observabilityConfigurationName", GoGetter: "ObservabilityConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "observabilityConfigurationRevision", GoGetter: "ObservabilityConfigurationRevision"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IObservabilityConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
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
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apprunner-alpha.IpAddressType",
		reflect.TypeOf((*IpAddressType)(nil)).Elem(),
		map[string]interface{}{
			"IPV4": IpAddressType_IPV4,
			"DUAL_STACK": IpAddressType_DUAL_STACK,
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
		"@aws-cdk/aws-apprunner-alpha.ObservabilityConfiguration",
		reflect.TypeOf((*ObservabilityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "observabilityConfigurationArn", GoGetter: "ObservabilityConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "observabilityConfigurationName", GoGetter: "ObservabilityConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "observabilityConfigurationRevision", GoGetter: "ObservabilityConfigurationRevision"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ObservabilityConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IObservabilityConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.ObservabilityConfigurationAttributes",
		reflect.TypeOf((*ObservabilityConfigurationAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.ObservabilityConfigurationProps",
		reflect.TypeOf((*ObservabilityConfigurationProps)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
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
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
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
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-apprunner-alpha.TcpHealthCheckOptions",
		reflect.TypeOf((*TcpHealthCheckOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-apprunner-alpha.TraceConfigurationVendor",
		reflect.TypeOf((*TraceConfigurationVendor)(nil)).Elem(),
		map[string]interface{}{
			"AWSXRAY": TraceConfigurationVendor_AWSXRAY,
		},
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
