package previewawsecsevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var environment interface{}
//   var inferenceAcceleratorOverrides interface{}
//   var networkInterfaces interface{}
//   var resourceRequirements interface{}
//   var tags interface{}
//
//   responseElements := &ResponseElements{
//   	Acknowledgment: []*string{
//   		jsii.String("acknowledgment"),
//   	},
//   	Endpoint: []*string{
//   		jsii.String("endpoint"),
//   	},
//   	Failures: []*string{
//   		jsii.String("failures"),
//   	},
//   	Tasks: []ResponseElementsItem{
//   		&ResponseElementsItem{
//   			Attachments: []ResponseElementsItemItem{
//   				&ResponseElementsItemItem{
//   					Details: []ResponseElementsItemItemItem{
//   						&ResponseElementsItemItemItem{
//   							Name: []*string{
//   								jsii.String("name"),
//   							},
//   							Value: []*string{
//   								jsii.String("value"),
//   							},
//   						},
//   					},
//   					Id: []*string{
//   						jsii.String("id"),
//   					},
//   					Status: []*string{
//   						jsii.String("status"),
//   					},
//   					Type: []*string{
//   						jsii.String("type"),
//   					},
//   				},
//   			},
//   			ClusterArn: []*string{
//   				jsii.String("clusterArn"),
//   			},
//   			ContainerInstanceArn: []*string{
//   				jsii.String("containerInstanceArn"),
//   			},
//   			Containers: []ResponseElementsItemItem1{
//   				&ResponseElementsItemItem1{
//   					ContainerArn: []*string{
//   						jsii.String("containerArn"),
//   					},
//   					Cpu: []*string{
//   						jsii.String("cpu"),
//   					},
//   					Image: []*string{
//   						jsii.String("image"),
//   					},
//   					LastStatus: []*string{
//   						jsii.String("lastStatus"),
//   					},
//   					Memory: []*string{
//   						jsii.String("memory"),
//   					},
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					NetworkInterfaces: []interface{}{
//   						networkInterfaces,
//   					},
//   					TaskArn: []*string{
//   						jsii.String("taskArn"),
//   					},
//   				},
//   			},
//   			Cpu: []*string{
//   				jsii.String("cpu"),
//   			},
//   			CreatedAt: []*string{
//   				jsii.String("createdAt"),
//   			},
//   			DesiredStatus: []*string{
//   				jsii.String("desiredStatus"),
//   			},
//   			Group: []*string{
//   				jsii.String("group"),
//   			},
//   			LastStatus: []*string{
//   				jsii.String("lastStatus"),
//   			},
//   			LaunchType: []*string{
//   				jsii.String("launchType"),
//   			},
//   			Memory: []*string{
//   				jsii.String("memory"),
//   			},
//   			Overrides: &Overrides1{
//   				ContainerOverrides: []Overrides1Item{
//   					&Overrides1Item{
//   						Command: []*string{
//   							jsii.String("command"),
//   						},
//   						Cpu: []*string{
//   							jsii.String("cpu"),
//   						},
//   						Environment: []interface{}{
//   							environment,
//   						},
//   						Memory: []*string{
//   							jsii.String("memory"),
//   						},
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						ResourceRequirements: []interface{}{
//   							resourceRequirements,
//   						},
//   					},
//   				},
//   				InferenceAcceleratorOverrides: []interface{}{
//   					inferenceAcceleratorOverrides,
//   				},
//   			},
//   			PlatformVersion: []*string{
//   				jsii.String("platformVersion"),
//   			},
//   			StartedBy: []*string{
//   				jsii.String("startedBy"),
//   			},
//   			Tags: []interface{}{
//   				tags,
//   			},
//   			TaskArn: []*string{
//   				jsii.String("taskArn"),
//   			},
//   			TaskDefinitionArn: []*string{
//   				jsii.String("taskDefinitionArn"),
//   			},
//   			Version: []*string{
//   				jsii.String("version"),
//   			},
//   		},
//   	},
//   	TelemetryEndpoint: []*string{
//   		jsii.String("telemetryEndpoint"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_ResponseElements struct {
	// acknowledgment property.
	//
	// Specify an array of string values to match this event if the actual value of acknowledgment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Acknowledgment *[]*string `field:"optional" json:"acknowledgment" yaml:"acknowledgment"`
	// endpoint property.
	//
	// Specify an array of string values to match this event if the actual value of endpoint is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Endpoint *[]*string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// failures property.
	//
	// Specify an array of string values to match this event if the actual value of failures is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Failures *[]*string `field:"optional" json:"failures" yaml:"failures"`
	// tasks property.
	//
	// Specify an array of string values to match this event if the actual value of tasks is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tasks *[]*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItem `field:"optional" json:"tasks" yaml:"tasks"`
	// telemetryEndpoint property.
	//
	// Specify an array of string values to match this event if the actual value of telemetryEndpoint is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TelemetryEndpoint *[]*string `field:"optional" json:"telemetryEndpoint" yaml:"telemetryEndpoint"`
}

