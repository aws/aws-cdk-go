package awselasticloadbalancingv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AddApplicationActionProps",
		reflect.TypeOf((*AddApplicationActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AddApplicationTargetGroupsProps",
		reflect.TypeOf((*AddApplicationTargetGroupsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AddApplicationTargetsProps",
		reflect.TypeOf((*AddApplicationTargetsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AddNetworkActionProps",
		reflect.TypeOf((*AddNetworkActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AddNetworkTargetsProps",
		reflect.TypeOf((*AddNetworkTargetsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AddRuleProps",
		reflect.TypeOf((*AddRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AlpnPolicy",
		reflect.TypeOf((*AlpnPolicy)(nil)).Elem(),
		map[string]interface{}{
			"HTTP1_ONLY": AlpnPolicy_HTTP1_ONLY,
			"HTTP2_ONLY": AlpnPolicy_HTTP2_ONLY,
			"HTTP2_OPTIONAL": AlpnPolicy_HTTP2_OPTIONAL,
			"HTTP2_PREFERRED": AlpnPolicy_HTTP2_PREFERRED,
			"NONE": AlpnPolicy_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		reflect.TypeOf((*ApplicationListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "addCertificates", GoMethod: "AddCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "addTargetGroups", GoMethod: "AddTargetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addTargets", GoMethod: "AddTargets"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "registerConnectable", GoMethod: "RegisterConnectable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateListener", GoMethod: "ValidateListener"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationListener{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseListener)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplicationListener)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerAttributes",
		reflect.TypeOf((*ApplicationListenerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
		reflect.TypeOf((*ApplicationListenerCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationListenerCertificate{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificateProps",
		reflect.TypeOf((*ApplicationListenerCertificateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerLookupOptions",
		reflect.TypeOf((*ApplicationListenerLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerProps",
		reflect.TypeOf((*ApplicationListenerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		reflect.TypeOf((*ApplicationListenerRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCondition", GoMethod: "AddCondition"},
			_jsii_.MemberMethod{JsiiMethod: "configureAction", GoMethod: "ConfigureAction"},
			_jsii_.MemberProperty{JsiiProperty: "listenerRuleArn", GoGetter: "ListenerRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationListenerRule{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRuleProps",
		reflect.TypeOf((*ApplicationListenerRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		reflect.TypeOf((*ApplicationLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addListener", GoMethod: "AddListener"},
			_jsii_.MemberMethod{JsiiMethod: "addRedirect", GoMethod: "AddRedirect"},
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "listeners", GoGetter: "Listeners"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCanonicalHostedZoneId", GoGetter: "LoadBalancerCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerDnsName", GoGetter: "LoadBalancerDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerFullName", GoGetter: "LoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerName", GoGetter: "LoadBalancerName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerSecurityGroups", GoGetter: "LoadBalancerSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "logAccessLogs", GoMethod: "LogAccessLogs"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveConnectionCount", GoMethod: "MetricActiveConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricClientTlsNegotiationErrorCount", GoMethod: "MetricClientTlsNegotiationErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumedLCUs", GoMethod: "MetricConsumedLCUs"},
			_jsii_.MemberMethod{JsiiMethod: "metricElbAuthError", GoMethod: "MetricElbAuthError"},
			_jsii_.MemberMethod{JsiiMethod: "metricElbAuthFailure", GoMethod: "MetricElbAuthFailure"},
			_jsii_.MemberMethod{JsiiMethod: "metricElbAuthLatency", GoMethod: "MetricElbAuthLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricElbAuthSuccess", GoMethod: "MetricElbAuthSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "metricHttpCodeElb", GoMethod: "MetricHttpCodeElb"},
			_jsii_.MemberMethod{JsiiMethod: "metricHttpCodeTarget", GoMethod: "MetricHttpCodeTarget"},
			_jsii_.MemberMethod{JsiiMethod: "metricHttpFixedResponseCount", GoMethod: "MetricHttpFixedResponseCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHttpRedirectCount", GoMethod: "MetricHttpRedirectCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHttpRedirectUrlLimitExceededCount", GoMethod: "MetricHttpRedirectUrlLimitExceededCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIpv6ProcessedBytes", GoMethod: "MetricIpv6ProcessedBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIpv6RequestCount", GoMethod: "MetricIpv6RequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricNewConnectionCount", GoMethod: "MetricNewConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricProcessedBytes", GoMethod: "MetricProcessedBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricRejectedConnectionCount", GoMethod: "MetricRejectedConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRequestCount", GoMethod: "MetricRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluations", GoMethod: "MetricRuleEvaluations"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetConnectionErrorCount", GoMethod: "MetricTargetConnectionErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetResponseTime", GoMethod: "MetricTargetResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetTLSNegotiationErrorCount", GoMethod: "MetricTargetTLSNegotiationErrorCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "removeAttribute", GoMethod: "RemoveAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resourcePolicyPrincipal", GoMethod: "ResourcePolicyPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "setAttribute", GoMethod: "SetAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateLoadBalancer", GoMethod: "ValidateLoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseLoadBalancer)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplicationLoadBalancer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancerAttributes",
		reflect.TypeOf((*ApplicationLoadBalancerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancerLookupOptions",
		reflect.TypeOf((*ApplicationLoadBalancerLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancerProps",
		reflect.TypeOf((*ApplicationLoadBalancerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancerRedirectConfig",
		reflect.TypeOf((*ApplicationLoadBalancerRedirectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationProtocol",
		reflect.TypeOf((*ApplicationProtocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": ApplicationProtocol_HTTP,
			"HTTPS": ApplicationProtocol_HTTPS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationProtocolVersion",
		reflect.TypeOf((*ApplicationProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"GRPC": ApplicationProtocolVersion_GRPC,
			"HTTP1": ApplicationProtocolVersion_HTTP1,
			"HTTP2": ApplicationProtocolVersion_HTTP2,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationTargetGroup",
		reflect.TypeOf((*ApplicationTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLoadBalancerTarget", GoMethod: "AddLoadBalancerTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberMethod{JsiiMethod: "configureHealthCheck", GoMethod: "ConfigureHealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPort", GoGetter: "DefaultPort"},
			_jsii_.MemberMethod{JsiiMethod: "enableCookieStickiness", GoMethod: "EnableCookieStickiness"},
			_jsii_.MemberProperty{JsiiProperty: "firstLoadBalancerFullName", GoGetter: "FirstLoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheck", GoGetter: "HealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArns", GoGetter: "LoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttached", GoGetter: "LoadBalancerAttached"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttachedDependencies", GoGetter: "LoadBalancerAttachedDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricHealthyHostCount", GoMethod: "MetricHealthyHostCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricHttpCodeTarget", GoMethod: "MetricHttpCodeTarget"},
			_jsii_.MemberMethod{JsiiMethod: "metricIpv6RequestCount", GoMethod: "MetricIpv6RequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRequestCount", GoMethod: "MetricRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricRequestCountPerTarget", GoMethod: "MetricRequestCountPerTarget"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetConnectionErrorCount", GoMethod: "MetricTargetConnectionErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetResponseTime", GoMethod: "MetricTargetResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTargetTLSNegotiationErrorCount", GoMethod: "MetricTargetTLSNegotiationErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricUnhealthyHostCount", GoMethod: "MetricUnhealthyHostCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerConnectable", GoMethod: "RegisterConnectable"},
			_jsii_.MemberMethod{JsiiMethod: "registerListener", GoMethod: "RegisterListener"},
			_jsii_.MemberMethod{JsiiMethod: "setAttribute", GoMethod: "SetAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupFullName", GoGetter: "TargetGroupFullName"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupLoadBalancerArns", GoGetter: "TargetGroupLoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateHealthCheck", GoMethod: "ValidateHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "validateTargetGroup", GoMethod: "ValidateTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationTargetGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TargetGroupBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApplicationTargetGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationTargetGroupProps",
		reflect.TypeOf((*ApplicationTargetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.AuthenticateOidcOptions",
		reflect.TypeOf((*AuthenticateOidcOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseApplicationListenerProps",
		reflect.TypeOf((*BaseApplicationListenerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseApplicationListenerRuleProps",
		reflect.TypeOf((*BaseApplicationListenerRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseListener",
		reflect.TypeOf((*BaseListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateListener", GoMethod: "ValidateListener"},
		},
		func() interface{} {
			j := jsiiProxy_BaseListener{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IListener)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseListenerLookupOptions",
		reflect.TypeOf((*BaseListenerLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		reflect.TypeOf((*BaseLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCanonicalHostedZoneId", GoGetter: "LoadBalancerCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerDnsName", GoGetter: "LoadBalancerDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerFullName", GoGetter: "LoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerName", GoGetter: "LoadBalancerName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerSecurityGroups", GoGetter: "LoadBalancerSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "logAccessLogs", GoMethod: "LogAccessLogs"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "removeAttribute", GoMethod: "RemoveAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resourcePolicyPrincipal", GoMethod: "ResourcePolicyPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "setAttribute", GoMethod: "SetAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateLoadBalancer", GoMethod: "ValidateLoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_BaseLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancerLookupOptions",
		reflect.TypeOf((*BaseLoadBalancerLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancerProps",
		reflect.TypeOf((*BaseLoadBalancerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseNetworkListenerProps",
		reflect.TypeOf((*BaseNetworkListenerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseTargetGroupProps",
		reflect.TypeOf((*BaseTargetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		reflect.TypeOf((*CfnListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alpnPolicy", GoGetter: "AlpnPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrListenerArn", GoGetter: "AttrListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "certificates", GoGetter: "Certificates"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultActions", GoGetter: "DefaultActions"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mutualAuthentication", GoGetter: "MutualAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sslPolicy", GoGetter: "SslPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListener{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.ActionProperty",
		reflect.TypeOf((*CfnListener_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.AuthenticateCognitoConfigProperty",
		reflect.TypeOf((*CfnListener_AuthenticateCognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.AuthenticateOidcConfigProperty",
		reflect.TypeOf((*CfnListener_AuthenticateOidcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.CertificateProperty",
		reflect.TypeOf((*CfnListener_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.FixedResponseConfigProperty",
		reflect.TypeOf((*CfnListener_FixedResponseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.ForwardConfigProperty",
		reflect.TypeOf((*CfnListener_ForwardConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.MutualAuthenticationProperty",
		reflect.TypeOf((*CfnListener_MutualAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.RedirectConfigProperty",
		reflect.TypeOf((*CfnListener_RedirectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.TargetGroupStickinessConfigProperty",
		reflect.TypeOf((*CfnListener_TargetGroupStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener.TargetGroupTupleProperty",
		reflect.TypeOf((*CfnListener_TargetGroupTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		reflect.TypeOf((*CfnListenerCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "certificates", GoGetter: "Certificates"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerCertificate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate.CertificateProperty",
		reflect.TypeOf((*CfnListenerCertificate_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificateProps",
		reflect.TypeOf((*CfnListenerCertificateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerProps",
		reflect.TypeOf((*CfnListenerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		reflect.TypeOf((*CfnListenerRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefault", GoGetter: "AttrIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "attrRuleArn", GoGetter: "AttrRuleArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "conditions", GoGetter: "Conditions"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnListenerRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.ActionProperty",
		reflect.TypeOf((*CfnListenerRule_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.AuthenticateCognitoConfigProperty",
		reflect.TypeOf((*CfnListenerRule_AuthenticateCognitoConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.AuthenticateOidcConfigProperty",
		reflect.TypeOf((*CfnListenerRule_AuthenticateOidcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.FixedResponseConfigProperty",
		reflect.TypeOf((*CfnListenerRule_FixedResponseConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.ForwardConfigProperty",
		reflect.TypeOf((*CfnListenerRule_ForwardConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.HostHeaderConfigProperty",
		reflect.TypeOf((*CfnListenerRule_HostHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.HttpHeaderConfigProperty",
		reflect.TypeOf((*CfnListenerRule_HttpHeaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.HttpRequestMethodConfigProperty",
		reflect.TypeOf((*CfnListenerRule_HttpRequestMethodConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.PathPatternConfigProperty",
		reflect.TypeOf((*CfnListenerRule_PathPatternConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.QueryStringConfigProperty",
		reflect.TypeOf((*CfnListenerRule_QueryStringConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.QueryStringKeyValueProperty",
		reflect.TypeOf((*CfnListenerRule_QueryStringKeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.RedirectConfigProperty",
		reflect.TypeOf((*CfnListenerRule_RedirectConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.RuleConditionProperty",
		reflect.TypeOf((*CfnListenerRule_RuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.SourceIpConfigProperty",
		reflect.TypeOf((*CfnListenerRule_SourceIpConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.TargetGroupStickinessConfigProperty",
		reflect.TypeOf((*CfnListenerRule_TargetGroupStickinessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule.TargetGroupTupleProperty",
		reflect.TypeOf((*CfnListenerRule_TargetGroupTupleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRuleProps",
		reflect.TypeOf((*CfnListenerRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		reflect.TypeOf((*CfnLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCanonicalHostedZoneId", GoGetter: "AttrCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "attrDnsName", GoGetter: "AttrDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoadBalancerArn", GoGetter: "AttrLoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoadBalancerFullName", GoGetter: "AttrLoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoadBalancerName", GoGetter: "AttrLoadBalancerName"},
			_jsii_.MemberProperty{JsiiProperty: "attrSecurityGroups", GoGetter: "AttrSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enforceSecurityGroupInboundRulesOnPrivateLinkTraffic", GoGetter: "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttributes", GoGetter: "LoadBalancerAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetMappings", GoGetter: "SubnetMappings"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer.LoadBalancerAttributeProperty",
		reflect.TypeOf((*CfnLoadBalancer_LoadBalancerAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer.SubnetMappingProperty",
		reflect.TypeOf((*CfnLoadBalancer_SubnetMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancerProps",
		reflect.TypeOf((*CfnLoadBalancerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		reflect.TypeOf((*CfnTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrLoadBalancerArns", GoGetter: "AttrLoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "attrTargetGroupArn", GoGetter: "AttrTargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrTargetGroupFullName", GoGetter: "AttrTargetGroupFullName"},
			_jsii_.MemberProperty{JsiiProperty: "attrTargetGroupName", GoGetter: "AttrTargetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckEnabled", GoGetter: "HealthCheckEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckIntervalSeconds", GoGetter: "HealthCheckIntervalSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckPath", GoGetter: "HealthCheckPath"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckPort", GoGetter: "HealthCheckPort"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckProtocol", GoGetter: "HealthCheckProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTimeoutSeconds", GoGetter: "HealthCheckTimeoutSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThresholdCount", GoGetter: "HealthyThresholdCount"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "matcher", GoGetter: "Matcher"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolVersion", GoGetter: "ProtocolVersion"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupAttributes", GoGetter: "TargetGroupAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targets", GoGetter: "Targets"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThresholdCount", GoGetter: "UnhealthyThresholdCount"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTargetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup.MatcherProperty",
		reflect.TypeOf((*CfnTargetGroup_MatcherProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup.TargetDescriptionProperty",
		reflect.TypeOf((*CfnTargetGroup_TargetDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup.TargetGroupAttributeProperty",
		reflect.TypeOf((*CfnTargetGroup_TargetGroupAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroupProps",
		reflect.TypeOf((*CfnTargetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTrustStore",
		reflect.TypeOf((*CfnTrustStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrNumberOfCaCertificates", GoGetter: "AttrNumberOfCaCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrTrustStoreArn", GoGetter: "AttrTrustStoreArn"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificatesBundleS3Bucket", GoGetter: "CaCertificatesBundleS3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificatesBundleS3Key", GoGetter: "CaCertificatesBundleS3Key"},
			_jsii_.MemberProperty{JsiiProperty: "caCertificatesBundleS3ObjectVersion", GoGetter: "CaCertificatesBundleS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTrustStoreProps",
		reflect.TypeOf((*CfnTrustStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTrustStoreRevocation",
		reflect.TypeOf((*CfnTrustStoreRevocation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrRevocationId", GoGetter: "AttrRevocationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTrustStoreRevocations", GoGetter: "AttrTrustStoreRevocations"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "revocationContents", GoGetter: "RevocationContents"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustStoreArn", GoGetter: "TrustStoreArn"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrustStoreRevocation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTrustStoreRevocation.RevocationContentProperty",
		reflect.TypeOf((*CfnTrustStoreRevocation_RevocationContentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTrustStoreRevocation.TrustStoreRevocationProperty",
		reflect.TypeOf((*CfnTrustStoreRevocation_TrustStoreRevocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTrustStoreRevocationProps",
		reflect.TypeOf((*CfnTrustStoreRevocationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ClientRoutingPolicy",
		reflect.TypeOf((*ClientRoutingPolicy)(nil)).Elem(),
		map[string]interface{}{
			"AVAILABILITY_ZONE_AFFINITY": ClientRoutingPolicy_AVAILABILITY_ZONE_AFFINITY,
			"PARTIAL_AVAILABILITY_ZONE_AFFINITY": ClientRoutingPolicy_PARTIAL_AVAILABILITY_ZONE_AFFINITY,
			"ANY_AVAILABILITY_ZONE": ClientRoutingPolicy_ANY_AVAILABILITY_ZONE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.DesyncMitigationMode",
		reflect.TypeOf((*DesyncMitigationMode)(nil)).Elem(),
		map[string]interface{}{
			"MONITOR": DesyncMitigationMode_MONITOR,
			"DEFENSIVE": DesyncMitigationMode_DEFENSIVE,
			"STRICTEST": DesyncMitigationMode_STRICTEST,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.FixedResponseOptions",
		reflect.TypeOf((*FixedResponseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ForwardOptions",
		reflect.TypeOf((*ForwardOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.HealthCheck",
		reflect.TypeOf((*HealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.HttpCodeElb",
		reflect.TypeOf((*HttpCodeElb)(nil)).Elem(),
		map[string]interface{}{
			"ELB_3XX_COUNT": HttpCodeElb_ELB_3XX_COUNT,
			"ELB_4XX_COUNT": HttpCodeElb_ELB_4XX_COUNT,
			"ELB_5XX_COUNT": HttpCodeElb_ELB_5XX_COUNT,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.HttpCodeTarget",
		reflect.TypeOf((*HttpCodeTarget)(nil)).Elem(),
		map[string]interface{}{
			"TARGET_2XX_COUNT": HttpCodeTarget_TARGET_2XX_COUNT,
			"TARGET_3XX_COUNT": HttpCodeTarget_TARGET_3XX_COUNT,
			"TARGET_4XX_COUNT": HttpCodeTarget_TARGET_4XX_COUNT,
			"TARGET_5XX_COUNT": HttpCodeTarget_TARGET_5XX_COUNT,
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IApplicationListener",
		reflect.TypeOf((*IApplicationListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "addCertificates", GoMethod: "AddCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "addTargetGroups", GoMethod: "AddTargetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addTargets", GoMethod: "AddTargets"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerConnectable", GoMethod: "RegisterConnectable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IApplicationListener{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IListener)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IApplicationLoadBalancer",
		reflect.TypeOf((*IApplicationLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addListener", GoMethod: "AddListener"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "listeners", GoGetter: "Listeners"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCanonicalHostedZoneId", GoGetter: "LoadBalancerCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerDnsName", GoGetter: "LoadBalancerDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_IApplicationLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILoadBalancerV2)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IApplicationLoadBalancerMetrics",
		reflect.TypeOf((*IApplicationLoadBalancerMetrics)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activeConnectionCount", GoMethod: "ActiveConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "clientTlsNegotiationErrorCount", GoMethod: "ClientTlsNegotiationErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "consumedLCUs", GoMethod: "ConsumedLCUs"},
			_jsii_.MemberMethod{JsiiMethod: "custom", GoMethod: "Custom"},
			_jsii_.MemberMethod{JsiiMethod: "elbAuthError", GoMethod: "ElbAuthError"},
			_jsii_.MemberMethod{JsiiMethod: "elbAuthFailure", GoMethod: "ElbAuthFailure"},
			_jsii_.MemberMethod{JsiiMethod: "elbAuthLatency", GoMethod: "ElbAuthLatency"},
			_jsii_.MemberMethod{JsiiMethod: "elbAuthSuccess", GoMethod: "ElbAuthSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "httpCodeElb", GoMethod: "HttpCodeElb"},
			_jsii_.MemberMethod{JsiiMethod: "httpCodeTarget", GoMethod: "HttpCodeTarget"},
			_jsii_.MemberMethod{JsiiMethod: "httpFixedResponseCount", GoMethod: "HttpFixedResponseCount"},
			_jsii_.MemberMethod{JsiiMethod: "httpRedirectCount", GoMethod: "HttpRedirectCount"},
			_jsii_.MemberMethod{JsiiMethod: "httpRedirectUrlLimitExceededCount", GoMethod: "HttpRedirectUrlLimitExceededCount"},
			_jsii_.MemberMethod{JsiiMethod: "ipv6ProcessedBytes", GoMethod: "Ipv6ProcessedBytes"},
			_jsii_.MemberMethod{JsiiMethod: "ipv6RequestCount", GoMethod: "Ipv6RequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "newConnectionCount", GoMethod: "NewConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "processedBytes", GoMethod: "ProcessedBytes"},
			_jsii_.MemberMethod{JsiiMethod: "rejectedConnectionCount", GoMethod: "RejectedConnectionCount"},
			_jsii_.MemberMethod{JsiiMethod: "requestCount", GoMethod: "RequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "ruleEvaluations", GoMethod: "RuleEvaluations"},
			_jsii_.MemberMethod{JsiiMethod: "targetConnectionErrorCount", GoMethod: "TargetConnectionErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "targetResponseTime", GoMethod: "TargetResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "targetTLSNegotiationErrorCount", GoMethod: "TargetTLSNegotiationErrorCount"},
		},
		func() interface{} {
			return &jsiiProxy_IApplicationLoadBalancerMetrics{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IApplicationLoadBalancerTarget",
		reflect.TypeOf((*IApplicationLoadBalancerTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToApplicationTargetGroup", GoMethod: "AttachToApplicationTargetGroup"},
		},
		func() interface{} {
			return &jsiiProxy_IApplicationLoadBalancerTarget{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IApplicationTargetGroup",
		reflect.TypeOf((*IApplicationTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArns", GoGetter: "LoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttached", GoGetter: "LoadBalancerAttached"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerConnectable", GoMethod: "RegisterConnectable"},
			_jsii_.MemberMethod{JsiiMethod: "registerListener", GoMethod: "RegisterListener"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
		},
		func() interface{} {
			j := jsiiProxy_IApplicationTargetGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITargetGroup)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IApplicationTargetGroupMetrics",
		reflect.TypeOf((*IApplicationTargetGroupMetrics)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "custom", GoMethod: "Custom"},
			_jsii_.MemberMethod{JsiiMethod: "healthyHostCount", GoMethod: "HealthyHostCount"},
			_jsii_.MemberMethod{JsiiMethod: "httpCodeTarget", GoMethod: "HttpCodeTarget"},
			_jsii_.MemberMethod{JsiiMethod: "ipv6RequestCount", GoMethod: "Ipv6RequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "requestCount", GoMethod: "RequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "requestCountPerTarget", GoMethod: "RequestCountPerTarget"},
			_jsii_.MemberMethod{JsiiMethod: "targetConnectionErrorCount", GoMethod: "TargetConnectionErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "targetResponseTime", GoMethod: "TargetResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "targetTLSNegotiationErrorCount", GoMethod: "TargetTLSNegotiationErrorCount"},
			_jsii_.MemberMethod{JsiiMethod: "unhealthyHostCount", GoMethod: "UnhealthyHostCount"},
		},
		func() interface{} {
			return &jsiiProxy_IApplicationTargetGroupMetrics{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IListener",
		reflect.TypeOf((*IListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IListener{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IListenerAction",
		reflect.TypeOf((*IListenerAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderActions", GoMethod: "RenderActions"},
			_jsii_.MemberMethod{JsiiMethod: "renderRuleActions", GoMethod: "RenderRuleActions"},
		},
		func() interface{} {
			return &jsiiProxy_IListenerAction{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IListenerCertificate",
		reflect.TypeOf((*IListenerCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
		},
		func() interface{} {
			return &jsiiProxy_IListenerCertificate{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ILoadBalancerV2",
		reflect.TypeOf((*ILoadBalancerV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCanonicalHostedZoneId", GoGetter: "LoadBalancerCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerDnsName", GoGetter: "LoadBalancerDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ILoadBalancerV2{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.INetworkListener",
		reflect.TypeOf((*INetworkListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_INetworkListener{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IListener)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.INetworkLoadBalancer",
		reflect.TypeOf((*INetworkLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addListener", GoMethod: "AddListener"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enforceSecurityGroupInboundRulesOnPrivateLinkTraffic", GoGetter: "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCanonicalHostedZoneId", GoGetter: "LoadBalancerCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerDnsName", GoGetter: "LoadBalancerDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_INetworkLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILoadBalancerV2)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IVpcEndpointServiceLoadBalancer)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.INetworkLoadBalancerMetrics",
		reflect.TypeOf((*INetworkLoadBalancerMetrics)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "activeFlowCount", GoMethod: "ActiveFlowCount"},
			_jsii_.MemberMethod{JsiiMethod: "consumedLCUs", GoMethod: "ConsumedLCUs"},
			_jsii_.MemberMethod{JsiiMethod: "custom", GoMethod: "Custom"},
			_jsii_.MemberMethod{JsiiMethod: "newFlowCount", GoMethod: "NewFlowCount"},
			_jsii_.MemberMethod{JsiiMethod: "processedBytes", GoMethod: "ProcessedBytes"},
			_jsii_.MemberMethod{JsiiMethod: "tcpClientResetCount", GoMethod: "TcpClientResetCount"},
			_jsii_.MemberMethod{JsiiMethod: "tcpElbResetCount", GoMethod: "TcpElbResetCount"},
			_jsii_.MemberMethod{JsiiMethod: "tcpTargetResetCount", GoMethod: "TcpTargetResetCount"},
		},
		func() interface{} {
			return &jsiiProxy_INetworkLoadBalancerMetrics{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.INetworkLoadBalancerTarget",
		reflect.TypeOf((*INetworkLoadBalancerTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "attachToNetworkTargetGroup", GoMethod: "AttachToNetworkTargetGroup"},
		},
		func() interface{} {
			return &jsiiProxy_INetworkLoadBalancerTarget{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.INetworkTargetGroup",
		reflect.TypeOf((*INetworkTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArns", GoGetter: "LoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttached", GoGetter: "LoadBalancerAttached"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerListener", GoMethod: "RegisterListener"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
		},
		func() interface{} {
			j := jsiiProxy_INetworkTargetGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITargetGroup)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.INetworkTargetGroupMetrics",
		reflect.TypeOf((*INetworkTargetGroupMetrics)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "custom", GoMethod: "Custom"},
			_jsii_.MemberMethod{JsiiMethod: "healthyHostCount", GoMethod: "HealthyHostCount"},
			_jsii_.MemberMethod{JsiiMethod: "unHealthyHostCount", GoMethod: "UnHealthyHostCount"},
		},
		func() interface{} {
			return &jsiiProxy_INetworkTargetGroupMetrics{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ITargetGroup",
		reflect.TypeOf((*ITargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArns", GoGetter: "LoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttached", GoGetter: "LoadBalancerAttached"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
		},
		func() interface{} {
			j := jsiiProxy_ITargetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.IpAddressType",
		reflect.TypeOf((*IpAddressType)(nil)).Elem(),
		map[string]interface{}{
			"IPV4": IpAddressType_IPV4,
			"DUAL_STACK": IpAddressType_DUAL_STACK,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		reflect.TypeOf((*ListenerAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRuleAction", GoMethod: "AddRuleAction"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "next", GoGetter: "Next"},
			_jsii_.MemberMethod{JsiiMethod: "renderActions", GoMethod: "RenderActions"},
			_jsii_.MemberMethod{JsiiMethod: "renderRuleActions", GoMethod: "RenderRuleActions"},
			_jsii_.MemberMethod{JsiiMethod: "renumber", GoMethod: "Renumber"},
		},
		func() interface{} {
			j := jsiiProxy_ListenerAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IListenerAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		reflect.TypeOf((*ListenerCertificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
		},
		func() interface{} {
			j := jsiiProxy_ListenerCertificate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IListenerCertificate)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		reflect.TypeOf((*ListenerCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderRawCondition", GoMethod: "RenderRawCondition"},
		},
		func() interface{} {
			return &jsiiProxy_ListenerCondition{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.LoadBalancerTargetProps",
		reflect.TypeOf((*LoadBalancerTargetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkForwardOptions",
		reflect.TypeOf((*NetworkForwardOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		reflect.TypeOf((*NetworkListener)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAction", GoMethod: "AddAction"},
			_jsii_.MemberMethod{JsiiMethod: "addCertificates", GoMethod: "AddCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "addTargetGroups", GoMethod: "AddTargetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addTargets", GoMethod: "AddTargets"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArn", GoGetter: "ListenerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateListener", GoMethod: "ValidateListener"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkListener{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseListener)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkListener)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		reflect.TypeOf((*NetworkListenerAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "next", GoGetter: "Next"},
			_jsii_.MemberMethod{JsiiMethod: "renderActions", GoMethod: "RenderActions"},
			_jsii_.MemberMethod{JsiiMethod: "renderRuleActions", GoMethod: "RenderRuleActions"},
			_jsii_.MemberMethod{JsiiMethod: "renumber", GoMethod: "Renumber"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkListenerAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IListenerAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerLookupOptions",
		reflect.TypeOf((*NetworkListenerLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerProps",
		reflect.TypeOf((*NetworkListenerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		reflect.TypeOf((*NetworkLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addListener", GoMethod: "AddListener"},
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enforceSecurityGroupInboundRulesOnPrivateLinkTraffic", GoGetter: "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressType", GoGetter: "IpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCanonicalHostedZoneId", GoGetter: "LoadBalancerCanonicalHostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerDnsName", GoGetter: "LoadBalancerDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerFullName", GoGetter: "LoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerName", GoGetter: "LoadBalancerName"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerSecurityGroups", GoGetter: "LoadBalancerSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "logAccessLogs", GoMethod: "LogAccessLogs"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveFlowCount", GoMethod: "MetricActiveFlowCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricConsumedLCUs", GoMethod: "MetricConsumedLCUs"},
			_jsii_.MemberMethod{JsiiMethod: "metricNewFlowCount", GoMethod: "MetricNewFlowCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricProcessedBytes", GoMethod: "MetricProcessedBytes"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberMethod{JsiiMethod: "metricTcpClientResetCount", GoMethod: "MetricTcpClientResetCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTcpElbResetCount", GoMethod: "MetricTcpElbResetCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricTcpTargetResetCount", GoMethod: "MetricTcpTargetResetCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "removeAttribute", GoMethod: "RemoveAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resourcePolicyPrincipal", GoMethod: "ResourcePolicyPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "setAttribute", GoMethod: "SetAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateLoadBalancer", GoMethod: "ValidateLoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkLoadBalancer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseLoadBalancer)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkLoadBalancer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancerAttributes",
		reflect.TypeOf((*NetworkLoadBalancerAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancerLookupOptions",
		reflect.TypeOf((*NetworkLoadBalancerLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancerProps",
		reflect.TypeOf((*NetworkLoadBalancerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		reflect.TypeOf((*NetworkTargetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLoadBalancerTarget", GoMethod: "AddLoadBalancerTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addTarget", GoMethod: "AddTarget"},
			_jsii_.MemberMethod{JsiiMethod: "configureHealthCheck", GoMethod: "ConfigureHealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPort", GoGetter: "DefaultPort"},
			_jsii_.MemberProperty{JsiiProperty: "firstLoadBalancerFullName", GoGetter: "FirstLoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheck", GoGetter: "HealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArns", GoGetter: "LoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttached", GoGetter: "LoadBalancerAttached"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttachedDependencies", GoGetter: "LoadBalancerAttachedDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "metricHealthyHostCount", GoMethod: "MetricHealthyHostCount"},
			_jsii_.MemberProperty{JsiiProperty: "metrics", GoGetter: "Metrics"},
			_jsii_.MemberMethod{JsiiMethod: "metricUnHealthyHostCount", GoMethod: "MetricUnHealthyHostCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "registerListener", GoMethod: "RegisterListener"},
			_jsii_.MemberMethod{JsiiMethod: "setAttribute", GoMethod: "SetAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupFullName", GoGetter: "TargetGroupFullName"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupLoadBalancerArns", GoGetter: "TargetGroupLoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateHealthCheck", GoMethod: "ValidateHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "validateTargetGroup", GoMethod: "ValidateTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkTargetGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TargetGroupBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkTargetGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroupProps",
		reflect.TypeOf((*NetworkTargetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkWeightedTargetGroup",
		reflect.TypeOf((*NetworkWeightedTargetGroup)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": Protocol_HTTP,
			"HTTPS": Protocol_HTTPS,
			"TCP": Protocol_TCP,
			"TLS": Protocol_TLS,
			"UDP": Protocol_UDP,
			"TCP_UDP": Protocol_TCP_UDP,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.QueryStringCondition",
		reflect.TypeOf((*QueryStringCondition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.RedirectOptions",
		reflect.TypeOf((*RedirectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.SslPolicy",
		reflect.TypeOf((*SslPolicy)(nil)).Elem(),
		map[string]interface{}{
			"RECOMMENDED_TLS": SslPolicy_RECOMMENDED_TLS,
			"RECOMMENDED": SslPolicy_RECOMMENDED,
			"TLS13_RES": SslPolicy_TLS13_RES,
			"TLS13_EXT1": SslPolicy_TLS13_EXT1,
			"TLS13_EXT2": SslPolicy_TLS13_EXT2,
			"TLS13_10": SslPolicy_TLS13_10,
			"TLS13_11": SslPolicy_TLS13_11,
			"TLS13_13": SslPolicy_TLS13_13,
			"FIPS_TLS13_13": SslPolicy_FIPS_TLS13_13,
			"FIPS_TLS13_12_RES": SslPolicy_FIPS_TLS13_12_RES,
			"FIPS_TLS13_12": SslPolicy_FIPS_TLS13_12,
			"FIPS_TLS13_12_EXT0": SslPolicy_FIPS_TLS13_12_EXT0,
			"FIPS_TLS13_12_EXT1": SslPolicy_FIPS_TLS13_12_EXT1,
			"FIPS_TLS13_12_EXT2": SslPolicy_FIPS_TLS13_12_EXT2,
			"FIPS_TLS13_11": SslPolicy_FIPS_TLS13_11,
			"FIPS_TLS13_10": SslPolicy_FIPS_TLS13_10,
			"FORWARD_SECRECY_TLS12_RES_GCM": SslPolicy_FORWARD_SECRECY_TLS12_RES_GCM,
			"FORWARD_SECRECY_TLS12_RES": SslPolicy_FORWARD_SECRECY_TLS12_RES,
			"FORWARD_SECRECY_TLS12": SslPolicy_FORWARD_SECRECY_TLS12,
			"FORWARD_SECRECY_TLS11": SslPolicy_FORWARD_SECRECY_TLS11,
			"FORWARD_SECRECY": SslPolicy_FORWARD_SECRECY,
			"TLS12": SslPolicy_TLS12,
			"TLS12_EXT": SslPolicy_TLS12_EXT,
			"TLS11": SslPolicy_TLS11,
			"LEGACY": SslPolicy_LEGACY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupAttributes",
		reflect.TypeOf((*TargetGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupBase",
		reflect.TypeOf((*TargetGroupBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addLoadBalancerTarget", GoMethod: "AddLoadBalancerTarget"},
			_jsii_.MemberMethod{JsiiMethod: "configureHealthCheck", GoMethod: "ConfigureHealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPort", GoGetter: "DefaultPort"},
			_jsii_.MemberProperty{JsiiProperty: "firstLoadBalancerFullName", GoGetter: "FirstLoadBalancerFullName"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheck", GoGetter: "HealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArns", GoGetter: "LoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttached", GoGetter: "LoadBalancerAttached"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerAttachedDependencies", GoGetter: "LoadBalancerAttachedDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "setAttribute", GoMethod: "SetAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupArn", GoGetter: "TargetGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupFullName", GoGetter: "TargetGroupFullName"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupLoadBalancerArns", GoGetter: "TargetGroupLoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupName", GoGetter: "TargetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "targetType", GoGetter: "TargetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateHealthCheck", GoMethod: "ValidateHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "validateTargetGroup", GoMethod: "ValidateTargetGroup"},
		},
		func() interface{} {
			j := jsiiProxy_TargetGroupBase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITargetGroup)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupLoadBalancingAlgorithmType",
		reflect.TypeOf((*TargetGroupLoadBalancingAlgorithmType)(nil)).Elem(),
		map[string]interface{}{
			"ROUND_ROBIN": TargetGroupLoadBalancingAlgorithmType_ROUND_ROBIN,
			"LEAST_OUTSTANDING_REQUESTS": TargetGroupLoadBalancingAlgorithmType_LEAST_OUTSTANDING_REQUESTS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetType",
		reflect.TypeOf((*TargetType)(nil)).Elem(),
		map[string]interface{}{
			"INSTANCE": TargetType_INSTANCE,
			"IP": TargetType_IP,
			"LAMBDA": TargetType_LAMBDA,
			"ALB": TargetType_ALB,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_elasticloadbalancingv2.UnauthenticatedAction",
		reflect.TypeOf((*UnauthenticatedAction)(nil)).Elem(),
		map[string]interface{}{
			"DENY": UnauthenticatedAction_DENY,
			"ALLOW": UnauthenticatedAction_ALLOW,
			"AUTHENTICATE": UnauthenticatedAction_AUTHENTICATE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_elasticloadbalancingv2.WeightedTargetGroup",
		reflect.TypeOf((*WeightedTargetGroup)(nil)).Elem(),
	)
}
