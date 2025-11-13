// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultIamRole", GoMethod: "AddDefaultIamRole"},
			_jsii_.MemberMethod{JsiiMethod: "addIamRole", GoMethod: "AddIamRole"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationMultiUser", GoMethod: "AddRotationMultiUser"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSingleUser", GoMethod: "AddRotationSingleUser"},
			_jsii_.MemberMethod{JsiiMethod: "addToParameterGroup", GoMethod: "AddToParameterGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "enableRebootForParameterChanges", GoMethod: "EnableRebootForParameterChanges"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameterGroup", GoGetter: "ParameterGroup"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.ClusterAttributes",
		reflect.TypeOf((*ClusterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.ClusterParameterGroup",
		reflect.TypeOf((*ClusterParameterGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addParameter", GoMethod: "AddParameter"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterParameterGroupName", GoGetter: "ClusterParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterParameterGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IClusterParameterGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.ClusterParameterGroupProps",
		reflect.TypeOf((*ClusterParameterGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.ClusterSubnetGroup",
		reflect.TypeOf((*ClusterSubnetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSubnetGroupName", GoGetter: "ClusterSubnetGroupName"},
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
			j := jsiiProxy_ClusterSubnetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IClusterSubnetGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.ClusterSubnetGroupProps",
		reflect.TypeOf((*ClusterSubnetGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.ClusterType",
		reflect.TypeOf((*ClusterType)(nil)).Elem(),
		map[string]interface{}{
			"SINGLE_NODE": ClusterType_SINGLE_NODE,
			"MULTI_NODE": ClusterType_MULTI_NODE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.Column",
		reflect.TypeOf((*Column)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.ColumnEncoding",
		reflect.TypeOf((*ColumnEncoding)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": ColumnEncoding_AUTO,
			"RAW": ColumnEncoding_RAW,
			"AZ64": ColumnEncoding_AZ64,
			"BYTEDICT": ColumnEncoding_BYTEDICT,
			"DELTA": ColumnEncoding_DELTA,
			"DELTA32K": ColumnEncoding_DELTA32K,
			"LZO": ColumnEncoding_LZO,
			"MOSTLY8": ColumnEncoding_MOSTLY8,
			"MOSTLY16": ColumnEncoding_MOSTLY16,
			"MOSTLY32": ColumnEncoding_MOSTLY32,
			"RUNLENGTH": ColumnEncoding_RUNLENGTH,
			"TEXT255": ColumnEncoding_TEXT255,
			"TEXT32K": ColumnEncoding_TEXT32K,
			"ZSTD": ColumnEncoding_ZSTD,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.DatabaseOptions",
		reflect.TypeOf((*DatabaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.DatabaseSecret",
		reflect.TypeOf((*DatabaseSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addReplicaRegion", GoMethod: "AddReplicaRegion"},
			_jsii_.MemberMethod{JsiiMethod: "addRotationSchedule", GoMethod: "AddRotationSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "arnForPolicies", GoGetter: "ArnForPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "attach", GoMethod: "Attach"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "cfnDynamicReferenceKey", GoMethod: "CfnDynamicReferenceKey"},
			_jsii_.MemberMethod{JsiiMethod: "denyAccountRootDelete", GoMethod: "DenyAccountRootDelete"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharacters", GoGetter: "ExcludeCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretFullArn", GoGetter: "SecretFullArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretValue", GoGetter: "SecretValue"},
			_jsii_.MemberMethod{JsiiMethod: "secretValueFromJson", GoMethod: "SecretValueFromJson"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseSecret{}
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerSecret)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.DatabaseSecretProps",
		reflect.TypeOf((*DatabaseSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "socketAddress", GoGetter: "SocketAddress"},
		},
		func() interface{} {
			return &jsiiProxy_Endpoint{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-redshift-alpha.ICluster",
		reflect.TypeOf((*ICluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "asSecretAttachmentTarget", GoMethod: "AsSecretAttachmentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ICluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__awssecretsmanagerISecretAttachmentTarget)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-redshift-alpha.IClusterParameterGroup",
		reflect.TypeOf((*IClusterParameterGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterParameterGroupName", GoGetter: "ClusterParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IClusterParameterGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-redshift-alpha.IClusterSubnetGroup",
		reflect.TypeOf((*IClusterSubnetGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSubnetGroupName", GoGetter: "ClusterSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IClusterSubnetGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-redshift-alpha.ITable",
		reflect.TypeOf((*ITable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tableColumns", GoGetter: "TableColumns"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
		},
		func() interface{} {
			j := jsiiProxy_ITable{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-redshift-alpha.IUser",
		reflect.TypeOf((*IUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTablePrivileges", GoMethod: "AddTablePrivileges"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			j := jsiiProxy_IUser{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.LoggingProperties",
		reflect.TypeOf((*LoggingProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.Login",
		reflect.TypeOf((*Login)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.MaintenanceTrackName",
		reflect.TypeOf((*MaintenanceTrackName)(nil)).Elem(),
		map[string]interface{}{
			"CURRENT": MaintenanceTrackName_CURRENT,
			"TRAILING": MaintenanceTrackName_TRAILING,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.NodeType",
		reflect.TypeOf((*NodeType)(nil)).Elem(),
		map[string]interface{}{
			"DS2_XLARGE": NodeType_DS2_XLARGE,
			"DS2_8XLARGE": NodeType_DS2_8XLARGE,
			"DC1_LARGE": NodeType_DC1_LARGE,
			"DC1_8XLARGE": NodeType_DC1_8XLARGE,
			"DC2_LARGE": NodeType_DC2_LARGE,
			"DC2_8XLARGE": NodeType_DC2_8XLARGE,
			"RA3_LARGE": NodeType_RA3_LARGE,
			"RA3_XLPLUS": NodeType_RA3_XLPLUS,
			"RA3_4XLARGE": NodeType_RA3_4XLARGE,
			"RA3_16XLARGE": NodeType_RA3_16XLARGE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.ResourceAction",
		reflect.TypeOf((*ResourceAction)(nil)).Elem(),
		map[string]interface{}{
			"PAUSE_CLUSTER": ResourceAction_PAUSE_CLUSTER,
			"RESUME_CLUSTER": ResourceAction_RESUME_CLUSTER,
			"FAILOVER_PRIMARY_COMPUTE": ResourceAction_FAILOVER_PRIMARY_COMPUTE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.RotationMultiUserOptions",
		reflect.TypeOf((*RotationMultiUserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.Table",
		reflect.TypeOf((*Table)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tableColumns", GoGetter: "TableColumns"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Table{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITable)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.TableAction",
		reflect.TypeOf((*TableAction)(nil)).Elem(),
		map[string]interface{}{
			"SELECT": TableAction_SELECT,
			"INSERT": TableAction_INSERT,
			"UPDATE": TableAction_UPDATE,
			"DELETE": TableAction_DELETE,
			"DROP": TableAction_DROP,
			"REFERENCES": TableAction_REFERENCES,
			"ALL": TableAction_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.TableAttributes",
		reflect.TypeOf((*TableAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.TableDistStyle",
		reflect.TypeOf((*TableDistStyle)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": TableDistStyle_AUTO,
			"EVEN": TableDistStyle_EVEN,
			"KEY": TableDistStyle_KEY,
			"ALL": TableDistStyle_ALL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.TableProps",
		reflect.TypeOf((*TableProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-redshift-alpha.TableSortStyle",
		reflect.TypeOf((*TableSortStyle)(nil)).Elem(),
		map[string]interface{}{
			"AUTO": TableSortStyle_AUTO,
			"COMPOUND": TableSortStyle_COMPOUND,
			"INTERLEAVED": TableSortStyle_INTERLEAVED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-redshift-alpha.User",
		reflect.TypeOf((*User)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTablePrivileges", GoMethod: "AddTablePrivileges"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "databaseProps", GoGetter: "DatabaseProps"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			j := jsiiProxy_User{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUser)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.UserAttributes",
		reflect.TypeOf((*UserAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-redshift-alpha.UserProps",
		reflect.TypeOf((*UserProps)(nil)).Elem(),
	)
}
