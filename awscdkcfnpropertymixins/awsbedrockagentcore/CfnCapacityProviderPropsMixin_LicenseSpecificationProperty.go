package awsbedrockagentcore


// Describes a license configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   licenseSpecificationProperty := &LicenseSpecificationProperty{
//   	LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-licensespecification.html
//
type CfnCapacityProviderPropsMixin_LicenseSpecificationProperty struct {
	// The ARN of the license configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-licensespecification.html#cfn-bedrockagentcore-capacityprovider-licensespecification-licenseconfigurationarn
	//
	LicenseConfigurationArn *string `field:"optional" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

