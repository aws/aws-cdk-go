package previewawsecsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Cluster aws.ecs@ECSTaskStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSTaskStateChangeProps := &ECSTaskStateChangeProps{
//   	Attachments: []AttachmentDetails{
//   		&AttachmentDetails{
//   			Details: &Details{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   			Id: []*string{
//   				jsii.String("id"),
//   			},
//   			Status: []*string{
//   				jsii.String("status"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	Attributes: []AttributesDetails{
//   		&AttributesDetails{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	ClusterArn: []*string{
//   		jsii.String("clusterArn"),
//   	},
//   	Connectivity: []*string{
//   		jsii.String("connectivity"),
//   	},
//   	ConnectivityAt: []*string{
//   		jsii.String("connectivityAt"),
//   	},
//   	ContainerInstanceArn: []*string{
//   		jsii.String("containerInstanceArn"),
//   	},
//   	Containers: []ContainerDetails{
//   		&ContainerDetails{
//   			ContainerArn: []*string{
//   				jsii.String("containerArn"),
//   			},
//   			Cpu: []*string{
//   				jsii.String("cpu"),
//   			},
//   			ExitCode: []*string{
//   				jsii.String("exitCode"),
//   			},
//   			GpuIds: []*string{
//   				jsii.String("gpuIds"),
//   			},
//   			Image: []*string{
//   				jsii.String("image"),
//   			},
//   			ImageDigest: []*string{
//   				jsii.String("imageDigest"),
//   			},
//   			LastStatus: []*string{
//   				jsii.String("lastStatus"),
//   			},
//   			Memory: []*string{
//   				jsii.String("memory"),
//   			},
//   			MemoryReservation: []*string{
//   				jsii.String("memoryReservation"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			NetworkBindings: []NetworkBindingDetails{
//   				&NetworkBindingDetails{
//   					BindIp: []*string{
//   						jsii.String("bindIp"),
//   					},
//   					ContainerPort: []*string{
//   						jsii.String("containerPort"),
//   					},
//   					HostPort: []*string{
//   						jsii.String("hostPort"),
//   					},
//   					Protocol: []*string{
//   						jsii.String("protocol"),
//   					},
//   				},
//   			},
//   			NetworkInterfaces: []NetworkInterfaceDetails{
//   				&NetworkInterfaceDetails{
//   					AttachmentId: []*string{
//   						jsii.String("attachmentId"),
//   					},
//   					Ipv6Address: []*string{
//   						jsii.String("ipv6Address"),
//   					},
//   					PrivateIpv4Address: []*string{
//   						jsii.String("privateIpv4Address"),
//   					},
//   				},
//   			},
//   			Reason: []*string{
//   				jsii.String("reason"),
//   			},
//   			RuntimeId: []*string{
//   				jsii.String("runtimeId"),
//   			},
//   			TaskArn: []*string{
//   				jsii.String("taskArn"),
//   			},
//   		},
//   	},
//   	Cpu: []*string{
//   		jsii.String("cpu"),
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	DesiredStatus: []*string{
//   		jsii.String("desiredStatus"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	ExecutionStoppedAt: []*string{
//   		jsii.String("executionStoppedAt"),
//   	},
//   	Group: []*string{
//   		jsii.String("group"),
//   	},
//   	LastStatus: []*string{
//   		jsii.String("lastStatus"),
//   	},
//   	LaunchType: []*string{
//   		jsii.String("launchType"),
//   	},
//   	Memory: []*string{
//   		jsii.String("memory"),
//   	},
//   	Overrides: &Overrides{
//   		ContainerOverrides: []OverridesItem{
//   			&OverridesItem{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Cpu: []*string{
//   					jsii.String("cpu"),
//   				},
//   				Environment: []map[string]*string{
//   					map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   				},
//   				Memory: []*string{
//   					jsii.String("memory"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	PlatformVersion: []*string{
//   		jsii.String("platformVersion"),
//   	},
//   	PullStartedAt: []*string{
//   		jsii.String("pullStartedAt"),
//   	},
//   	PullStoppedAt: []*string{
//   		jsii.String("pullStoppedAt"),
//   	},
//   	StartedAt: []*string{
//   		jsii.String("startedAt"),
//   	},
//   	StartedBy: []*string{
//   		jsii.String("startedBy"),
//   	},
//   	StopCode: []*string{
//   		jsii.String("stopCode"),
//   	},
//   	StoppedAt: []*string{
//   		jsii.String("stoppedAt"),
//   	},
//   	StoppedReason: []*string{
//   		jsii.String("stoppedReason"),
//   	},
//   	StoppingAt: []*string{
//   		jsii.String("stoppingAt"),
//   	},
//   	TaskArn: []*string{
//   		jsii.String("taskArn"),
//   	},
//   	TaskDefinitionArn: []*string{
//   		jsii.String("taskDefinitionArn"),
//   	},
//   	UpdatedAt: []*string{
//   		jsii.String("updatedAt"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSTaskStateChange_ECSTaskStateChangeProps struct {
	// attachments property.
	//
	// Specify an array of string values to match this event if the actual value of attachments is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attachments *[]*ClusterEvents_ECSTaskStateChange_AttachmentDetails `field:"optional" json:"attachments" yaml:"attachments"`
	// attributes property.
	//
	// Specify an array of string values to match this event if the actual value of attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attributes *[]*ClusterEvents_ECSTaskStateChange_AttributesDetails `field:"optional" json:"attributes" yaml:"attributes"`
	// availabilityZone property.
	//
	// Specify an array of string values to match this event if the actual value of availabilityZone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// clusterArn property.
	//
	// Specify an array of string values to match this event if the actual value of clusterArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Cluster reference.
	//
	// Experimental.
	ClusterArn *[]*string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// connectivity property.
	//
	// Specify an array of string values to match this event if the actual value of connectivity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Connectivity *[]*string `field:"optional" json:"connectivity" yaml:"connectivity"`
	// connectivityAt property.
	//
	// Specify an array of string values to match this event if the actual value of connectivityAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConnectivityAt *[]*string `field:"optional" json:"connectivityAt" yaml:"connectivityAt"`
	// containerInstanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of containerInstanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerInstanceArn *[]*string `field:"optional" json:"containerInstanceArn" yaml:"containerInstanceArn"`
	// containers property.
	//
	// Specify an array of string values to match this event if the actual value of containers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Containers *[]*ClusterEvents_ECSTaskStateChange_ContainerDetails `field:"optional" json:"containers" yaml:"containers"`
	// cpu property.
	//
	// Specify an array of string values to match this event if the actual value of cpu is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cpu *[]*string `field:"optional" json:"cpu" yaml:"cpu"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// desiredStatus property.
	//
	// Specify an array of string values to match this event if the actual value of desiredStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DesiredStatus *[]*string `field:"optional" json:"desiredStatus" yaml:"desiredStatus"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// executionStoppedAt property.
	//
	// Specify an array of string values to match this event if the actual value of executionStoppedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionStoppedAt *[]*string `field:"optional" json:"executionStoppedAt" yaml:"executionStoppedAt"`
	// group property.
	//
	// Specify an array of string values to match this event if the actual value of group is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// lastStatus property.
	//
	// Specify an array of string values to match this event if the actual value of lastStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastStatus *[]*string `field:"optional" json:"lastStatus" yaml:"lastStatus"`
	// launchType property.
	//
	// Specify an array of string values to match this event if the actual value of launchType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchType *[]*string `field:"optional" json:"launchType" yaml:"launchType"`
	// memory property.
	//
	// Specify an array of string values to match this event if the actual value of memory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Memory *[]*string `field:"optional" json:"memory" yaml:"memory"`
	// overrides property.
	//
	// Specify an array of string values to match this event if the actual value of overrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Overrides *ClusterEvents_ECSTaskStateChange_Overrides `field:"optional" json:"overrides" yaml:"overrides"`
	// platformVersion property.
	//
	// Specify an array of string values to match this event if the actual value of platformVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlatformVersion *[]*string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// pullStartedAt property.
	//
	// Specify an array of string values to match this event if the actual value of pullStartedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullStartedAt *[]*string `field:"optional" json:"pullStartedAt" yaml:"pullStartedAt"`
	// pullStoppedAt property.
	//
	// Specify an array of string values to match this event if the actual value of pullStoppedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PullStoppedAt *[]*string `field:"optional" json:"pullStoppedAt" yaml:"pullStoppedAt"`
	// startedAt property.
	//
	// Specify an array of string values to match this event if the actual value of startedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedAt *[]*string `field:"optional" json:"startedAt" yaml:"startedAt"`
	// startedBy property.
	//
	// Specify an array of string values to match this event if the actual value of startedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedBy *[]*string `field:"optional" json:"startedBy" yaml:"startedBy"`
	// stopCode property.
	//
	// Specify an array of string values to match this event if the actual value of stopCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StopCode *[]*string `field:"optional" json:"stopCode" yaml:"stopCode"`
	// stoppedAt property.
	//
	// Specify an array of string values to match this event if the actual value of stoppedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppedAt *[]*string `field:"optional" json:"stoppedAt" yaml:"stoppedAt"`
	// stoppedReason property.
	//
	// Specify an array of string values to match this event if the actual value of stoppedReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppedReason *[]*string `field:"optional" json:"stoppedReason" yaml:"stoppedReason"`
	// stoppingAt property.
	//
	// Specify an array of string values to match this event if the actual value of stoppingAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppingAt *[]*string `field:"optional" json:"stoppingAt" yaml:"stoppingAt"`
	// taskArn property.
	//
	// Specify an array of string values to match this event if the actual value of taskArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskArn *[]*string `field:"optional" json:"taskArn" yaml:"taskArn"`
	// taskDefinitionArn property.
	//
	// Specify an array of string values to match this event if the actual value of taskDefinitionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskDefinitionArn *[]*string `field:"optional" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// updatedAt property.
	//
	// Specify an array of string values to match this event if the actual value of updatedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UpdatedAt *[]*string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

