package previewawsbatchevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.batch@BatchJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var networkInterfaces interface{}
//
//   batchJobStateChangeProps := &BatchJobStateChangeProps{
//   	Attempts: []BatchJobStateChangeItem{
//   		&BatchJobStateChangeItem{
//   			Container: &Container1{
//   				ContainerInstanceArn: []*string{
//   					jsii.String("containerInstanceArn"),
//   				},
//   				ExitCode: []*string{
//   					jsii.String("exitCode"),
//   				},
//   				LogStreamName: []*string{
//   					jsii.String("logStreamName"),
//   				},
//   				NetworkInterfaces: []interface{}{
//   					networkInterfaces,
//   				},
//   				TaskArn: []*string{
//   					jsii.String("taskArn"),
//   				},
//   			},
//   			StartedAt: []*string{
//   				jsii.String("startedAt"),
//   			},
//   			StatusReason: []*string{
//   				jsii.String("statusReason"),
//   			},
//   			StoppedAt: []*string{
//   				jsii.String("stoppedAt"),
//   			},
//   		},
//   	},
//   	Container: &Container{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		ContainerInstanceArn: []*string{
//   			jsii.String("containerInstanceArn"),
//   		},
//   		Environment: []ContainerItem{
//   			&ContainerItem{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		ExitCode: []*string{
//   			jsii.String("exitCode"),
//   		},
//   		Image: []*string{
//   			jsii.String("image"),
//   		},
//   		LogStreamName: []*string{
//   			jsii.String("logStreamName"),
//   		},
//   		Memory: []*string{
//   			jsii.String("memory"),
//   		},
//   		MountPoints: []MountPoint{
//   			&MountPoint{
//   				ContainerPath: []*string{
//   					jsii.String("containerPath"),
//   				},
//   				ReadOnly: []*string{
//   					jsii.String("readOnly"),
//   				},
//   				SourceVolume: []*string{
//   					jsii.String("sourceVolume"),
//   				},
//   			},
//   		},
//   		NetworkInterfaces: []NetworkInterface{
//   			&NetworkInterface{
//   				AttachmentId: []*string{
//   					jsii.String("attachmentId"),
//   				},
//   				Ipv6Address: []*string{
//   					jsii.String("ipv6Address"),
//   				},
//   				PrivateIpv4Address: []*string{
//   					jsii.String("privateIpv4Address"),
//   				},
//   			},
//   		},
//   		ResourceRequirements: []ResourceRequirement{
//   			&ResourceRequirement{
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		TaskArn: []*string{
//   			jsii.String("taskArn"),
//   		},
//   		Ulimits: []ULimit{
//   			&ULimit{
//   				HardLimit: []*string{
//   					jsii.String("hardLimit"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				SoftLimit: []*string{
//   					jsii.String("softLimit"),
//   				},
//   			},
//   		},
//   		Vcpus: []*string{
//   			jsii.String("vcpus"),
//   		},
//   		Volumes: []Volumes{
//   			&Volumes{
//   				Host: &Host{
//   					SourcePath: []*string{
//   						jsii.String("sourcePath"),
//   					},
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	DependsOn: []JobDependency{
//   		&JobDependency{
//   			JobId: []*string{
//   				jsii.String("jobId"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
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
//   	JobDefinition: []*string{
//   		jsii.String("jobDefinition"),
//   	},
//   	JobId: []*string{
//   		jsii.String("jobId"),
//   	},
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   	JobQueue: []*string{
//   		jsii.String("jobQueue"),
//   	},
//   	Parameters: []*string{
//   		jsii.String("parameters"),
//   	},
//   	RetryStrategy: &RetryStrategy{
//   		Attempts: []*string{
//   			jsii.String("attempts"),
//   		},
//   	},
//   	StartedAt: []*string{
//   		jsii.String("startedAt"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StatusReason: []*string{
//   		jsii.String("statusReason"),
//   	},
//   	StoppedAt: []*string{
//   		jsii.String("stoppedAt"),
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_BatchJobStateChangeProps struct {
	// attempts property.
	//
	// Specify an array of string values to match this event if the actual value of attempts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attempts *[]*BatchJobStateChange_BatchJobStateChangeItem `field:"optional" json:"attempts" yaml:"attempts"`
	// container property.
	//
	// Specify an array of string values to match this event if the actual value of container is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Container *BatchJobStateChange_Container `field:"optional" json:"container" yaml:"container"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// dependsOn property.
	//
	// Specify an array of string values to match this event if the actual value of dependsOn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DependsOn *[]*BatchJobStateChange_JobDependency `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// jobDefinition property.
	//
	// Specify an array of string values to match this event if the actual value of jobDefinition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobDefinition *[]*string `field:"optional" json:"jobDefinition" yaml:"jobDefinition"`
	// jobId property.
	//
	// Specify an array of string values to match this event if the actual value of jobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobId *[]*string `field:"optional" json:"jobId" yaml:"jobId"`
	// jobName property.
	//
	// Specify an array of string values to match this event if the actual value of jobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
	// jobQueue property.
	//
	// Specify an array of string values to match this event if the actual value of jobQueue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobQueue *[]*string `field:"optional" json:"jobQueue" yaml:"jobQueue"`
	// parameters property.
	//
	// Specify an array of string values to match this event if the actual value of parameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// retryStrategy property.
	//
	// Specify an array of string values to match this event if the actual value of retryStrategy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetryStrategy *BatchJobStateChange_RetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// startedAt property.
	//
	// Specify an array of string values to match this event if the actual value of startedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedAt *[]*string `field:"optional" json:"startedAt" yaml:"startedAt"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// statusReason property.
	//
	// Specify an array of string values to match this event if the actual value of statusReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusReason *[]*string `field:"optional" json:"statusReason" yaml:"statusReason"`
	// stoppedAt property.
	//
	// Specify an array of string values to match this event if the actual value of stoppedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StoppedAt *[]*string `field:"optional" json:"stoppedAt" yaml:"stoppedAt"`
}

