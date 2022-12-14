package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   ecsParametersProperty := &ecsParametersProperty{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	capacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	enableEcsManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	group: jsii.String("group"),
//   	launchType: jsii.String("launchType"),
//   	networkConfiguration: &networkConfigurationProperty{
//   		awsvpcConfiguration: &awsVpcConfigurationProperty{
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//
//   			// the properties below are optional
//   			assignPublicIp: jsii.String("assignPublicIp"),
//   			securityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   		},
//   	},
//   	placementConstraints: []interface{}{
//   		&placementConstraintProperty{
//   			expression: jsii.String("expression"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	placementStrategy: []interface{}{
//   		&placementStrategyProperty{
//   			field: jsii.String("field"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	platformVersion: jsii.String("platformVersion"),
//   	propagateTags: jsii.String("propagateTags"),
//   	referenceId: jsii.String("referenceId"),
//   	tags: tags,
//   	taskCount: jsii.Number(123),
//   }
//
type CfnSchedule_EcsParametersProperty struct {
	// `CfnSchedule.EcsParametersProperty.TaskDefinitionArn`.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// `CfnSchedule.EcsParametersProperty.CapacityProviderStrategy`.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// `CfnSchedule.EcsParametersProperty.EnableECSManagedTags`.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// `CfnSchedule.EcsParametersProperty.EnableExecuteCommand`.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// `CfnSchedule.EcsParametersProperty.Group`.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// `CfnSchedule.EcsParametersProperty.LaunchType`.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// `CfnSchedule.EcsParametersProperty.NetworkConfiguration`.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// `CfnSchedule.EcsParametersProperty.PlacementConstraints`.
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// `CfnSchedule.EcsParametersProperty.PlacementStrategy`.
	PlacementStrategy interface{} `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
	// `CfnSchedule.EcsParametersProperty.PlatformVersion`.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// `CfnSchedule.EcsParametersProperty.PropagateTags`.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// `CfnSchedule.EcsParametersProperty.ReferenceId`.
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// `CfnSchedule.EcsParametersProperty.Tags`.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// `CfnSchedule.EcsParametersProperty.TaskCount`.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

