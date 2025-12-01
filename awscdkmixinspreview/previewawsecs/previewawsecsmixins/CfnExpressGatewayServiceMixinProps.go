package previewawsecsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnExpressGatewayServicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExpressGatewayServiceMixinProps := &CfnExpressGatewayServiceMixinProps{
//   	Cluster: jsii.String("cluster"),
//   	Cpu: jsii.String("cpu"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	Memory: jsii.String("memory"),
//   	NetworkConfiguration: &ExpressGatewayServiceNetworkConfigurationProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   	PrimaryContainer: &ExpressGatewayContainerProperty{
//   		AwsLogsConfiguration: &ExpressGatewayServiceAwsLogsConfigurationProperty{
//   			LogGroup: jsii.String("logGroup"),
//   			LogStreamPrefix: jsii.String("logStreamPrefix"),
//   		},
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		ContainerPort: jsii.Number(123),
//   		Environment: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Image: jsii.String("image"),
//   		RepositoryCredentials: &ExpressGatewayRepositoryCredentialsProperty{
//   			CredentialsParameter: jsii.String("credentialsParameter"),
//   		},
//   		Secrets: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	ScalingTarget: &ExpressGatewayScalingTargetProperty{
//   		AutoScalingMetric: jsii.String("autoScalingMetric"),
//   		AutoScalingTargetValue: jsii.Number(123),
//   		MaxTaskCount: jsii.Number(123),
//   		MinTaskCount: jsii.Number(123),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html
//
type CfnExpressGatewayServiceMixinProps struct {
	// The short name or full ARN of the cluster that hosts the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-cluster
	//
	// Default: - "default".
	//
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// The CPU allocation for tasks in this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-cpu
	//
	// Default: - "256".
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The ARN of the task execution role for the service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The health check path for this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-healthcheckpath
	//
	// Default: - "HTTP:80/ping".
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The ARN of the infrastructure role that manages AWS resources for the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-infrastructurerolearn
	//
	InfrastructureRoleArn *string `field:"optional" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// The memory allocation for tasks in this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-memory
	//
	// Default: - "512".
	//
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The network configuration for tasks in this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The primary container configuration for this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-primarycontainer
	//
	PrimaryContainer interface{} `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// The auto-scaling configuration for this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-scalingtarget
	//
	ScalingTarget interface{} `field:"optional" json:"scalingTarget" yaml:"scalingTarget"`
	// The name of the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The metadata applied to the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the task role for the service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html#cfn-ecs-expressgatewayservice-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}

