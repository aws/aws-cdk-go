package awsec2


// Specifies the license configuration to use.
//
// `LicenseSpecification` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseSpecificationProperty := &licenseSpecificationProperty{
//   	licenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   }
//
type CfnInstance_LicenseSpecificationProperty struct {
	// The Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string `field:"required" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

