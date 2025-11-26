package previewawsecsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Cluster aws.ecs@ECSContainerInstanceStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSContainerInstanceStateChangeProps := &ECSContainerInstanceStateChangeProps{
//   	AccountType: []*string{
//   		jsii.String("accountType"),
//   	},
//   	AgentConnected: []*string{
//   		jsii.String("agentConnected"),
//   	},
//   	AgentUpdateStatus: []*string{
//   		jsii.String("agentUpdateStatus"),
//   	},
//   	Attachments: []AttachmentDetails{
//   		&AttachmentDetails{
//   			Details: []DetailsItems{
//   				&DetailsItems{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
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
//   	ClusterArn: []*string{
//   		jsii.String("clusterArn"),
//   	},
//   	ContainerInstanceArn: []*string{
//   		jsii.String("containerInstanceArn"),
//   	},
//   	Ec2InstanceId: []*string{
//   		jsii.String("ec2InstanceId"),
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
//   	PendingTasksCount: []*string{
//   		jsii.String("pendingTasksCount"),
//   	},
//   	RegisteredAt: []*string{
//   		jsii.String("registeredAt"),
//   	},
//   	RegisteredResources: []ResourceDetails{
//   		&ResourceDetails{
//   			DoubleValue: []*string{
//   				jsii.String("doubleValue"),
//   			},
//   			IntegerValue: []*string{
//   				jsii.String("integerValue"),
//   			},
//   			LongValue: []*string{
//   				jsii.String("longValue"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			StringSetValue: []*string{
//   				jsii.String("stringSetValue"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	RemainingResources: []ResourceDetails{
//   		&ResourceDetails{
//   			DoubleValue: []*string{
//   				jsii.String("doubleValue"),
//   			},
//   			IntegerValue: []*string{
//   				jsii.String("integerValue"),
//   			},
//   			LongValue: []*string{
//   				jsii.String("longValue"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			StringSetValue: []*string{
//   				jsii.String("stringSetValue"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	RunningTasksCount: []*string{
//   		jsii.String("runningTasksCount"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StatusReason: []*string{
//   		jsii.String("statusReason"),
//   	},
//   	UpdatedAt: []*string{
//   		jsii.String("updatedAt"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   	VersionInfo: &VersionInfo{
//   		AgentHash: []*string{
//   			jsii.String("agentHash"),
//   		},
//   		AgentVersion: []*string{
//   			jsii.String("agentVersion"),
//   		},
//   		DockerVersion: []*string{
//   			jsii.String("dockerVersion"),
//   		},
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps struct {
	// accountType property.
	//
	// Specify an array of string values to match this event if the actual value of accountType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountType *[]*string `field:"optional" json:"accountType" yaml:"accountType"`
	// agentConnected property.
	//
	// Specify an array of string values to match this event if the actual value of agentConnected is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentConnected *[]*string `field:"optional" json:"agentConnected" yaml:"agentConnected"`
	// agentUpdateStatus property.
	//
	// Specify an array of string values to match this event if the actual value of agentUpdateStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentUpdateStatus *[]*string `field:"optional" json:"agentUpdateStatus" yaml:"agentUpdateStatus"`
	// attachments property.
	//
	// Specify an array of string values to match this event if the actual value of attachments is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attachments *[]*ClusterEvents_ECSContainerInstanceStateChange_AttachmentDetails `field:"optional" json:"attachments" yaml:"attachments"`
	// attributes property.
	//
	// Specify an array of string values to match this event if the actual value of attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attributes *[]*ClusterEvents_ECSContainerInstanceStateChange_AttributesDetails `field:"optional" json:"attributes" yaml:"attributes"`
	// clusterArn property.
	//
	// Specify an array of string values to match this event if the actual value of clusterArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Cluster reference.
	//
	// Experimental.
	ClusterArn *[]*string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// containerInstanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of containerInstanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerInstanceArn *[]*string `field:"optional" json:"containerInstanceArn" yaml:"containerInstanceArn"`
	// ec2InstanceId property.
	//
	// Specify an array of string values to match this event if the actual value of ec2InstanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ec2InstanceId *[]*string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// pendingTasksCount property.
	//
	// Specify an array of string values to match this event if the actual value of pendingTasksCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PendingTasksCount *[]*string `field:"optional" json:"pendingTasksCount" yaml:"pendingTasksCount"`
	// registeredAt property.
	//
	// Specify an array of string values to match this event if the actual value of registeredAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegisteredAt *[]*string `field:"optional" json:"registeredAt" yaml:"registeredAt"`
	// registeredResources property.
	//
	// Specify an array of string values to match this event if the actual value of registeredResources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegisteredResources *[]*ClusterEvents_ECSContainerInstanceStateChange_ResourceDetails `field:"optional" json:"registeredResources" yaml:"registeredResources"`
	// remainingResources property.
	//
	// Specify an array of string values to match this event if the actual value of remainingResources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RemainingResources *[]*ClusterEvents_ECSContainerInstanceStateChange_ResourceDetails `field:"optional" json:"remainingResources" yaml:"remainingResources"`
	// runningTasksCount property.
	//
	// Specify an array of string values to match this event if the actual value of runningTasksCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RunningTasksCount *[]*string `field:"optional" json:"runningTasksCount" yaml:"runningTasksCount"`
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
	// versionInfo property.
	//
	// Specify an array of string values to match this event if the actual value of versionInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionInfo *ClusterEvents_ECSContainerInstanceStateChange_VersionInfo `field:"optional" json:"versionInfo" yaml:"versionInfo"`
}

