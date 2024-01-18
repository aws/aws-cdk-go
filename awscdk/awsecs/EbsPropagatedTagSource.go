package awsecs


// Propagate tags for EBS Volume Configuration from either service or task definition.
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
type EbsPropagatedTagSource string

const (
	// SERVICE.
	EbsPropagatedTagSource_SERVICE EbsPropagatedTagSource = "SERVICE"
	// TASK_DEFINITION.
	EbsPropagatedTagSource_TASK_DEFINITION EbsPropagatedTagSource = "TASK_DEFINITION"
)

