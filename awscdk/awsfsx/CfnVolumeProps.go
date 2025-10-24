package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   		StorageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   		// the properties below are optional
//   		AggregateConfiguration: &AggregateConfigurationProperty{
//   			Aggregates: []*string{
//   				jsii.String("aggregates"),
//   			},
//   			ConstituentsPerAggregate: jsii.Number(123),
//   		},
//   		CopyTagsToBackups: jsii.String("copyTagsToBackups"),
//   		JunctionPath: jsii.String("junctionPath"),
//   		OntapVolumeType: jsii.String("ontapVolumeType"),
//   		SecurityStyle: jsii.String("securityStyle"),
//   		SizeInBytes: jsii.String("sizeInBytes"),
//   		SizeInMegabytes: jsii.String("sizeInMegabytes"),
//   		SnaplockConfiguration: &SnaplockConfigurationProperty{
//   			SnaplockType: jsii.String("snaplockType"),
//
//   			// the properties below are optional
//   			AuditLogVolume: jsii.String("auditLogVolume"),
//   			AutocommitPeriod: &AutocommitPeriodProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Value: jsii.Number(123),
//   			},
//   			PrivilegedDelete: jsii.String("privilegedDelete"),
//   			RetentionPeriod: &SnaplockRetentionPeriodProperty{
//   				DefaultRetention: &RetentionPeriodProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Value: jsii.Number(123),
//   				},
//   				MaximumRetention: &RetentionPeriodProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Value: jsii.Number(123),
//   				},
//   				MinimumRetention: &RetentionPeriodProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			VolumeAppendModeEnabled: jsii.String("volumeAppendModeEnabled"),
//   		},
//   		SnapshotPolicy: jsii.String("snapshotPolicy"),
//   		StorageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   		TieringPolicy: &TieringPolicyProperty{
//   			CoolingPeriod: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   		VolumeStyle: jsii.String("volumeStyle"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html
//
type CfnVolumeProps struct {
	// The name of the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html#cfn-fsx-volume-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the ID of the volume backup to use to create a new volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html#cfn-fsx-volume-backupid
	//
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The configuration of an Amazon FSx for NetApp ONTAP volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html#cfn-fsx-volume-ontapconfiguration
	//
	OntapConfiguration interface{} `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The configuration of an Amazon FSx for OpenZFS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html#cfn-fsx-volume-openzfsconfiguration
	//
	OpenZfsConfiguration interface{} `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html#cfn-fsx-volume-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-volume.html#cfn-fsx-volume-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

