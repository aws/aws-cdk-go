package awsdms


// Provides information that defines a PostgreSQL endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postgreSqlSettingsProperty := &postgreSqlSettingsProperty{
//   	afterConnectScript: jsii.String("afterConnectScript"),
//   	captureDdls: jsii.Boolean(false),
//   	ddlArtifactsSchema: jsii.String("ddlArtifactsSchema"),
//   	executeTimeout: jsii.Number(123),
//   	failTasksOnLobTruncation: jsii.Boolean(false),
//   	heartbeatEnable: jsii.Boolean(false),
//   	heartbeatFrequency: jsii.Number(123),
//   	heartbeatSchema: jsii.String("heartbeatSchema"),
//   	maxFileSize: jsii.Number(123),
//   	pluginName: jsii.String("pluginName"),
//   	secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	slotName: jsii.String("slotName"),
//   }
//
type CfnEndpoint_PostgreSqlSettingsProperty struct {
	// For use with change data capture (CDC) only, this attribute has AWS DMS bypass foreign keys and user triggers to reduce the time it takes to bulk load data.
	//
	// Example: `afterConnectScript=SET session_replication_role='replica'`.
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// To capture DDL events, AWS DMS creates various artifacts in the PostgreSQL database when the task starts.
	//
	// You can later remove these artifacts.
	//
	// If this value is set to `N` , you don't have to create tables or triggers on the source database.
	CaptureDdls interface{} `field:"optional" json:"captureDdls" yaml:"captureDdls"`
	// The schema in which the operational DDL database artifacts are created.
	//
	// Example: `ddlArtifactsSchema=xyzddlschema;`.
	DdlArtifactsSchema *string `field:"optional" json:"ddlArtifactsSchema" yaml:"ddlArtifactsSchema"`
	// Sets the client statement timeout for the PostgreSQL instance, in seconds. The default value is 60 seconds.
	//
	// Example: `executeTimeout=100;`.
	ExecuteTimeout *float64 `field:"optional" json:"executeTimeout" yaml:"executeTimeout"`
	// When set to `true` , this value causes a task to fail if the actual size of a LOB column is greater than the specified `LobMaxSize` .
	//
	// If task is set to Limited LOB mode and this option is set to true, the task fails instead of truncating the LOB data.
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// The write-ahead log (WAL) heartbeat feature mimics a dummy transaction.
	//
	// By doing this, it prevents idle logical replication slots from holding onto old WAL logs, which can result in storage full situations on the source. This heartbeat keeps `restart_lsn` moving and prevents storage full scenarios.
	HeartbeatEnable interface{} `field:"optional" json:"heartbeatEnable" yaml:"heartbeatEnable"`
	// Sets the WAL heartbeat frequency (in minutes).
	HeartbeatFrequency *float64 `field:"optional" json:"heartbeatFrequency" yaml:"heartbeatFrequency"`
	// Sets the schema in which the heartbeat artifacts are created.
	HeartbeatSchema *string `field:"optional" json:"heartbeatSchema" yaml:"heartbeatSchema"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to PostgreSQL.
	//
	// Example: `maxFileSize=512`.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Specifies the plugin to use to create a replication slot.
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the PostgreSQL endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the PostgreSQL endpoint connection details.
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Sets the name of a previously created logical replication slot for a change data capture (CDC) load of the PostgreSQL source instance.
	//
	// When used with the `CdcStartPosition` request parameter for the AWS DMS API , this attribute also makes it possible to use native CDC start points. DMS verifies that the specified logical replication slot exists before starting the CDC load task. It also verifies that the task was created with a valid setting of `CdcStartPosition` . If the specified slot doesn't exist or the task doesn't have a valid `CdcStartPosition` setting, DMS raises an error.
	//
	// For more information about setting the `CdcStartPosition` request parameter, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native) in the *AWS Database Migration Service User Guide* . For more information about using `CdcStartPosition` , see [CreateReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateReplicationTask.html) , [StartReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_StartReplicationTask.html) , and [ModifyReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_ModifyReplicationTask.html) .
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

