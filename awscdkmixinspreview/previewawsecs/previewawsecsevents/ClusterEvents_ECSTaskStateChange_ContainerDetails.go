package previewawsecsevents


// Type definition for ContainerDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerDetails := &ContainerDetails{
//   	ContainerArn: []*string{
//   		jsii.String("containerArn"),
//   	},
//   	Cpu: []*string{
//   		jsii.String("cpu"),
//   	},
//   	ExitCode: []*string{
//   		jsii.String("exitCode"),
//   	},
//   	GpuIds: []*string{
//   		jsii.String("gpuIds"),
//   	},
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	ImageDigest: []*string{
//   		jsii.String("imageDigest"),
//   	},
//   	LastStatus: []*string{
//   		jsii.String("lastStatus"),
//   	},
//   	Memory: []*string{
//   		jsii.String("memory"),
//   	},
//   	MemoryReservation: []*string{
//   		jsii.String("memoryReservation"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	NetworkBindings: []NetworkBindingDetails{
//   		&NetworkBindingDetails{
//   			BindIp: []*string{
//   				jsii.String("bindIp"),
//   			},
//   			ContainerPort: []*string{
//   				jsii.String("containerPort"),
//   			},
//   			HostPort: []*string{
//   				jsii.String("hostPort"),
//   			},
//   			Protocol: []*string{
//   				jsii.String("protocol"),
//   			},
//   		},
//   	},
//   	NetworkInterfaces: []NetworkInterfaceDetails{
//   		&NetworkInterfaceDetails{
//   			AttachmentId: []*string{
//   				jsii.String("attachmentId"),
//   			},
//   			Ipv6Address: []*string{
//   				jsii.String("ipv6Address"),
//   			},
//   			PrivateIpv4Address: []*string{
//   				jsii.String("privateIpv4Address"),
//   			},
//   		},
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	RuntimeId: []*string{
//   		jsii.String("runtimeId"),
//   	},
//   	TaskArn: []*string{
//   		jsii.String("taskArn"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSTaskStateChange_ContainerDetails struct {
	// containerArn property.
	//
	// Specify an array of string values to match this event if the actual value of containerArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerArn *[]*string `field:"optional" json:"containerArn" yaml:"containerArn"`
	// cpu property.
	//
	// Specify an array of string values to match this event if the actual value of cpu is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cpu *[]*string `field:"optional" json:"cpu" yaml:"cpu"`
	// exitCode property.
	//
	// Specify an array of string values to match this event if the actual value of exitCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExitCode *[]*string `field:"optional" json:"exitCode" yaml:"exitCode"`
	// gpuIds property.
	//
	// Specify an array of string values to match this event if the actual value of gpuIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GpuIds *[]*string `field:"optional" json:"gpuIds" yaml:"gpuIds"`
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// imageDigest property.
	//
	// Specify an array of string values to match this event if the actual value of imageDigest is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageDigest *[]*string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// lastStatus property.
	//
	// Specify an array of string values to match this event if the actual value of lastStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastStatus *[]*string `field:"optional" json:"lastStatus" yaml:"lastStatus"`
	// memory property.
	//
	// Specify an array of string values to match this event if the actual value of memory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Memory *[]*string `field:"optional" json:"memory" yaml:"memory"`
	// memoryReservation property.
	//
	// Specify an array of string values to match this event if the actual value of memoryReservation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MemoryReservation *[]*string `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// networkBindings property.
	//
	// Specify an array of string values to match this event if the actual value of networkBindings is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkBindings *[]*ClusterEvents_ECSTaskStateChange_NetworkBindingDetails `field:"optional" json:"networkBindings" yaml:"networkBindings"`
	// networkInterfaces property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaces is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaces *[]*ClusterEvents_ECSTaskStateChange_NetworkInterfaceDetails `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// runtimeId property.
	//
	// Specify an array of string values to match this event if the actual value of runtimeId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RuntimeId *[]*string `field:"optional" json:"runtimeId" yaml:"runtimeId"`
	// taskArn property.
	//
	// Specify an array of string values to match this event if the actual value of taskArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskArn *[]*string `field:"optional" json:"taskArn" yaml:"taskArn"`
}

