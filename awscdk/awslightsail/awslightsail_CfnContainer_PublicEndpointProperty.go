package awslightsail


// `PublicEndpoint` is a property of the [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html) property. It describes describes the settings of the public endpoint of a container on a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicEndpointProperty := &publicEndpointProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	healthCheckConfig: &healthCheckConfigProperty{
//   		healthyThreshold: jsii.Number(123),
//   		intervalSeconds: jsii.Number(123),
//   		path: jsii.String("path"),
//   		successCodes: jsii.String("successCodes"),
//   		timeoutSeconds: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   }
//
type CfnContainer_PublicEndpointProperty struct {
	// The name of the container entry of the deployment that the endpoint configuration applies to.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port of the specified container to which traffic is forwarded to.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// An object that describes the health check configuration of the container.
	HealthCheckConfig interface{} `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
}

