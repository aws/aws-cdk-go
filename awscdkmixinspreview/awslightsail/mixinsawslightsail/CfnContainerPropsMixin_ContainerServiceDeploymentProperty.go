package mixinsawslightsail


// `ContainerServiceDeployment` is a property of the [AWS::Lightsail::Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html) resource. It describes a container deployment configuration of a container service.
//
// A deployment specifies the settings, such as the ports and launch command, of containers that are deployed to your container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerServiceDeploymentProperty := &ContainerServiceDeploymentProperty{
//   	Containers: []interface{}{
//   		&ContainerProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			ContainerName: jsii.String("containerName"),
//   			Environment: []interface{}{
//   				&EnvironmentVariableProperty{
//   					Value: jsii.String("value"),
//   					Variable: jsii.String("variable"),
//   				},
//   			},
//   			Image: jsii.String("image"),
//   			Ports: []interface{}{
//   				&PortInfoProperty{
//   					Port: jsii.String("port"),
//   					Protocol: jsii.String("protocol"),
//   				},
//   			},
//   		},
//   	},
//   	PublicEndpoint: &PublicEndpointProperty{
//   		ContainerName: jsii.String("containerName"),
//   		ContainerPort: jsii.Number(123),
//   		HealthCheckConfig: &HealthCheckConfigProperty{
//   			HealthyThreshold: jsii.Number(123),
//   			IntervalSeconds: jsii.Number(123),
//   			Path: jsii.String("path"),
//   			SuccessCodes: jsii.String("successCodes"),
//   			TimeoutSeconds: jsii.Number(123),
//   			UnhealthyThreshold: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html
//
type CfnContainerPropsMixin_ContainerServiceDeploymentProperty struct {
	// An object that describes the configuration for the containers of the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html#cfn-lightsail-container-containerservicedeployment-containers
	//
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// An object that describes the endpoint of the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html#cfn-lightsail-container-containerservicedeployment-publicendpoint
	//
	PublicEndpoint interface{} `field:"optional" json:"publicEndpoint" yaml:"publicEndpoint"`
}

