package awsec2


// Specifies the credit option for CPU usage of a T2, T3, or T3a instance.
//
// `CreditSpecification` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   creditSpecificationProperty := &creditSpecificationProperty{
//   	cpuCredits: jsii.String("cpuCredits"),
//   }
//
type CfnInstance_CreditSpecificationProperty struct {
	// The credit option for CPU usage of the instance.
	//
	// Valid values are `standard` and `unlimited` . `T3` instances launch as `unlimited` by default. `T2` instances launch as `standard` by default.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

