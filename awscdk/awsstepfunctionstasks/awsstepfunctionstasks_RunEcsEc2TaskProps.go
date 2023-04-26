package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties to run an ECS task on EC2 in StepFunctionsan ECS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var placementConstraint placementConstraint
//   var placementStrategy placementStrategy
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var taskDefinition taskDefinition
//
//   runEcsEc2TaskProps := &RunEcsEc2TaskProps{
//   	Cluster: cluster,
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MemoryLimit: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   		},
//   	},
//   	IntegrationPattern: awscdk.Aws_stepfunctions.ServiceIntegrationPattern_FIRE_AND_FORGET,
//   	PlacementConstraints: []*placementConstraint{
//   		placementConstraint,
//   	},
//   	PlacementStrategies: []*placementStrategy{
//   		placementStrategy,
//   	},
//   	SecurityGroup: securityGroup,
//   	Subnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		SubnetName: jsii.String("subnetName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   	},
//   }
//
// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
type RunEcsEc2TaskProps struct {
	// The topic to run the task on.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	TaskDefinition awsecs.TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	ContainerOverrides *[]*ContainerOverride `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// The service integration pattern indicates different ways to call RunTask in ECS.
	//
	// The valid value for Lambda is FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Placement constraints.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	PlacementConstraints *[]awsecs.PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// Placement strategies.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	PlacementStrategies *[]awsecs.PlacementStrategy `field:"optional" json:"placementStrategies" yaml:"placementStrategies"`
	// Existing security group to use for the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

