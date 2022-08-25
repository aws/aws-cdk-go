package awsec2


// Specifies a license configuration for an instance.
//
// `LicenseSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
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
type CfnLaunchTemplate_LicenseSpecificationProperty struct {
	// The Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string `field:"optional" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

