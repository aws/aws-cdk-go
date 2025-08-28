package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousIntegrationScanConfigurationProperty := &ContinuousIntegrationScanConfigurationProperty{
//   	SupportedEvents: []*string{
//   		jsii.String("supportedEvents"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration.html
//
type CfnCodeSecurityScanConfiguration_ContinuousIntegrationScanConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration-supportedevents
	//
	SupportedEvents *[]*string `field:"required" json:"supportedEvents" yaml:"supportedEvents"`
}

