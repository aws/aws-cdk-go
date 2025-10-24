package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReplicationConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var replicationSettings interface{}
//   var supplementalSettings interface{}
//   var tableMappings interface{}
//
//   cfnReplicationConfigProps := &CfnReplicationConfigProps{
//   	ComputeConfig: &ComputeConfigProperty{
//   		MaxCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		DnsNameServers: jsii.String("dnsNameServers"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		MinCapacityUnits: jsii.Number(123),
//   		MultiAz: jsii.Boolean(false),
//   		PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   		ReplicationSubnetGroupId: jsii.String("replicationSubnetGroupId"),
//   		VpcSecurityGroupIds: []*string{
//   			jsii.String("vpcSecurityGroupIds"),
//   		},
//   	},
//   	ReplicationConfigIdentifier: jsii.String("replicationConfigIdentifier"),
//   	ReplicationType: jsii.String("replicationType"),
//   	SourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	TableMappings: tableMappings,
//   	TargetEndpointArn: jsii.String("targetEndpointArn"),
//
//   	// the properties below are optional
//   	ReplicationSettings: replicationSettings,
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	SupplementalSettings: supplementalSettings,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html
//
type CfnReplicationConfigProps struct {
	// Configuration parameters for provisioning an AWS DMS Serverless replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-computeconfig
	//
	ComputeConfig interface{} `field:"required" json:"computeConfig" yaml:"computeConfig"`
	// A unique identifier that you want to use to create a `ReplicationConfigArn` that is returned as part of the output from this action.
	//
	// You can then pass this output `ReplicationConfigArn` as the value of the `ReplicationConfigArn` option for other actions to identify both AWS DMS Serverless replications and replication configurations that you want those actions to operate on. For some actions, you can also use either this unique identifier or a corresponding ARN in action filters to identify the specific replication and replication configuration to operate on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationconfigidentifier
	//
	ReplicationConfigIdentifier *string `field:"required" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration.
	//
	// Possible values:
	//
	// - `"full-load"`
	// - `"cdc"`
	// - `"full-load-and-cdc"`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationtype
	//
	ReplicationType *string `field:"required" json:"replicationType" yaml:"replicationType"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-sourceendpointarn
	//
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration.
	//
	// For more information, see [Specifying table selection and transformations rules using JSON](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.SelectionTransformation.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-tablemappings
	//
	TableMappings interface{} `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS serverless replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-targetendpointarn
	//
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Optional JSON settings for AWS DMS Serverless replications that are provisioned using this replication configuration.
	//
	// For example, see [Change processing tuning settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.ChangeProcessingTuning.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationsettings
	//
	ReplicationSettings interface{} `field:"optional" json:"replicationSettings" yaml:"replicationSettings"`
	// Optional unique value or name that you set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource.
	//
	// For more information, see [Fine-grained access control using resource names and tags](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.FineGrainedAccess) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Optional JSON settings for specifying supplemental data.
	//
	// For more information, see [Specifying supplemental data for task settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-supplementalsettings
	//
	SupplementalSettings interface{} `field:"optional" json:"supplementalSettings" yaml:"supplementalSettings"`
	// One or more optional tags associated with resources used by the AWS DMS Serverless replication.
	//
	// For more information, see [Tagging resources in AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

