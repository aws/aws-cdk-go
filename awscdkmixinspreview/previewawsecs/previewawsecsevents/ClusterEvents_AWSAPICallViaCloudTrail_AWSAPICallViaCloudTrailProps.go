package previewawsecsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Cluster aws.ecs@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var environment interface{}
//   var inferenceAcceleratorOverrides interface{}
//   var networkBindings interface{}
//   var networkInterfaces interface{}
//   var resourceRequirements interface{}
//   var tags interface{}
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
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
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		Cluster: []*string{
//   			jsii.String("cluster"),
//   		},
//   		ContainerInstance: []*string{
//   			jsii.String("containerInstance"),
//   		},
//   		Containers: []RequestParametersItem1{
//   			&RequestParametersItem1{
//   				ContainerName: []*string{
//   					jsii.String("containerName"),
//   				},
//   				ExitCode: []*string{
//   					jsii.String("exitCode"),
//   				},
//   				NetworkBindings: []interface{}{
//   					networkBindings,
//   				},
//   				Status: []*string{
//   					jsii.String("status"),
//   				},
//   			},
//   		},
//   		Count: []*string{
//   			jsii.String("count"),
//   		},
//   		EnableEcsManagedTags: []*string{
//   			jsii.String("enableEcsManagedTags"),
//   		},
//   		ExecutionStoppedAt: []*string{
//   			jsii.String("executionStoppedAt"),
//   		},
//   		LaunchType: []*string{
//   			jsii.String("launchType"),
//   		},
//   		NetworkConfiguration: &NetworkConfiguration{
//   			AwsvpcConfiguration: &AwsvpcConfiguration{
//   				AssignPublicIp: []*string{
//   					jsii.String("assignPublicIp"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		Overrides: &Overrides{
//   			ContainerOverrides: []OverridesItem{
//   				&OverridesItem{
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					Cpu: []*string{
//   						jsii.String("cpu"),
//   					},
//   					Environment: []interface{}{
//   						environment,
//   					},
//   					Memory: []*string{
//   						jsii.String("memory"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					ResourceRequirements: []interface{}{
//   						resourceRequirements,
//   					},
//   				},
//   			},
//   		},
//   		PlacementConstraints: []RequestParametersItem{
//   			&RequestParametersItem{
//   				Expression: []*string{
//   					jsii.String("expression"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   			},
//   		},
//   		PullStartedAt: []*string{
//   			jsii.String("pullStartedAt"),
//   		},
//   		PullStoppedAt: []*string{
//   			jsii.String("pullStoppedAt"),
//   		},
//   		Reason: []*string{
//   			jsii.String("reason"),
//   		},
//   		StartedBy: []*string{
//   			jsii.String("startedBy"),
//   		},
//   		Status: []*string{
//   			jsii.String("status"),
//   		},
//   		Task: []*string{
//   			jsii.String("task"),
//   		},
//   		TaskDefinition: []*string{
//   			jsii.String("taskDefinition"),
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		Acknowledgment: []*string{
//   			jsii.String("acknowledgment"),
//   		},
//   		Endpoint: []*string{
//   			jsii.String("endpoint"),
//   		},
//   		Failures: []*string{
//   			jsii.String("failures"),
//   		},
//   		Tasks: []ResponseElementsItem{
//   			&ResponseElementsItem{
//   				Attachments: []ResponseElementsItemItem{
//   					&ResponseElementsItemItem{
//   						Details: []ResponseElementsItemItemItem{
//   							&ResponseElementsItemItemItem{
//   								Name: []*string{
//   									jsii.String("name"),
//   								},
//   								Value: []*string{
//   									jsii.String("value"),
//   								},
//   							},
//   						},
//   						Id: []*string{
//   							jsii.String("id"),
//   						},
//   						Status: []*string{
//   							jsii.String("status"),
//   						},
//   						Type: []*string{
//   							jsii.String("type"),
//   						},
//   					},
//   				},
//   				ClusterArn: []*string{
//   					jsii.String("clusterArn"),
//   				},
//   				ContainerInstanceArn: []*string{
//   					jsii.String("containerInstanceArn"),
//   				},
//   				Containers: []ResponseElementsItemItem1{
//   					&ResponseElementsItemItem1{
//   						ContainerArn: []*string{
//   							jsii.String("containerArn"),
//   						},
//   						Cpu: []*string{
//   							jsii.String("cpu"),
//   						},
//   						Image: []*string{
//   							jsii.String("image"),
//   						},
//   						LastStatus: []*string{
//   							jsii.String("lastStatus"),
//   						},
//   						Memory: []*string{
//   							jsii.String("memory"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						NetworkInterfaces: []interface{}{
//   							networkInterfaces,
//   						},
//   						TaskArn: []*string{
//   							jsii.String("taskArn"),
//   						},
//   					},
//   				},
//   				Cpu: []*string{
//   					jsii.String("cpu"),
//   				},
//   				CreatedAt: []*string{
//   					jsii.String("createdAt"),
//   				},
//   				DesiredStatus: []*string{
//   					jsii.String("desiredStatus"),
//   				},
//   				Group: []*string{
//   					jsii.String("group"),
//   				},
//   				LastStatus: []*string{
//   					jsii.String("lastStatus"),
//   				},
//   				LaunchType: []*string{
//   					jsii.String("launchType"),
//   				},
//   				Memory: []*string{
//   					jsii.String("memory"),
//   				},
//   				Overrides: &Overrides1{
//   					ContainerOverrides: []Overrides1Item{
//   						&Overrides1Item{
//   							Command: []*string{
//   								jsii.String("command"),
//   							},
//   							Cpu: []*string{
//   								jsii.String("cpu"),
//   							},
//   							Environment: []interface{}{
//   								environment,
//   							},
//   							Memory: []*string{
//   								jsii.String("memory"),
//   							},
//   							Name: []*string{
//   								jsii.String("name"),
//   							},
//   							ResourceRequirements: []interface{}{
//   								resourceRequirements,
//   							},
//   						},
//   					},
//   					InferenceAcceleratorOverrides: []interface{}{
//   						inferenceAcceleratorOverrides,
//   					},
//   				},
//   				PlatformVersion: []*string{
//   					jsii.String("platformVersion"),
//   				},
//   				StartedBy: []*string{
//   					jsii.String("startedBy"),
//   				},
//   				Tags: []interface{}{
//   					tags,
//   				},
//   				TaskArn: []*string{
//   					jsii.String("taskArn"),
//   				},
//   				TaskDefinitionArn: []*string{
//   					jsii.String("taskDefinitionArn"),
//   				},
//   				Version: []*string{
//   					jsii.String("version"),
//   				},
//   			},
//   		},
//   		TelemetryEndpoint: []*string{
//   			jsii.String("telemetryEndpoint"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		InvokedBy: []*string{
//   			jsii.String("invokedBy"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   			WebIdFederationData: []*string{
//   				jsii.String("webIdFederationData"),
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *ClusterEvents_AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *ClusterEvents_AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *ClusterEvents_AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}

