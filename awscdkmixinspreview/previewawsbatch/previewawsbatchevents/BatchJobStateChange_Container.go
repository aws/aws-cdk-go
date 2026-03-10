package previewawsbatchevents


// Type definition for Container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   container := &Container{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	ContainerInstanceArn: []*string{
//   		jsii.String("containerInstanceArn"),
//   	},
//   	Environment: []ContainerItem{
//   		&ContainerItem{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	ExitCode: []*string{
//   		jsii.String("exitCode"),
//   	},
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	LogStreamName: []*string{
//   		jsii.String("logStreamName"),
//   	},
//   	Memory: []*string{
//   		jsii.String("memory"),
//   	},
//   	MountPoints: []MountPoint{
//   		&MountPoint{
//   			ContainerPath: []*string{
//   				jsii.String("containerPath"),
//   			},
//   			ReadOnly: []*string{
//   				jsii.String("readOnly"),
//   			},
//   			SourceVolume: []*string{
//   				jsii.String("sourceVolume"),
//   			},
//   		},
//   	},
//   	NetworkInterfaces: []NetworkInterface{
//   		&NetworkInterface{
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
//   	ResourceRequirements: []ResourceRequirement{
//   		&ResourceRequirement{
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	TaskArn: []*string{
//   		jsii.String("taskArn"),
//   	},
//   	Ulimits: []ULimit{
//   		&ULimit{
//   			HardLimit: []*string{
//   				jsii.String("hardLimit"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			SoftLimit: []*string{
//   				jsii.String("softLimit"),
//   			},
//   		},
//   	},
//   	Vcpus: []*string{
//   		jsii.String("vcpus"),
//   	},
//   	Volumes: []Volumes{
//   		&Volumes{
//   			Host: &Host{
//   				SourcePath: []*string{
//   					jsii.String("sourcePath"),
//   				},
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_Container struct {
	// command property.
	//
	// Specify an array of string values to match this event if the actual value of command is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// containerInstanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of containerInstanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerInstanceArn *[]*string `field:"optional" json:"containerInstanceArn" yaml:"containerInstanceArn"`
	// environment property.
	//
	// Specify an array of string values to match this event if the actual value of environment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Environment *[]*BatchJobStateChange_ContainerItem `field:"optional" json:"environment" yaml:"environment"`
	// exitCode property.
	//
	// Specify an array of string values to match this event if the actual value of exitCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExitCode *[]*string `field:"optional" json:"exitCode" yaml:"exitCode"`
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// logStreamName property.
	//
	// Specify an array of string values to match this event if the actual value of logStreamName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LogStreamName *[]*string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// memory property.
	//
	// Specify an array of string values to match this event if the actual value of memory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Memory *[]*string `field:"optional" json:"memory" yaml:"memory"`
	// mountPoints property.
	//
	// Specify an array of string values to match this event if the actual value of mountPoints is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MountPoints *[]*BatchJobStateChange_MountPoint `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// networkInterfaces property.
	//
	// Specify an array of string values to match this event if the actual value of networkInterfaces is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkInterfaces *[]*BatchJobStateChange_NetworkInterface `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// resourceRequirements property.
	//
	// Specify an array of string values to match this event if the actual value of resourceRequirements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceRequirements *[]*BatchJobStateChange_ResourceRequirement `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// taskArn property.
	//
	// Specify an array of string values to match this event if the actual value of taskArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskArn *[]*string `field:"optional" json:"taskArn" yaml:"taskArn"`
	// ulimits property.
	//
	// Specify an array of string values to match this event if the actual value of ulimits is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ulimits *[]*BatchJobStateChange_ULimit `field:"optional" json:"ulimits" yaml:"ulimits"`
	// vcpus property.
	//
	// Specify an array of string values to match this event if the actual value of vcpus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Vcpus *[]*string `field:"optional" json:"vcpus" yaml:"vcpus"`
	// volumes property.
	//
	// Specify an array of string values to match this event if the actual value of volumes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Volumes *[]*BatchJobStateChange_Volumes `field:"optional" json:"volumes" yaml:"volumes"`
}

