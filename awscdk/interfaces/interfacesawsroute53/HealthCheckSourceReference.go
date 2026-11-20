package interfacesawsroute53


// A reference to a HealthCheckSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckSourceReference := &HealthCheckSourceReference{
//   	HealthCheckSourceArn: jsii.String("healthCheckSourceArn"),
//   }
//
type HealthCheckSourceReference struct {
	// The Arn of the HealthCheckSource resource.
	HealthCheckSourceArn *string `field:"required" json:"healthCheckSourceArn" yaml:"healthCheckSourceArn"`
}

