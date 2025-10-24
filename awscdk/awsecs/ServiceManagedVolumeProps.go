package awsecs


// Represents the Volume configuration for an ECS service.
//
// Example:
//   var cluster Cluster
//
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))
//
//   container := taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	PortMappings: []PortMapping{
//   		&PortMapping{
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
//   		TagSpecifications: []EBSTagSpecification{
//   			&EBSTagSpecification{
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
type ServiceManagedVolumeProps struct {
	// The name of the volume.
	//
	// This corresponds to the name provided in the ECS TaskDefinition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration for an Amazon Elastic Block Store (EBS) volume managed by ECS.
	// Default: - undefined.
	//
	ManagedEBSVolume *ServiceManagedEBSVolumeConfiguration `field:"optional" json:"managedEBSVolume" yaml:"managedEBSVolume"`
}

