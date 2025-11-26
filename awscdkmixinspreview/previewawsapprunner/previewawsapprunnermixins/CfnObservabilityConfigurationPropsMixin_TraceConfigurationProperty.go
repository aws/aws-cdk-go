package previewawsapprunnermixins


// Describes the configuration of the tracing feature within an AWS App Runner observability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   traceConfigurationProperty := &TraceConfigurationProperty{
//   	Vendor: jsii.String("vendor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-observabilityconfiguration-traceconfiguration.html
//
type CfnObservabilityConfigurationPropsMixin_TraceConfigurationProperty struct {
	// The implementation provider chosen for tracing App Runner services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-observabilityconfiguration-traceconfiguration.html#cfn-apprunner-observabilityconfiguration-traceconfiguration-vendor
	//
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

