package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReplicationTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationTaskProps := &cfnReplicationTaskProps{
//   	migrationType: jsii.String("migrationType"),
//   	replicationInstanceArn: jsii.String("replicationInstanceArn"),
//   	sourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	tableMappings: jsii.String("tableMappings"),
//   	targetEndpointArn: jsii.String("targetEndpointArn"),
//
//   	// the properties below are optional
//   	cdcStartPosition: jsii.String("cdcStartPosition"),
//   	cdcStartTime: jsii.Number(123),
//   	cdcStopPosition: jsii.String("cdcStopPosition"),
//   	replicationTaskIdentifier: jsii.String("replicationTaskIdentifier"),
//   	replicationTaskSettings: jsii.String("replicationTaskSettings"),
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskData: jsii.String("taskData"),
//   }
//
type CfnReplicationTaskProps struct {
	// The migration type.
	//
	// Valid values: `full-load` | `cdc` | `full-load-and-cdc`.
	MigrationType *string `field:"required" json:"migrationType" yaml:"migrationType"`
	// The Amazon Resource Name (ARN) of a replication instance.
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// The table mappings for the task, in JSON format.
	//
	// For more information, see [Using Table Mapping to Specify Task Settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html) in the *AWS Database Migration Service User Guide* .
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Indicates when you want a change data capture (CDC) operation to start.
	//
	// Use either `CdcStartPosition` or `CdcStartTime` to specify when you want a CDC operation to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, log sequence number (LSN), or system change number (SCN) format.
	//
	// Here is a date example: `--cdc-start-position "2018-03-08T12:12:12"`
	//
	// Here is a checkpoint example: `--cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"`
	//
	// Here is an LSN example: `--cdc-start-position “mysql-bin-changelog.000024:373”`
	//
	// > When you use this task setting with a source PostgreSQL database, a logical replication slot should already be created and associated with the source endpoint. You can verify this by setting the `slotName` extra connection attribute to the name of this logical replication slot. For more information, see [Extra Connection Attributes When Using PostgreSQL as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// Indicates the start time for a change data capture (CDC) operation.
	CdcStartTime *float64 `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// Indicates when you want a change data capture (CDC) operation to stop.
	//
	// The value can be either server time or commit time.
	//
	// Here is a server time example: `--cdc-stop-position "server_time:2018-02-09T12:12:12"`
	//
	// Here is a commit time example: `--cdc-stop-position "commit_time: 2018-02-09T12:12:12"`.
	CdcStopPosition *string `field:"optional" json:"cdcStopPosition" yaml:"cdcStopPosition"`
	// An identifier for the replication task.
	//
	// Constraints:
	//
	// - Must contain 1-255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationTaskIdentifier *string `field:"optional" json:"replicationTaskIdentifier" yaml:"replicationTaskIdentifier"`
	// Overall settings for the task, in JSON format.
	//
	// For more information, see [Specifying Task Settings for AWS Database Migration Service Tasks](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html) in the *AWS Database Migration Service User Guide* .
	ReplicationTaskSettings *string `field:"optional" json:"replicationTaskSettings" yaml:"replicationTaskSettings"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// One or more tags to be assigned to the replication task.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::DMS::ReplicationTask.TaskData`.
	TaskData *string `field:"optional" json:"taskData" yaml:"taskData"`
}

