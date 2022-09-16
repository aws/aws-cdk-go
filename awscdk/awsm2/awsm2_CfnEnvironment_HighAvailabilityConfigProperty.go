package awsm2


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
	// `CfnEnvironment.HighAvailabilityConfigProperty.DesiredCapacity`.
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
}

