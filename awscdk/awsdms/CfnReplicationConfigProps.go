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
//   	ReplicationConfigArn: jsii.String("replicationConfigArn"),
//   	ReplicationConfigIdentifier: jsii.String("replicationConfigIdentifier"),
//   	ReplicationSettings: replicationSettings,
//   	ReplicationType: jsii.String("replicationType"),
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	SourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	SupplementalSettings: supplementalSettings,
//   	TableMappings: tableMappings,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetEndpointArn: jsii.String("targetEndpointArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html
//
type CfnReplicationConfigProps struct {
	// Configuration parameters for provisioning a AWS DMS Serverless replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-computeconfig
	//
	ComputeConfig interface{} `field:"optional" json:"computeConfig" yaml:"computeConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationconfigarn
	//
	ReplicationConfigArn *string `field:"optional" json:"replicationConfigArn" yaml:"replicationConfigArn"`
	// A unique identifier of replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationconfigidentifier
	//
	ReplicationConfigIdentifier *string `field:"optional" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// JSON settings for Servereless replications that are provisioned using this replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationsettings
	//
	ReplicationSettings interface{} `field:"optional" json:"replicationSettings" yaml:"replicationSettings"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-replicationtype
	//
	ReplicationType *string `field:"optional" json:"replicationType" yaml:"replicationType"`
	// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-sourceendpointarn
	//
	SourceEndpointArn *string `field:"optional" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// JSON settings for specifying supplemental data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-supplementalsettings
	//
	SupplementalSettings interface{} `field:"optional" json:"supplementalSettings" yaml:"supplementalSettings"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-tablemappings
	//
	TableMappings interface{} `field:"optional" json:"tableMappings" yaml:"tableMappings"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationconfig.html#cfn-dms-replicationconfig-targetendpointarn
	//
	TargetEndpointArn *string `field:"optional" json:"targetEndpointArn" yaml:"targetEndpointArn"`
}

