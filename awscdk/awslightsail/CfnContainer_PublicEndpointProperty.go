package awslightsail


// `PublicEndpoint` is a property of the [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html) property. It describes describes the settings of the public endpoint of a container on a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicEndpointProperty := &PublicEndpointProperty{
//   	ContainerName: jsii.String("containerName"),
//   	ContainerPort: jsii.Number(123),
//   	HealthCheckConfig: &HealthCheckConfigProperty{
//   		HealthyThreshold: jsii.Number(123),
//   		IntervalSeconds: jsii.Number(123),
//   		Path: jsii.String("path"),
//   		SuccessCodes: jsii.String("successCodes"),
//   		TimeoutSeconds: jsii.Number(123),
//   		UnhealthyThreshold: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html
//
type CfnContainer_PublicEndpointProperty struct {
	// The name of the container entry of the deployment that the endpoint configuration applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html#cfn-lightsail-container-publicendpoint-containername
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port of the specified container to which traffic is forwarded to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html#cfn-lightsail-container-publicendpoint-containerport
	//
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// An object that describes the health check configuration of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html#cfn-lightsail-container-publicendpoint-healthcheckconfig
	//
	HealthCheckConfig interface{} `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
}

