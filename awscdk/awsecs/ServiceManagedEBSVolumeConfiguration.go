package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents the configuration for an ECS Service managed EBS volume.
//
// Example:
//   var cluster cluster
//
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))
//
//   container := taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	PortMappings: []portMapping{
//   		&portMapping{
//   			ContainerPort: jsii.Number(80),
//   			Protocol: ecs.Protocol_TCP,
//   		},
//   	},
//   })
//
//   volume := ecs.NewServiceManagedVolume(this, jsii.String("EBSVolume"), &ServiceManagedVolumeProps{
//   	Name: jsii.String("ebs1"),
//   	ManagedEBSVolume: &ServiceManagedEBSVolumeConfiguration{
//   		Size: awscdk.Size_Gibibytes(jsii.Number(15)),
//   		VolumeType: ec2.EbsDeviceVolumeType_GP3,
//   		FileSystemType: ecs.FileSystemType_XFS,
//   		TagSpecifications: []eBSTagSpecification{
//   			&eBSTagSpecification{
//   				Tags: map[string]*string{
//   					"purpose": jsii.String("production"),
//   				},
//   				PropagateTags: ecs.EbsPropagatedTagSource_SERVICE,
//   			},
//   		},
//   	},
//   })
//
//   volume.MountIn(container, &ContainerMountPoint{
//   	ContainerPath: jsii.String("/var/lib"),
//   	ReadOnly: jsii.Boolean(false),
//   })
//
//   taskDefinition.AddVolume(volume)
//
//   service := ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   })
//
//   service.AddVolume(volume)
//
type ServiceManagedEBSVolumeConfiguration struct {
	// Indicates whether the volume should be encrypted.
	// Default: - Default Amazon EBS encryption.
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The Linux filesystem type for the volume.
	//
	// For volumes created from a snapshot, you must specify the same filesystem type that
	// the volume was using when the snapshot was created.
	// The available filesystem types are ext3, ext4, and xfs.
	// Default: - FileSystemType.XFS
	//
	FileSystemType FileSystemType `field:"optional" json:"fileSystemType" yaml:"fileSystemType"`
	// The number of I/O operations per second (IOPS).
	//
	// For gp3, io1, and io2 volumes, this represents the number of IOPS that are provisioned
	// for the volume. For gp2 volumes, this represents the baseline performance of the volume
	// and the rate at which the volume accumulates I/O credits for bursting.
	//
	// The following are the supported values for each volume type.
	//   - gp3: 3,000 - 16,000 IOPS
	//   - io1: 100 - 64,000 IOPS
	//   - io2: 100 - 256,000 IOPS
	//
	// This parameter is required for io1 and io2 volume types. The default for gp3 volumes is
	// 3,000 IOPS. This parameter is not supported for st1, sc1, or standard volume types.
	// Default: - undefined.
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// AWS Key Management Service key to use for Amazon EBS encryption.
	// Default: - When `encryption` is turned on and no `kmsKey` is specified,
	// the default AWS managed key for Amazon EBS volumes is used.
	//
	KmsKeyId awskms.IKey `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An IAM role that allows ECS to make calls to EBS APIs on your behalf.
	//
	// This role is required to create and manage the Amazon EBS volume.
	// Default: - automatically generated role.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The size of the volume in GiB.
	//
	// You must specify either `size` or `snapshotId`.
	// You can optionally specify a volume size greater than or equal to the snapshot size.
	//
	// The following are the supported volume size values for each volume type.
	//   - gp2 and gp3: 1-16,384
	//   - io1 and io2: 4-16,384
	//   - st1 and sc1: 125-16,384
	// - standard: 1-1,024.
	// Default: - The snapshot size is used for the volume size if you specify `snapshotId`,
	// otherwise this parameter is required.
	//
	Size awscdk.Size `field:"optional" json:"size" yaml:"size"`
	// The snapshot that Amazon ECS uses to create the volume.
	//
	// You must specify either `size` or `snapshotId`.
	// Default: - No snapshot.
	//
	SnapShotId *string `field:"optional" json:"snapShotId" yaml:"snapShotId"`
	// Specifies the tags to apply to the volume and whether to propagate those tags to the volume.
	// Default: - No tags are specified.
	//
	TagSpecifications *[]*EBSTagSpecification `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The throughput to provision for a volume, in MiB/s, with a maximum of 1,000 MiB/s.
	//
	// This parameter is only supported for the gp3 volume type.
	// Default: - No throughput.
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume type.
	// Default: - ec2.EbsDeviceVolumeType.GP2
	//
	VolumeType awsec2.EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
}

