package awslightsail


// `ContainerServiceDeployment` is a property of the [AWS::Lightsail::Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html) resource. It describes a container deployment configuration of a container service.
//
// A deployment specifies the settings, such as the ports and launch command, of containers that are deployed to your container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerServiceDeploymentProperty := &containerServiceDeploymentProperty{
//   	containers: []interface{}{
//   		&containerProperty{
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			containerName: jsii.String("containerName"),
//   			environment: []interface{}{
//   				&environmentVariableProperty{
//   					value: jsii.String("value"),
//   					variable: jsii.String("variable"),
//   				},
//   			},
//   			image: jsii.String("image"),
//   			ports: []interface{}{
//   				&portInfoProperty{
//   					port: jsii.String("port"),
//   					protocol: jsii.String("protocol"),
//   				},
//   			},
//   		},
//   	},
//   	publicEndpoint: &publicEndpointProperty{
//   		containerName: jsii.String("containerName"),
//   		containerPort: jsii.Number(123),
//   		healthCheckConfig: &healthCheckConfigProperty{
//   			healthyThreshold: jsii.Number(123),
//   			intervalSeconds: jsii.Number(123),
//   			path: jsii.String("path"),
//   			successCodes: jsii.String("successCodes"),
//   			timeoutSeconds: jsii.Number(123),
//   			unhealthyThreshold: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnContainer_ContainerServiceDeploymentProperty struct {
	// An object that describes the configuration for the containers of the deployment.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// An object that describes the endpoint of the deployment.
	PublicEndpoint interface{} `field:"optional" json:"publicEndpoint" yaml:"publicEndpoint"`
}

