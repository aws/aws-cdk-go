package awslightsail


// `HealthCheckConfig` is a property of the [PublicEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html) property. It describes the healthcheck configuration of a container deployment on a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &healthCheckConfigProperty{
//   	healthyThreshold: jsii.Number(123),
//   	intervalSeconds: jsii.Number(123),
//   	path: jsii.String("path"),
//   	successCodes: jsii.String("successCodes"),
//   	timeoutSeconds: jsii.Number(123),
//   	unhealthyThreshold: jsii.Number(123),
//   }
//
type CfnContainer_HealthCheckConfigProperty struct {
	// The number of consecutive health check successes required before moving the container to the `Healthy` state.
	//
	// The default value is `2` .
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual container.
	//
	// You can specify between `5` and `300` seconds. The default value is `5` .
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The path on the container on which to perform the health check.
	//
	// The default value is `/` .
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP codes to use when checking for a successful response from a container.
	//
	// You can specify values between `200` and `499` . You can specify multiple values (for example, `200,202` ) or a range of values (for example, `200-299` ).
	SuccessCodes *string `field:"optional" json:"successCodes" yaml:"successCodes"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// You can specify between `2` and `60` seconds. The default value is `2` .
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The number of consecutive health check failures required before moving the container to the `Unhealthy` state.
	//
	// The default value is `2` .
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

