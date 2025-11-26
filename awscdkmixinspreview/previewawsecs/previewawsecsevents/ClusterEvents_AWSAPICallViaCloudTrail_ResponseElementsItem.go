package previewawsecsevents


// Type definition for ResponseElementsItem.
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
//   responseElementsItem := &ResponseElementsItem{
//   	Attachments: []ResponseElementsItemItem{
//   		&ResponseElementsItemItem{
//   			Details: []ResponseElementsItemItemItem{
//   				&ResponseElementsItemItemItem{
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
//   	ClusterArn: []*string{
//   		jsii.String("clusterArn"),
//   	},
//   	ContainerInstanceArn: []*string{
//   		jsii.String("containerInstanceArn"),
//   	},
//   	Containers: []ResponseElementsItemItem1{
//   		&ResponseElementsItemItem1{
//   			ContainerArn: []*string{
//   				jsii.String("containerArn"),
//   			},
//   			Cpu: []*string{
//   				jsii.String("cpu"),
//   			},
//   			Image: []*string{
//   				jsii.String("image"),
//   			},
//   			LastStatus: []*string{
//   				jsii.String("lastStatus"),
//   			},
//   			Memory: []*string{
//   				jsii.String("memory"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			NetworkInterfaces: []interface{}{
//   				networkInterfaces,
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
//   	Overrides: &Overrides1{
//   		ContainerOverrides: []Overrides1Item{
//   			&Overrides1Item{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Cpu: []*string{
//   					jsii.String("cpu"),
//   				},
//   				Environment: []interface{}{
//   					environment,
//   				},
//   				Memory: []*string{
//   					jsii.String("memory"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				ResourceRequirements: []interface{}{
//   					resourceRequirements,
//   				},
//   			},
//   		},
//   		InferenceAcceleratorOverrides: []interface{}{
//   			inferenceAcceleratorOverrides,
//   		},
//   	},
//   	PlatformVersion: []*string{
//   		jsii.String("platformVersion"),
//   	},
//   	StartedBy: []*string{
//   		jsii.String("startedBy"),
//   	},
//   	Tags: []interface{}{
//   		tags,
//   	},
//   	TaskArn: []*string{
//   		jsii.String("taskArn"),
//   	},
//   	TaskDefinitionArn: []*string{
//   		jsii.String("taskDefinitionArn"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItem struct {
	// attachments property.
	//
	// Specify an array of string values to match this event if the actual value of attachments is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Attachments *[]*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItemItem `field:"optional" json:"attachments" yaml:"attachments"`
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
	// containers property.
	//
	// Specify an array of string values to match this event if the actual value of containers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Containers *[]*ClusterEvents_AWSAPICallViaCloudTrail_ResponseElementsItemItem1 `field:"optional" json:"containers" yaml:"containers"`
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
	Overrides *ClusterEvents_AWSAPICallViaCloudTrail_Overrides1 `field:"optional" json:"overrides" yaml:"overrides"`
	// platformVersion property.
	//
	// Specify an array of string values to match this event if the actual value of platformVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlatformVersion *[]*string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// startedBy property.
	//
	// Specify an array of string values to match this event if the actual value of startedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedBy *[]*string `field:"optional" json:"startedBy" yaml:"startedBy"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
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
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

