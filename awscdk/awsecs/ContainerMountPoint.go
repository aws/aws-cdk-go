package awsecs


// Defines the mount point details for attaching a volume to a container.
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
//   	MinHealthyPercent: jsii.Number(100),
//   })
//
//   service.AddVolume(volume)
//
type ContainerMountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Specifies whether to give the container read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
}

