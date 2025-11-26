package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   expressGatewayContainerProperty := &ExpressGatewayContainerProperty{
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
//   	Image: jsii.String("image"),
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
type CfnExpressGatewayServicePropsMixin_ExpressGatewayContainerProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-awslogsconfiguration
	//
	AwsLogsConfiguration interface{} `field:"optional" json:"awsLogsConfiguration" yaml:"awsLogsConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-containerport
	//
	// Default: - 80.
	//
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-repositorycredentials
	//
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html#cfn-ecs-expressgatewayservice-expressgatewaycontainer-secrets
	//
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
}

