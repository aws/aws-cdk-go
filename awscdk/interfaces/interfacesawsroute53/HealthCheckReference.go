package interfacesawsroute53


// A reference to a HealthCheck resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckReference := &HealthCheckReference{
//   	HealthCheckId: jsii.String("healthCheckId"),
//   }
//
type HealthCheckReference struct {
	// The HealthCheckId of the HealthCheck resource.
	HealthCheckId *string `field:"required" json:"healthCheckId" yaml:"healthCheckId"`
}

