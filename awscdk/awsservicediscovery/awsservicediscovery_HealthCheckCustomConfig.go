package awsservicediscovery


// Specifies information about an optional custom health check.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckCustomConfig := &healthCheckCustomConfig{
//   	failureThreshold: jsii.Number(123),
//   }
//
type HealthCheckCustomConfig struct {
	// The number of 30-second intervals that you want Cloud Map to wait after receiving an UpdateInstanceCustomHealthStatus request before it changes the health status of a service instance.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
}

