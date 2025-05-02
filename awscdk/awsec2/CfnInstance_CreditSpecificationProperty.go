package awsec2


// Specifies the credit option for CPU usage of a T instance.
//
// `CreditSpecification` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   creditSpecificationProperty := &CreditSpecificationProperty{
//   	CpuCredits: jsii.String("cpuCredits"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-creditspecification.html
//
type CfnInstance_CreditSpecificationProperty struct {
	// The credit option for CPU usage of the instance.
	//
	// Valid values: `standard` | `unlimited`
	//
	// T3 instances with `host` tenancy do not support the `unlimited` CPU credit option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-creditspecification.html#cfn-ec2-instance-creditspecification-cpucredits
	//
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

