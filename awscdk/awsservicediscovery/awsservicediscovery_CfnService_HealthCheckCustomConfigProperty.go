package awsservicediscovery


// A complex type that contains information about an optional custom health check.
//
// A custom health check, which requires that you use a third-party health checker to evaluate the health of your resources, is useful in the following circumstances:
//
// - You can't use a health check that's defined by `HealthCheckConfig` because the resource isn't available over the internet. For example, you can use a custom health check when the instance is in an Amazon VPC. (To check the health of resources in a VPC, the health checker must also be in the VPC.)
// - You want to use a third-party health checker regardless of where your resources are located.
//
// > If you specify a health check configuration, you can specify either `HealthCheckCustomConfig` or `HealthCheckConfig` but not both.
//
// To change the status of a custom health check, submit an `UpdateInstanceCustomHealthStatus` request. AWS Cloud Map doesn't monitor the status of the resource, it just keeps a record of the status specified in the most recent `UpdateInstanceCustomHealthStatus` request.
//
// Here's how custom health checks work:
//
// - You create a service.
// - You register an instance.
// - You configure a third-party health checker to monitor the resource that's associated with the new instance.
//
// > AWS Cloud Map doesn't check the health of the resource directly.
// - The third-party health-checker determines that the resource is unhealthy and notifies your application.
// - Your application submits an `UpdateInstanceCustomHealthStatus` request.
// - AWS Cloud Map waits for 30 seconds.
// - If another `UpdateInstanceCustomHealthStatus` request doesn't arrive during that time to change the status back to healthy, AWS Cloud Map stops routing traffic to the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckCustomConfigProperty := &healthCheckCustomConfigProperty{
//   	failureThreshold: jsii.Number(123),
//   }
//
type CfnService_HealthCheckCustomConfigProperty struct {
	// > This parameter is no longer supported and is always set to 1.
	//
	// AWS Cloud Map waits for approximately 30 seconds after receiving an `UpdateInstanceCustomHealthStatus` request before changing the status of the service instance.
	//
	// The number of 30-second intervals that you want AWS Cloud Map to wait after receiving an `UpdateInstanceCustomHealthStatus` request before it changes the health status of a service instance.
	//
	// Sending a second or subsequent `UpdateInstanceCustomHealthStatus` request with the same value before 30 seconds has passed doesn't accelerate the change. AWS Cloud Map still waits `30` seconds after the first request to make the change.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
}

