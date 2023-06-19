package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVolume`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeProps := &CfnVolumeProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	BackupId: jsii.String("backupId"),
//   	OntapConfiguration: &OntapConfigurationProperty{
//   		SizeInMegabytes: jsii.String("sizeInMegabytes"),
//   		StorageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   		// the properties below are optional
//   		CopyTagsToBackups: jsii.String("copyTagsToBackups"),
//   		JunctionPath: jsii.String("junctionPath"),
//   		OntapVolumeType: jsii.String("ontapVolumeType"),
//   		SecurityStyle: jsii.String("securityStyle"),
//   		SnapshotPolicy: jsii.String("snapshotPolicy"),
//   		StorageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   		TieringPolicy: &TieringPolicyProperty{
//   			CoolingPeriod: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	OpenZfsConfiguration: &OpenZFSConfigurationProperty{
//   		ParentVolumeId: jsii.String("parentVolumeId"),
//
//   		// the properties below are optional
//   		CopyTagsToSnapshots: jsii.Boolean(false),
//   		DataCompressionType: jsii.String("dataCompressionType"),
//   		NfsExports: []interface{}{
//   			&NfsExportsProperty{
//   				ClientConfigurations: []interface{}{
//   					&ClientConfigurationsProperty{
//   						Clients: jsii.String("clients"),
//   						Options: []*string{
//   							jsii.String("options"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Options: []*string{
//   			jsii.String("options"),
//   		},
//   		OriginSnapshot: &OriginSnapshotProperty{
//   			CopyStrategy: jsii.String("copyStrategy"),
//   			SnapshotArn: jsii.String("snapshotArn"),
//   		},
//   		ReadOnly: jsii.Boolean(false),
//   		RecordSizeKiB: jsii.Number(123),
//   		StorageCapacityQuotaGiB: jsii.Number(123),
//   		StorageCapacityReservationGiB: jsii.Number(123),
//   		UserAndGroupQuotas: []interface{}{
//   			&UserAndGroupQuotasProperty{
//   				Id: jsii.Number(123),
//   				StorageCapacityQuotaGiB: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VolumeType: jsii.String("volumeType"),
//   }
//
type CfnVolumeProps struct {
	// The name of the volume.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the ID of the volume backup to use to create a new volume.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The configuration of an Amazon FSx for NetApp ONTAP volume.
	OntapConfiguration interface{} `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The configuration of an Amazon FSx for OpenZFS volume.
	OpenZfsConfiguration interface{} `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the volume.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

