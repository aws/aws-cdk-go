package previewawsinspectorv2mixins


// Configuration settings for continuous integration scans that run automatically when code changes are made.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   continuousIntegrationScanConfigurationProperty := &ContinuousIntegrationScanConfigurationProperty{
//   	SupportedEvents: []*string{
//   		jsii.String("supportedEvents"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration.html
//
type CfnCodeSecurityScanConfigurationPropsMixin_ContinuousIntegrationScanConfigurationProperty struct {
	// The repository events that trigger continuous integration scans, such as pull requests or commits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-continuousintegrationscanconfiguration-supportedevents
	//
	SupportedEvents *[]*string `field:"optional" json:"supportedEvents" yaml:"supportedEvents"`
}

