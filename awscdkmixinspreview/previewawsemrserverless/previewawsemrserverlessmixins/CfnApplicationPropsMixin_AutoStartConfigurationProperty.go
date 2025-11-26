package previewawsemrserverlessmixins


// The configuration for an application to automatically start on job submission.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoStartConfigurationProperty := &AutoStartConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostartconfiguration.html
//
type CfnApplicationPropsMixin_AutoStartConfigurationProperty struct {
	// Enables the application to automatically start on job submission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostartconfiguration.html#cfn-emrserverless-application-autostartconfiguration-enabled
	//
	// Default: - true.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

