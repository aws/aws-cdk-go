package awsecs


// Represents a specific configuration revision of an Express service, containing all the settings and parameters for that revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressGatewayServiceConfigurationProperty := &ExpressGatewayServiceConfigurationProperty{
//   	Cpu: jsii.String("cpu"),
//   	CreatedAt: jsii.String("createdAt"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	IngressPaths: []interface{}{
//   		&IngressPathSummaryProperty{
//   			AccessType: jsii.String("accessType"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   	},
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
//   		Image: jsii.String("image"),
//
//   		// the properties below are optional
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
//   	ServiceRevisionArn: jsii.String("serviceRevisionArn"),
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html
//
type CfnExpressGatewayService_ExpressGatewayServiceConfigurationProperty struct {
	// The CPU allocation for tasks in this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-cpu
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The Unix timestamp for when this service revision was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The ARN of the task execution role for the service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The health check path for this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-healthcheckpath
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The entry point into this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-ingresspaths
	//
	IngressPaths interface{} `field:"optional" json:"ingressPaths" yaml:"ingressPaths"`
	// The memory allocation for tasks in this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-memory
	//
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The network configuration for tasks in this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The primary container configuration for this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-primarycontainer
	//
	PrimaryContainer interface{} `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// The auto-scaling configuration for this service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-scalingtarget
	//
	ScalingTarget interface{} `field:"optional" json:"scalingTarget" yaml:"scalingTarget"`
	// The ARN of the service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-servicerevisionarn
	//
	ServiceRevisionArn *string `field:"optional" json:"serviceRevisionArn" yaml:"serviceRevisionArn"`
	// The ARN of the task role for the service revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayserviceconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayserviceconfiguration-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}

