package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// The properties used to define an ECS cluster.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	memoryMiB: jsii.String("512"),
//   	cpu: jsii.String("256"),
//   	compatibility: ecs.compatibility_FARGATE,
//   })
//
//   containerDefinition := taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	assignPublicIp: jsii.Boolean(true),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("SOME_KEY"),
//   					value: sfn.jsonPath.stringAt(jsii.String("$.SomeKey")),
//   				},
//   			},
//   		},
//   	},
//   	launchTarget: tasks.NewEcsFargateLaunchTarget(),
//   })
//
// Experimental.
type ClusterProps struct {
	// The ec2 capacity to add to the cluster.
	// Experimental.
	Capacity *AddCapacityOptions `field:"optional" json:"capacity" yaml:"capacity"`
	// The capacity providers to add to the cluster.
	// Deprecated: Use {@link ClusterProps.enableFargateCapacityProviders} instead.
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
	// The name for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// If true CloudWatch Container Insights will be enabled for the cluster.
	// Experimental.
	ContainerInsights *bool `field:"optional" json:"containerInsights" yaml:"containerInsights"`
	// The service discovery namespace created in this cluster.
	// Experimental.
	DefaultCloudMapNamespace *CloudMapNamespaceOptions `field:"optional" json:"defaultCloudMapNamespace" yaml:"defaultCloudMapNamespace"`
	// Whether to enable Fargate Capacity Providers.
	// Experimental.
	EnableFargateCapacityProviders *bool `field:"optional" json:"enableFargateCapacityProviders" yaml:"enableFargateCapacityProviders"`
	// The execute command configuration for the cluster.
	// Experimental.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// The VPC where your ECS instances will be running or your ENIs will be deployed.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

