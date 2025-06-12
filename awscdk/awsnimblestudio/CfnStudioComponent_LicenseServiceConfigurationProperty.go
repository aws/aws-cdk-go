package awsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseServiceConfigurationProperty := &LicenseServiceConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-licenseserviceconfiguration.html
//
type CfnStudioComponent_LicenseServiceConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-licenseserviceconfiguration.html#cfn-nimblestudio-studiocomponent-licenseserviceconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

