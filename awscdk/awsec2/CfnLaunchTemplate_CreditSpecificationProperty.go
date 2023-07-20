package awsec2


// Specifies the credit option for CPU usage of a T2, T3, or T3a instance.
//
// `CreditSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-creditspecification.html
//
type CfnLaunchTemplate_CreditSpecificationProperty struct {
	// The credit option for CPU usage of a T instance.
	//
	// Valid values: `standard` | `unlimited`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-creditspecification.html#cfn-ec2-launchtemplate-creditspecification-cpucredits
	//
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

