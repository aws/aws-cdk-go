package awsdrs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReplicationConfigurationTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationConfigurationTemplateProps := &CfnReplicationConfigurationTemplateProps{
//   	BandwidthThrottling: jsii.Number(123),
//   	EbsEncryption: jsii.String("ebsEncryption"),
//   	PitPolicy: []interface{}{
//   		&PITPolicyRuleProperty{
//   			Interval: jsii.Number(123),
//   			RetentionDuration: jsii.Number(123),
//   			Units: jsii.String("units"),
//
//   			// the properties below are optional
//   			Enabled: jsii.Boolean(false),
//   			RuleId: jsii.Number(123),
//   		},
//   	},
//   	ReplicationServersSecurityGroupsIDs: []*string{
//   		jsii.String("replicationServersSecurityGroupsIDs"),
//   	},
//   	StagingAreaSubnetId: jsii.String("stagingAreaSubnetId"),
//   	StagingAreaTags: map[string]*string{
//   		"stagingAreaTagsKey": jsii.String("stagingAreaTags"),
//   	},
//
//   	// the properties below are optional
//   	AssociateDefaultSecurityGroup: jsii.Boolean(false),
//   	AutoReplicateNewDisks: jsii.Boolean(false),
//   	CreatePublicIp: jsii.Boolean(false),
//   	DataPlaneRouting: jsii.String("dataPlaneRouting"),
//   	DefaultLargeStagingDiskType: jsii.String("defaultLargeStagingDiskType"),
//   	EbsEncryptionKeyArn: jsii.String("ebsEncryptionKeyArn"),
//   	InternetProtocol: jsii.String("internetProtocol"),
//   	ReplicationServerInstanceType: jsii.String("replicationServerInstanceType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseDedicatedReplicationServer: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html
//
type CfnReplicationConfigurationTemplateProps struct {
	// Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-bandwidththrottling
	//
	BandwidthThrottling *float64 `field:"required" json:"bandwidthThrottling" yaml:"bandwidthThrottling"`
	// The type of EBS encryption to be used during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-ebsencryption
	//
	EbsEncryption *string `field:"required" json:"ebsEncryption" yaml:"ebsEncryption"`
	// The Point in time (PIT) policy to manage snapshots taken during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-pitpolicy
	//
	PitPolicy interface{} `field:"required" json:"pitPolicy" yaml:"pitPolicy"`
	// The security group IDs that will be used by the replication server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-replicationserverssecuritygroupsids
	//
	ReplicationServersSecurityGroupsIDs *[]*string `field:"required" json:"replicationServersSecurityGroupsIDs" yaml:"replicationServersSecurityGroupsIDs"`
	// The subnet to be used by the replication staging area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-stagingareasubnetid
	//
	StagingAreaSubnetId *string `field:"required" json:"stagingAreaSubnetId" yaml:"stagingAreaSubnetId"`
	// A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-stagingareatags
	//
	StagingAreaTags interface{} `field:"required" json:"stagingAreaTags" yaml:"stagingAreaTags"`
	// Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-associatedefaultsecuritygroup
	//
	AssociateDefaultSecurityGroup interface{} `field:"optional" json:"associateDefaultSecurityGroup" yaml:"associateDefaultSecurityGroup"`
	// Whether to allow the AWS replication agent to automatically replicate newly added disks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-autoreplicatenewdisks
	//
	AutoReplicateNewDisks interface{} `field:"optional" json:"autoReplicateNewDisks" yaml:"autoReplicateNewDisks"`
	// Whether to create a Public IP for the Recovery Instance by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-createpublicip
	//
	CreatePublicIp interface{} `field:"optional" json:"createPublicIp" yaml:"createPublicIp"`
	// The data plane routing mechanism that will be used for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-dataplanerouting
	//
	DataPlaneRouting *string `field:"optional" json:"dataPlaneRouting" yaml:"dataPlaneRouting"`
	// The Staging Disk EBS volume type to be used during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-defaultlargestagingdisktype
	//
	DefaultLargeStagingDiskType *string `field:"optional" json:"defaultLargeStagingDiskType" yaml:"defaultLargeStagingDiskType"`
	// The ARN of the EBS encryption key to be used during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-ebsencryptionkeyarn
	//
	EbsEncryptionKeyArn *string `field:"optional" json:"ebsEncryptionKeyArn" yaml:"ebsEncryptionKeyArn"`
	// Which version of the Internet Protocol to use for replication of data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-internetprotocol
	//
	InternetProtocol *string `field:"optional" json:"internetProtocol" yaml:"internetProtocol"`
	// The instance type to be used for the replication server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-replicationserverinstancetype
	//
	ReplicationServerInstanceType *string `field:"optional" json:"replicationServerInstanceType" yaml:"replicationServerInstanceType"`
	// A set of tags to be associated with the Replication Configuration Template resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Whether to use a dedicated Replication Server in the replication staging area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-replicationconfigurationtemplate.html#cfn-drs-replicationconfigurationtemplate-usededicatedreplicationserver
	//
	UseDedicatedReplicationServer interface{} `field:"optional" json:"useDedicatedReplicationServer" yaml:"useDedicatedReplicationServer"`
}

