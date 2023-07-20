package awsnimblestudio


// The configuration for a license service that is associated with a studio resource.
//
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
	// The endpoint of the license service that is accessed by the studio component resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-licenseserviceconfiguration.html#cfn-nimblestudio-studiocomponent-licenseserviceconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

