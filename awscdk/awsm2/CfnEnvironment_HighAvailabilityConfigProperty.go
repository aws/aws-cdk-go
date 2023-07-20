package awsm2


// Defines the details of a high availability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   highAvailabilityConfigProperty := &HighAvailabilityConfigProperty{
//   	DesiredCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-highavailabilityconfig.html
//
type CfnEnvironment_HighAvailabilityConfigProperty struct {
	// The number of instances in a high availability configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-highavailabilityconfig.html#cfn-m2-environment-highavailabilityconfig-desiredcapacity
	//
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
}

