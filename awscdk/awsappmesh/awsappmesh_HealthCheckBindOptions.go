package awsappmesh


// Options used for creating the Health Check object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckBindOptions := &healthCheckBindOptions{
//   	defaultPort: jsii.Number(123),
//   }
//
type HealthCheckBindOptions struct {
	// Port for Health Check interface.
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
}

