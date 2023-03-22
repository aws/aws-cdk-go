package awsm2


// Defines the details of a high availability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   highAvailabilityConfigProperty := &highAvailabilityConfigProperty{
//   	desiredCapacity: jsii.Number(123),
//   }
//
type CfnEnvironment_HighAvailabilityConfigProperty struct {
	// The number of instances in a high availability configuration.
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
}

