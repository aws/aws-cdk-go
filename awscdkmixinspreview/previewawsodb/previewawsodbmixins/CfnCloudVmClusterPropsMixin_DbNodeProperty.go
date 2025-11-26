package previewawsodbmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Information about a DB node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dbNodeProperty := &DbNodeProperty{
//   	BackupIpId: jsii.String("backupIpId"),
//   	BackupVnic2Id: jsii.String("backupVnic2Id"),
//   	CpuCoreCount: jsii.Number(123),
//   	DbNodeArn: jsii.String("dbNodeArn"),
//   	DbNodeId: jsii.String("dbNodeId"),
//   	DbNodeStorageSizeInGBs: jsii.Number(123),
//   	DbServerId: jsii.String("dbServerId"),
//   	DbSystemId: jsii.String("dbSystemId"),
//   	HostIpId: jsii.String("hostIpId"),
//   	Hostname: jsii.String("hostname"),
//   	MemorySizeInGBs: jsii.Number(123),
//   	Ocid: jsii.String("ocid"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Vnic2Id: jsii.String("vnic2Id"),
//   	VnicId: jsii.String("vnicId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html
//
type CfnCloudVmClusterPropsMixin_DbNodeProperty struct {
	// The Oracle Cloud ID (OCID) of the backup IP address that's associated with the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-backupipid
	//
	BackupIpId *string `field:"optional" json:"backupIpId" yaml:"backupIpId"`
	// The OCID of the second backup VNIC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-backupvnic2id
	//
	BackupVnic2Id *string `field:"optional" json:"backupVnic2Id" yaml:"backupVnic2Id"`
	// Number of CPU cores enabled on the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-cpucorecount
	//
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// The Amazon Resource Name (ARN) of the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-dbnodearn
	//
	DbNodeArn *string `field:"optional" json:"dbNodeArn" yaml:"dbNodeArn"`
	// The unique identifier of the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-dbnodeid
	//
	DbNodeId *string `field:"optional" json:"dbNodeId" yaml:"dbNodeId"`
	// The amount of local node storage, in gigabytes (GBs), that's allocated on the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-dbnodestoragesizeingbs
	//
	DbNodeStorageSizeInGBs *float64 `field:"optional" json:"dbNodeStorageSizeInGBs" yaml:"dbNodeStorageSizeInGBs"`
	// The unique identifier of the Db server that is associated with the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-dbserverid
	//
	DbServerId *string `field:"optional" json:"dbServerId" yaml:"dbServerId"`
	// The OCID of the DB system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-dbsystemid
	//
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// The OCID of the host IP address that's associated with the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-hostipid
	//
	HostIpId *string `field:"optional" json:"hostIpId" yaml:"hostIpId"`
	// The host name for the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-hostname
	//
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The allocated memory in GBs on the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-memorysizeingbs
	//
	MemorySizeInGBs *float64 `field:"optional" json:"memorySizeInGBs" yaml:"memorySizeInGBs"`
	// The OCID of the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-ocid
	//
	Ocid *string `field:"optional" json:"ocid" yaml:"ocid"`
	// The current status of the DB node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The OCID of the second VNIC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-vnic2id
	//
	Vnic2Id *string `field:"optional" json:"vnic2Id" yaml:"vnic2Id"`
	// The OCID of the VNIC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-cloudvmcluster-dbnode.html#cfn-odb-cloudvmcluster-dbnode-vnicid
	//
	VnicId *string `field:"optional" json:"vnicId" yaml:"vnicId"`
}

