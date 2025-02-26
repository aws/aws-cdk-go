package awsecs


// Type of resource to set a limit on.
//
// Example:
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	Ulimits: []ulimit{
//   		&ulimit{
//   			HardLimit: jsii.Number(128),
//   			Name: ecs.UlimitName_RSS,
//   			SoftLimit: jsii.Number(128),
//   		},
//   	},
//   })
//
type UlimitName string

const (
	UlimitName_CORE UlimitName = "CORE"
	UlimitName_CPU UlimitName = "CPU"
	UlimitName_DATA UlimitName = "DATA"
	UlimitName_FSIZE UlimitName = "FSIZE"
	UlimitName_LOCKS UlimitName = "LOCKS"
	UlimitName_MEMLOCK UlimitName = "MEMLOCK"
	UlimitName_MSGQUEUE UlimitName = "MSGQUEUE"
	UlimitName_NICE UlimitName = "NICE"
	UlimitName_NOFILE UlimitName = "NOFILE"
	UlimitName_NPROC UlimitName = "NPROC"
	UlimitName_RSS UlimitName = "RSS"
	UlimitName_RTPRIO UlimitName = "RTPRIO"
	UlimitName_RTTIME UlimitName = "RTTIME"
	UlimitName_SIGPENDING UlimitName = "SIGPENDING"
	UlimitName_STACK UlimitName = "STACK"
)

