package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration and status for a replica location in a multi-location canary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   replicaProperty := &ReplicaProperty{
//   	CanaryState: jsii.String("canaryState"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LastModified: jsii.Number(123),
//   	Location: jsii.String("location"),
//   	ReplicationStatus: &ReplicaReplicationStatusProperty{
//   		State: jsii.String("state"),
//   	},
//   	ResourcesToReplicateTags: []*string{
//   		jsii.String("resourcesToReplicateTags"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VPCConfigProperty{
//   		Ipv6AllowedForDualStack: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []interface{}{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html
//
type CfnCanaryPropsMixin_ReplicaProperty struct {
	// State of the replica canary (CREATING, READY, RUNNING, etc.).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-canarystate
	//
	CanaryState *string `field:"optional" json:"canaryState" yaml:"canaryState"`
	// ARN of the KMS key used to encrypt the replica canary's Lambda function environment variables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Last modified timestamp of the replica.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-lastmodified
	//
	LastModified *float64 `field:"optional" json:"lastModified" yaml:"lastModified"`
	// AWS region for the replica (e.g., us-east-1).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Replication status details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-replicationstatus
	//
	ReplicationStatus interface{} `field:"optional" json:"replicationStatus" yaml:"replicationStatus"`
	// Resources to replicate tags to for this replica (e.g., lambda-function).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-resourcestoreplicatetags
	//
	ResourcesToReplicateTags *[]*string `field:"optional" json:"resourcesToReplicateTags" yaml:"resourcesToReplicateTags"`
	// Tags to apply to this replica canary and optionally its Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replica.html#cfn-synthetics-canary-replica-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

