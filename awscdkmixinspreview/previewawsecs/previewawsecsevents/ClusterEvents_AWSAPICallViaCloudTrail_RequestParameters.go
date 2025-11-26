package previewawsecsevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var environment interface{}
//   var networkBindings interface{}
//   var resourceRequirements interface{}
//
//   requestParameters := &RequestParameters{
//   	Cluster: []*string{
//   		jsii.String("cluster"),
//   	},
//   	ContainerInstance: []*string{
//   		jsii.String("containerInstance"),
//   	},
//   	Containers: []RequestParametersItem1{
//   		&RequestParametersItem1{
//   			ContainerName: []*string{
//   				jsii.String("containerName"),
//   			},
//   			ExitCode: []*string{
//   				jsii.String("exitCode"),
//   			},
//   			NetworkBindings: []interface{}{
//   				networkBindings,
//   			},
//   			Status: []*string{
//   				jsii.String("status"),
//   			},
//   		},
//   	},
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   	EnableEcsManagedTags: []*string{
//   		jsii.String("enableEcsManagedTags"),
//   	},
//   	ExecutionStoppedAt: []*string{
//   		jsii.String("executionStoppedAt"),
//   	},
//   	LaunchType: []*string{
//   		jsii.String("launchType"),
//   	},
//   	NetworkConfiguration: &NetworkConfiguration{
//   		AwsvpcConfiguration: &AwsvpcConfiguration{
//   			AssignPublicIp: []*string{
//   				jsii.String("assignPublicIp"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
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
//   	},
//   	PlacementConstraints: []RequestParametersItem{
//   		&RequestParametersItem{
//   			Expression: []*string{
//   				jsii.String("expression"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   	PullStartedAt: []*string{
//   		jsii.String("pullStartedAt"),
//   	},
//   	PullStoppedAt: []*string{
//   		jsii.String("pullStoppedAt"),
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	StartedBy: []*string{
//   		jsii.String("startedBy"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	Task: []*string{
//   		jsii.String("task"),
//   	},
//   	TaskDefinition: []*string{
//   		jsii.String("taskDefinition"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// cluster property.
	//
	// Specify an array of string values to match this event if the actual value of cluster is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cluster *[]*string `field:"optional" json:"cluster" yaml:"cluster"`
	// containerInstance property.
	//
	// Specify an array of string values to match this event if the actual value of containerInstance is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerInstance *[]*string `field:"optional" json:"containerInstance" yaml:"containerInstance"`
	// containers property.
	//
	// Specify an array of string values to match this event if the actual value of containers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Containers *[]*ClusterEvents_AWSAPICallViaCloudTrail_RequestParametersItem1 `field:"optional" json:"containers" yaml:"containers"`
	// count property.
	//
	// Specify an array of string values to match this event if the actual value of count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
	// enableECSManagedTags property.
	//
	// Specify an array of string values to match this event if the actual value of enableECSManagedTags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnableEcsManagedTags *[]*string `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// executionStoppedAt property.
	//
	// Specify an array of string values to match this event if the actual value of executionStoppedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionStoppedAt *[]*string `field:"optional" json:"executionStoppedAt" yaml:"executionStoppedAt"`
	// launchType property.
	//
	// Specify an array of string values to match this event if the actual value of launchType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchType *[]*string `field:"optional" json:"launchType" yaml:"launchType"`
	// networkConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of networkConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NetworkConfiguration *ClusterEvents_AWSAPICallViaCloudTrail_NetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// overrides property.
	//
	// Specify an array of string values to match this event if the actual value of overrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Overrides *ClusterEvents_AWSAPICallViaCloudTrail_Overrides `field:"optional" json:"overrides" yaml:"overrides"`
	// placementConstraints property.
	//
	// Specify an array of string values to match this event if the actual value of placementConstraints is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlacementConstraints *[]*ClusterEvents_AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
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
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// startedBy property.
	//
	// Specify an array of string values to match this event if the actual value of startedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedBy *[]*string `field:"optional" json:"startedBy" yaml:"startedBy"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// task property.
	//
	// Specify an array of string values to match this event if the actual value of task is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Task *[]*string `field:"optional" json:"task" yaml:"task"`
	// taskDefinition property.
	//
	// Specify an array of string values to match this event if the actual value of taskDefinition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskDefinition *[]*string `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
}

