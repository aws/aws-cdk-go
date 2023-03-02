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
//   cfnVolumeProps := &cfnVolumeProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	backupId: jsii.String("backupId"),
//   	ontapConfiguration: &ontapConfigurationProperty{
//   		sizeInMegabytes: jsii.String("sizeInMegabytes"),
//   		storageVirtualMachineId: jsii.String("storageVirtualMachineId"),
//
//   		// the properties below are optional
//   		copyTagsToBackups: jsii.String("copyTagsToBackups"),
//   		junctionPath: jsii.String("junctionPath"),
//   		ontapVolumeType: jsii.String("ontapVolumeType"),
//   		securityStyle: jsii.String("securityStyle"),
//   		snapshotPolicy: jsii.String("snapshotPolicy"),
//   		storageEfficiencyEnabled: jsii.String("storageEfficiencyEnabled"),
//   		tieringPolicy: &tieringPolicyProperty{
//   			coolingPeriod: jsii.Number(123),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	openZfsConfiguration: &openZFSConfigurationProperty{
//   		parentVolumeId: jsii.String("parentVolumeId"),
//
//   		// the properties below are optional
//   		copyTagsToSnapshots: jsii.Boolean(false),
//   		dataCompressionType: jsii.String("dataCompressionType"),
//   		nfsExports: []interface{}{
//   			&nfsExportsProperty{
//   				clientConfigurations: []interface{}{
//   					&clientConfigurationsProperty{
//   						clients: jsii.String("clients"),
//   						options: []*string{
//   							jsii.String("options"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		options: []*string{
//   			jsii.String("options"),
//   		},
//   		originSnapshot: &originSnapshotProperty{
//   			copyStrategy: jsii.String("copyStrategy"),
//   			snapshotArn: jsii.String("snapshotArn"),
//   		},
//   		readOnly: jsii.Boolean(false),
//   		recordSizeKiB: jsii.Number(123),
//   		storageCapacityQuotaGiB: jsii.Number(123),
//   		storageCapacityReservationGiB: jsii.Number(123),
//   		userAndGroupQuotas: []interface{}{
//   			&userAndGroupQuotasProperty{
//   				id: jsii.Number(123),
//   				storageCapacityQuotaGiB: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	volumeType: jsii.String("volumeType"),
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

