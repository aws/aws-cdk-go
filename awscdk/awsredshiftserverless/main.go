package awsredshiftserverless

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
		reflect.TypeOf((*CfnNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsername", GoGetter: "AdminUsername"},
			_jsii_.MemberProperty{JsiiProperty: "adminUserPassword", GoGetter: "AdminUserPassword"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceAdminUsername", GoGetter: "AttrNamespaceAdminUsername"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceCreationDate", GoGetter: "AttrNamespaceCreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceDbName", GoGetter: "AttrNamespaceDbName"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceDefaultIamRoleArn", GoGetter: "AttrNamespaceDefaultIamRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceIamRoles", GoGetter: "AttrNamespaceIamRoles"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceKmsKeyId", GoGetter: "AttrNamespaceKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceLogExports", GoGetter: "AttrNamespaceLogExports"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceNamespaceArn", GoGetter: "AttrNamespaceNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceNamespaceId", GoGetter: "AttrNamespaceNamespaceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceNamespaceName", GoGetter: "AttrNamespaceNamespaceName"},
			_jsii_.MemberProperty{JsiiProperty: "attrNamespaceStatus", GoGetter: "AttrNamespaceStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbName", GoGetter: "DbName"},
			_jsii_.MemberProperty{JsiiProperty: "defaultIamRoleArn", GoGetter: "DefaultIamRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "finalSnapshotName", GoGetter: "FinalSnapshotName"},
			_jsii_.MemberProperty{JsiiProperty: "finalSnapshotRetentionPeriod", GoGetter: "FinalSnapshotRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoles", GoGetter: "IamRoles"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logExports", GoGetter: "LogExports"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceName", GoGetter: "NamespaceName"},
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
			j := jsiiProxy_CfnNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace.NamespaceProperty",
		reflect.TypeOf((*CfnNamespace_NamespaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespaceProps",
		reflect.TypeOf((*CfnNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroup",
		reflect.TypeOf((*CfnWorkgroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupBaseCapacity", GoGetter: "AttrWorkgroupBaseCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupCreationDate", GoGetter: "AttrWorkgroupCreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupEndpointAddress", GoGetter: "AttrWorkgroupEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupEndpointPort", GoGetter: "AttrWorkgroupEndpointPort"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupEnhancedVpcRouting", GoGetter: "AttrWorkgroupEnhancedVpcRouting"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupNamespaceName", GoGetter: "AttrWorkgroupNamespaceName"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupPubliclyAccessible", GoGetter: "AttrWorkgroupPubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupSecurityGroupIds", GoGetter: "AttrWorkgroupSecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupStatus", GoGetter: "AttrWorkgroupStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupSubnetIds", GoGetter: "AttrWorkgroupSubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupWorkgroupArn", GoGetter: "AttrWorkgroupWorkgroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupWorkgroupId", GoGetter: "AttrWorkgroupWorkgroupId"},
			_jsii_.MemberProperty{JsiiProperty: "attrWorkgroupWorkgroupName", GoGetter: "AttrWorkgroupWorkgroupName"},
			_jsii_.MemberProperty{JsiiProperty: "baseCapacity", GoGetter: "BaseCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configParameters", GoGetter: "ConfigParameters"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enhancedVpcRouting", GoGetter: "EnhancedVpcRouting"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceName", GoGetter: "NamespaceName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workgroup", GoGetter: "Workgroup"},
			_jsii_.MemberProperty{JsiiProperty: "workgroupName", GoGetter: "WorkgroupName"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkgroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroup.ConfigParameterProperty",
		reflect.TypeOf((*CfnWorkgroup_ConfigParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroup.EndpointProperty",
		reflect.TypeOf((*CfnWorkgroup_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroup.NetworkInterfaceProperty",
		reflect.TypeOf((*CfnWorkgroup_NetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroup.VpcEndpointProperty",
		reflect.TypeOf((*CfnWorkgroup_VpcEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroup.WorkgroupProperty",
		reflect.TypeOf((*CfnWorkgroup_WorkgroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_redshiftserverless.CfnWorkgroupProps",
		reflect.TypeOf((*CfnWorkgroupProps)(nil)).Elem(),
	)
}
