package mixinsawsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCloudVmClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudVmClusterMixinProps := &CfnCloudVmClusterMixinProps{
//   	CloudExadataInfrastructureId: jsii.String("cloudExadataInfrastructureId"),
//   	ClusterName: jsii.String("clusterName"),
//   	CpuCoreCount: jsii.Number(123),
//   	DataCollectionOptions: &DataCollectionOptionsProperty{
//   		IsDiagnosticsEventsEnabled: jsii.Boolean(false),
//   		IsHealthMonitoringEnabled: jsii.Boolean(false),
//   		IsIncidentLogsEnabled: jsii.Boolean(false),
//   	},
//   	DataStorageSizeInTBs: jsii.Number(123),
//   	DbNodes: []interface{}{
//   		&DbNodeProperty{
//   			BackupIpId: jsii.String("backupIpId"),
//   			BackupVnic2Id: jsii.String("backupVnic2Id"),
//   			CpuCoreCount: jsii.Number(123),
//   			DbNodeArn: jsii.String("dbNodeArn"),
//   			DbNodeId: jsii.String("dbNodeId"),
//   			DbNodeStorageSizeInGBs: jsii.Number(123),
//   			DbServerId: jsii.String("dbServerId"),
//   			DbSystemId: jsii.String("dbSystemId"),
//   			HostIpId: jsii.String("hostIpId"),
//   			Hostname: jsii.String("hostname"),
//   			MemorySizeInGBs: jsii.Number(123),
//   			Ocid: jsii.String("ocid"),
//   			Status: jsii.String("status"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Vnic2Id: jsii.String("vnic2Id"),
//   			VnicId: jsii.String("vnicId"),
//   		},
//   	},
//   	DbNodeStorageSizeInGBs: jsii.Number(123),
//   	DbServers: []*string{
//   		jsii.String("dbServers"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	GiVersion: jsii.String("giVersion"),
//   	Hostname: jsii.String("hostname"),
//   	IsLocalBackupEnabled: jsii.Boolean(false),
//   	IsSparseDiskgroupEnabled: jsii.Boolean(false),
//   	LicenseModel: jsii.String("licenseModel"),
//   	MemorySizeInGBs: jsii.Number(123),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	ScanListenerPortTcp: jsii.Number(123),
//   	SshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	SystemVersion: jsii.String("systemVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html
//
type CfnCloudVmClusterMixinProps struct {
	// The unique identifier of the Exadata infrastructure that this VM cluster belongs to.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-cloudexadatainfrastructureid
	//
	CloudExadataInfrastructureId *string `field:"optional" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The name of the Grid Infrastructure (GI) cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The number of CPU cores enabled on the VM cluster.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-cpucorecount
	//
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// The set of diagnostic collection options enabled for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-datacollectionoptions
	//
	DataCollectionOptions interface{} `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-datastoragesizeintbs
	//
	DataStorageSizeInTBs *float64 `field:"optional" json:"dataStorageSizeInTBs" yaml:"dataStorageSizeInTBs"`
	// The DB nodes that are implicitly created and managed as part of this VM Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-dbnodes
	//
	DbNodes interface{} `field:"optional" json:"dbNodes" yaml:"dbNodes"`
	// The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-dbnodestoragesizeingbs
	//
	DbNodeStorageSizeInGBs *float64 `field:"optional" json:"dbNodeStorageSizeInGBs" yaml:"dbNodeStorageSizeInGBs"`
	// The list of database servers for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-dbservers
	//
	DbServers *[]*string `field:"optional" json:"dbServers" yaml:"dbServers"`
	// The user-friendly name for the VM cluster.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The software version of the Oracle Grid Infrastructure (GI) for the VM cluster.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-giversion
	//
	GiVersion *string `field:"optional" json:"giVersion" yaml:"giVersion"`
	// The host name for the VM cluster.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-hostname
	//
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Specifies whether database backups to local Exadata storage are enabled for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-islocalbackupenabled
	//
	IsLocalBackupEnabled interface{} `field:"optional" json:"isLocalBackupEnabled" yaml:"isLocalBackupEnabled"`
	// Specifies whether the VM cluster is configured with a sparse disk group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-issparsediskgroupenabled
	//
	IsSparseDiskgroupEnabled interface{} `field:"optional" json:"isSparseDiskgroupEnabled" yaml:"isSparseDiskgroupEnabled"`
	// The Oracle license model applied to the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-licensemodel
	//
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-memorysizeingbs
	//
	MemorySizeInGBs *float64 `field:"optional" json:"memorySizeInGBs" yaml:"memorySizeInGBs"`
	// The unique identifier of the ODB network for the VM cluster.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-odbnetworkid
	//
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// The port number for TCP connections to the single client access name (SCAN) listener.
	//
	// Valid values: `1024–8999` with the following exceptions: `2484` , `6100` , `6200` , `7060` , `7070` , `7085` , and `7879`
	//
	// Default: `1521`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-scanlistenerporttcp
	//
	ScanListenerPortTcp *float64 `field:"optional" json:"scanListenerPortTcp" yaml:"scanListenerPortTcp"`
	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	//
	// Required when creating a VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-sshpublickeys
	//
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// The operating system version of the image chosen for the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-systemversion
	//
	SystemVersion *string `field:"optional" json:"systemVersion" yaml:"systemVersion"`
	// Tags to assign to the Vm Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The time zone of the VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html#cfn-odb-cloudvmcluster-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

