package awslightsail


// `HealthCheckConfig` is a property of the [PublicEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html) property. It describes the healthcheck configuration of a container deployment on a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &HealthCheckConfigProperty{
//   	HealthyThreshold: jsii.Number(123),
//   	IntervalSeconds: jsii.Number(123),
//   	Path: jsii.String("path"),
//   	SuccessCodes: jsii.String("successCodes"),
//   	TimeoutSeconds: jsii.Number(123),
//   	UnhealthyThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html
//
type CfnContainer_HealthCheckConfigProperty struct {
	// The number of consecutive health check successes required before moving the container to the `Healthy` state.
	//
	// The default value is `2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html#cfn-lightsail-container-healthcheckconfig-healthythreshold
	//
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual container.
	//
	// You can specify between `5` and `300` seconds. The default value is `5` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html#cfn-lightsail-container-healthcheckconfig-intervalseconds
	//
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The path on the container on which to perform the health check.
	//
	// The default value is `/` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html#cfn-lightsail-container-healthcheckconfig-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP codes to use when checking for a successful response from a container.
	//
	// You can specify values between `200` and `499` . You can specify multiple values (for example, `200,202` ) or a range of values (for example, `200-299` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html#cfn-lightsail-container-healthcheckconfig-successcodes
	//
	SuccessCodes *string `field:"optional" json:"successCodes" yaml:"successCodes"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// You can specify between `2` and `60` seconds. The default value is `2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html#cfn-lightsail-container-healthcheckconfig-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The number of consecutive health check failures required before moving the container to the `Unhealthy` state.
	//
	// The default value is `2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-healthcheckconfig.html#cfn-lightsail-container-healthcheckconfig-unhealthythreshold
	//
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

