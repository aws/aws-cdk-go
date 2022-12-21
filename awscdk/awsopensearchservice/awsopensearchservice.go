package awsopensearchservice

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.AdvancedSecurityOptions",
		reflect.TypeOf((*AdvancedSecurityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CapacityConfig",
		reflect.TypeOf((*CapacityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain",
		reflect.TypeOf((*CfnDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessPolicies", GoGetter: "AccessPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "advancedOptions", GoGetter: "AdvancedOptions"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityOptions", GoGetter: "AdvancedSecurityOptions"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainEndpoint", GoGetter: "AttrDomainEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrDomainEndpoints", GoGetter: "AttrDomainEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsAutomatedUpdateDate", GoGetter: "AttrServiceSoftwareOptionsAutomatedUpdateDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsCancellable", GoGetter: "AttrServiceSoftwareOptionsCancellable"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsCurrentVersion", GoGetter: "AttrServiceSoftwareOptionsCurrentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsDescription", GoGetter: "AttrServiceSoftwareOptionsDescription"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsNewVersion", GoGetter: "AttrServiceSoftwareOptionsNewVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsOptionalDeployment", GoGetter: "AttrServiceSoftwareOptionsOptionalDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsUpdateAvailable", GoGetter: "AttrServiceSoftwareOptionsUpdateAvailable"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceSoftwareOptionsUpdateStatus", GoGetter: "AttrServiceSoftwareOptionsUpdateStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterConfig", GoGetter: "ClusterConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoOptions", GoGetter: "CognitoOptions"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainEndpointOptions", GoGetter: "DomainEndpointOptions"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptions", GoGetter: "EbsOptions"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtRestOptions", GoGetter: "EncryptionAtRestOptions"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "logPublishingOptions", GoGetter: "LogPublishingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeToNodeEncryptionOptions", GoGetter: "NodeToNodeEncryptionOptions"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotOptions", GoGetter: "SnapshotOptions"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcOptions", GoGetter: "VpcOptions"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.AdvancedSecurityOptionsInputProperty",
		reflect.TypeOf((*CfnDomain_AdvancedSecurityOptionsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.ClusterConfigProperty",
		reflect.TypeOf((*CfnDomain_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.CognitoOptionsProperty",
		reflect.TypeOf((*CfnDomain_CognitoOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.DomainEndpointOptionsProperty",
		reflect.TypeOf((*CfnDomain_DomainEndpointOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.EBSOptionsProperty",
		reflect.TypeOf((*CfnDomain_EBSOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.EncryptionAtRestOptionsProperty",
		reflect.TypeOf((*CfnDomain_EncryptionAtRestOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.LogPublishingOptionProperty",
		reflect.TypeOf((*CfnDomain_LogPublishingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.MasterUserOptionsProperty",
		reflect.TypeOf((*CfnDomain_MasterUserOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.NodeToNodeEncryptionOptionsProperty",
		reflect.TypeOf((*CfnDomain_NodeToNodeEncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.ServiceSoftwareOptionsProperty",
		reflect.TypeOf((*CfnDomain_ServiceSoftwareOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.SnapshotOptionsProperty",
		reflect.TypeOf((*CfnDomain_SnapshotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.VPCOptionsProperty",
		reflect.TypeOf((*CfnDomain_VPCOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomain.ZoneAwarenessConfigProperty",
		reflect.TypeOf((*CfnDomain_ZoneAwarenessConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CfnDomainProps",
		reflect.TypeOf((*CfnDomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CognitoOptions",
		reflect.TypeOf((*CognitoOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.CustomEndpointOptions",
		reflect.TypeOf((*CustomEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_opensearchservice.Domain",
		reflect.TypeOf((*Domain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAccessPolicies", GoMethod: "AddAccessPolicies"},
			_jsii_.MemberProperty{JsiiProperty: "appLogGroup", GoGetter: "AppLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "auditLogGroup", GoGetter: "AuditLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "domainArn", GoGetter: "DomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainEndpoint", GoGetter: "DomainEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantIndexRead", GoMethod: "GrantIndexRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantIndexReadWrite", GoMethod: "GrantIndexReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantIndexWrite", GoMethod: "GrantIndexWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantPathRead", GoMethod: "GrantPathRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantPathReadWrite", GoMethod: "GrantPathReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantPathWrite", GoMethod: "GrantPathWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "masterUserPassword", GoGetter: "MasterUserPassword"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAutomatedSnapshotFailure", GoMethod: "MetricAutomatedSnapshotFailure"},
			_jsii_.MemberMethod{JsiiMethod: "metricClusterIndexWritesBlocked", GoMethod: "MetricClusterIndexWritesBlocked"},
			_jsii_.MemberMethod{JsiiMethod: "metricClusterStatusRed", GoMethod: "MetricClusterStatusRed"},
			_jsii_.MemberMethod{JsiiMethod: "metricClusterStatusYellow", GoMethod: "MetricClusterStatusYellow"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricIndexingLatency", GoMethod: "MetricIndexingLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricJVMMemoryPressure", GoMethod: "MetricJVMMemoryPressure"},
			_jsii_.MemberMethod{JsiiMethod: "metricKMSKeyError", GoMethod: "MetricKMSKeyError"},
			_jsii_.MemberMethod{JsiiMethod: "metricKMSKeyInaccessible", GoMethod: "MetricKMSKeyInaccessible"},
			_jsii_.MemberMethod{JsiiMethod: "metricMasterCPUUtilization", GoMethod: "MetricMasterCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricMasterJVMMemoryPressure", GoMethod: "MetricMasterJVMMemoryPressure"},
			_jsii_.MemberMethod{JsiiMethod: "metricNodes", GoMethod: "MetricNodes"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchableDocuments", GoMethod: "MetricSearchableDocuments"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchLatency", GoMethod: "MetricSearchLatency"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "slowIndexLogGroup", GoGetter: "SlowIndexLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "slowSearchLogGroup", GoGetter: "SlowSearchLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Domain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDomain)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.DomainAttributes",
		reflect.TypeOf((*DomainAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.DomainProps",
		reflect.TypeOf((*DomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.EbsOptions",
		reflect.TypeOf((*EbsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.EncryptionAtRestOptions",
		reflect.TypeOf((*EncryptionAtRestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_opensearchservice.EngineVersion",
		reflect.TypeOf((*EngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_EngineVersion{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_opensearchservice.IDomain",
		reflect.TypeOf((*IDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "domainArn", GoGetter: "DomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainEndpoint", GoGetter: "DomainEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantIndexRead", GoMethod: "GrantIndexRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantIndexReadWrite", GoMethod: "GrantIndexReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantIndexWrite", GoMethod: "GrantIndexWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantPathRead", GoMethod: "GrantPathRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantPathReadWrite", GoMethod: "GrantPathReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantPathWrite", GoMethod: "GrantPathWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricAutomatedSnapshotFailure", GoMethod: "MetricAutomatedSnapshotFailure"},
			_jsii_.MemberMethod{JsiiMethod: "metricClusterIndexWritesBlocked", GoMethod: "MetricClusterIndexWritesBlocked"},
			_jsii_.MemberMethod{JsiiMethod: "metricClusterStatusRed", GoMethod: "MetricClusterStatusRed"},
			_jsii_.MemberMethod{JsiiMethod: "metricClusterStatusYellow", GoMethod: "MetricClusterStatusYellow"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricFreeStorageSpace", GoMethod: "MetricFreeStorageSpace"},
			_jsii_.MemberMethod{JsiiMethod: "metricIndexingLatency", GoMethod: "MetricIndexingLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricJVMMemoryPressure", GoMethod: "MetricJVMMemoryPressure"},
			_jsii_.MemberMethod{JsiiMethod: "metricKMSKeyError", GoMethod: "MetricKMSKeyError"},
			_jsii_.MemberMethod{JsiiMethod: "metricKMSKeyInaccessible", GoMethod: "MetricKMSKeyInaccessible"},
			_jsii_.MemberMethod{JsiiMethod: "metricMasterCPUUtilization", GoMethod: "MetricMasterCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricMasterJVMMemoryPressure", GoMethod: "MetricMasterJVMMemoryPressure"},
			_jsii_.MemberMethod{JsiiMethod: "metricNodes", GoMethod: "MetricNodes"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchableDocuments", GoMethod: "MetricSearchableDocuments"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchLatency", GoMethod: "MetricSearchLatency"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDomain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.LoggingOptions",
		reflect.TypeOf((*LoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_opensearchservice.TLSSecurityPolicy",
		reflect.TypeOf((*TLSSecurityPolicy)(nil)).Elem(),
		map[string]interface{}{
			"TLS_1_0": TLSSecurityPolicy_TLS_1_0,
			"TLS_1_2": TLSSecurityPolicy_TLS_1_2,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_opensearchservice.ZoneAwarenessConfig",
		reflect.TypeOf((*ZoneAwarenessConfig)(nil)).Elem(),
	)
}
