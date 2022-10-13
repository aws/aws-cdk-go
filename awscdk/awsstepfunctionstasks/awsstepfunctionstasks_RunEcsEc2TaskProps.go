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
//   runEcsEc2TaskProps := &runEcsEc2TaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			memoryLimit: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   		},
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	placementConstraints: []*placementConstraint{
//   		placementConstraint,
//   	},
//   	placementStrategies: []*placementStrategy{
//   		placementStrategy,
//   	},
//   	securityGroup: securityGroup,
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
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

