package previewawscleanroomsmixins


// A structure that defines the level of detail included in error messages returned by PySpark jobs.
//
// This configuration allows you to control the verbosity of error messages to help with troubleshooting PySpark jobs while maintaining appropriate security controls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   errorMessageConfigurationProperty := &ErrorMessageConfigurationProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-errormessageconfiguration.html
//
type CfnAnalysisTemplatePropsMixin_ErrorMessageConfigurationProperty struct {
	// The level of detail for error messages returned by the PySpark job.
	//
	// When set to DETAILED, error messages include more information to help troubleshoot issues with your PySpark job.
	//
	// Because this setting may expose sensitive data, it is recommended for development and testing environments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-errormessageconfiguration.html#cfn-cleanrooms-analysistemplate-errormessageconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

