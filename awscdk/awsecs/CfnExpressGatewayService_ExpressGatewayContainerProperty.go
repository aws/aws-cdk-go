package awsecs


// Defines the configuration for the primary container in an Express service.
//
// This container receives traffic from the Application Load Balancer and runs your application code.
//
// The container configuration includes the container image, port mapping, logging settings, environment variables, and secrets. The container image is the only required parameter, with sensible defaults provided for other settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressGatewayContainerProperty := &ExpressGatewayContainerProperty{
//   	Image: jsii.String("image"),
//
//   	// the properties below are optional
//   	AwsLogsConfiguration: &ExpressGatewayServiceAwsLogsConfigurationProperty{
//   		LogGroup: jsii.String("logGroup"),
//   		LogStreamPrefix: jsii.String("logStreamPrefix"),
//   	},
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	ContainerPort: jsii.Number(123),
//   	Environment: []interface{}{
//   		&KeyValuePairProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RepositoryCredentials: &ExpressGatewayRepositoryCredentialsProperty{
//   		CredentialsParameter: jsii.String("credentialsParameter"),
//   	},
//   	Secrets: []interface{}{
//   		&SecretProperty{
//   			Name: jsii.String("name"),
//   			ValueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html
//
type CfnExpressGatewayService_ExpressGatewayContainerProperty struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon. Images in the Docker Hub registry are available by default. Other repositories are specified with either `repository-url/image:tag` or `repository-url/image@digest` .
	//
	// For Express services, the image typically contains a web application that listens on the specified container port. The image can be stored in Amazon ECR, Docker Hub, or any other container registry accessible to your execution role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-image
	//
	Image *string `field:"required" json:"image" yaml:"image"`
	// The log configuration for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-awslogsconfiguration
	//
	AwsLogsConfiguration interface{} `field:"optional" json:"awsLogsConfiguration" yaml:"awsLogsConfiguration"`
	// The command that is passed to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The port number on the container that receives traffic from the load balancer.
	//
	// Default is 80.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-containerport
	//
	// Default: - 80.
	//
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The environment variables to pass to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The configuration for repository credentials for private registry authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-repositorycredentials
	//
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// The secrets to pass to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-secrets
	//
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
}

