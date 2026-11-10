package awsmgn

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
//   	AssociateDefaultSecurityGroup: jsii.Boolean(false),
//   	BandwidthThrottling: jsii.Number(123),
//   	CreatePublicIp: jsii.Boolean(false),
//   	DataPlaneRouting: jsii.String("dataPlaneRouting"),
//   	DefaultLargeStagingDiskType: jsii.String("defaultLargeStagingDiskType"),
//   	EbsEncryption: jsii.String("ebsEncryption"),
//   	ReplicationServerInstanceType: jsii.String("replicationServerInstanceType"),
//   	ReplicationServersSecurityGroupsIDs: []*string{
//   		jsii.String("replicationServersSecurityGroupsIDs"),
//   	},
//   	StagingAreaSubnetId: jsii.String("stagingAreaSubnetId"),
//   	StagingAreaTags: map[string]*string{
//   		"stagingAreaTagsKey": jsii.String("stagingAreaTags"),
//   	},
//   	UseDedicatedReplicationServer: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	EbsEncryptionKeyArn: jsii.String("ebsEncryptionKeyArn"),
//   	InternetProtocol: jsii.String("internetProtocol"),
//   	StoreSnapshotOnLocalZone: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseFipsEndpoint: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html
//
type CfnReplicationConfigurationTemplateProps struct {
	// Whether to associate the default Application Migration Service Security Group with the Replication Configuration Template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-associatedefaultsecuritygroup
	//
	AssociateDefaultSecurityGroup interface{} `field:"required" json:"associateDefaultSecurityGroup" yaml:"associateDefaultSecurityGroup"`
	// Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-bandwidththrottling
	//
	BandwidthThrottling *float64 `field:"required" json:"bandwidthThrottling" yaml:"bandwidthThrottling"`
	// Whether to create a Public IP for the Recovery Instance by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-createpublicip
	//
	CreatePublicIp interface{} `field:"required" json:"createPublicIp" yaml:"createPublicIp"`
	// The data plane routing mechanism that will be used for replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-dataplanerouting
	//
	DataPlaneRouting *string `field:"required" json:"dataPlaneRouting" yaml:"dataPlaneRouting"`
	// The Staging Disk EBS volume type to be used during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-defaultlargestagingdisktype
	//
	DefaultLargeStagingDiskType *string `field:"required" json:"defaultLargeStagingDiskType" yaml:"defaultLargeStagingDiskType"`
	// The type of EBS encryption to be used during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-ebsencryption
	//
	EbsEncryption *string `field:"required" json:"ebsEncryption" yaml:"ebsEncryption"`
	// The instance type to be used for the replication server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-replicationserverinstancetype
	//
	ReplicationServerInstanceType *string `field:"required" json:"replicationServerInstanceType" yaml:"replicationServerInstanceType"`
	// The security group IDs that will be used by the replication server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-replicationserverssecuritygroupsids
	//
	ReplicationServersSecurityGroupsIDs *[]*string `field:"required" json:"replicationServersSecurityGroupsIDs" yaml:"replicationServersSecurityGroupsIDs"`
	// The subnet to be used by the replication staging area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-stagingareasubnetid
	//
	StagingAreaSubnetId *string `field:"required" json:"stagingAreaSubnetId" yaml:"stagingAreaSubnetId"`
	// A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-stagingareatags
	//
	StagingAreaTags interface{} `field:"required" json:"stagingAreaTags" yaml:"stagingAreaTags"`
	// Whether to use a dedicated Replication Server in the replication staging area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-usededicatedreplicationserver
	//
	UseDedicatedReplicationServer interface{} `field:"required" json:"useDedicatedReplicationServer" yaml:"useDedicatedReplicationServer"`
	// The ARN of the EBS encryption key to be used during replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-ebsencryptionkeyarn
	//
	EbsEncryptionKeyArn *string `field:"optional" json:"ebsEncryptionKeyArn" yaml:"ebsEncryptionKeyArn"`
	// Which internet protocol to use for replication of data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-internetprotocol
	//
	InternetProtocol *string `field:"optional" json:"internetProtocol" yaml:"internetProtocol"`
	// Whether to store the Snapshot in the Local Zone of the replication staging area subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-storesnapshotonlocalzone
	//
	StoreSnapshotOnLocalZone interface{} `field:"optional" json:"storeSnapshotOnLocalZone" yaml:"storeSnapshotOnLocalZone"`
	// A set of tags to be associated with the Replication Configuration Template resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Whether to use a FIPS endpoint when initiating the agent connection to the Application Migration Service service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-replicationconfigurationtemplate.html#cfn-mgn-replicationconfigurationtemplate-usefipsendpoint
	//
	UseFipsEndpoint interface{} `field:"optional" json:"useFipsEndpoint" yaml:"useFipsEndpoint"`
}

