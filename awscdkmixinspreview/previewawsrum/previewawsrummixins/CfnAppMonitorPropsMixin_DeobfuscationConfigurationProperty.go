package previewawsrummixins


// A structure that contains the configuration for how an app monitor can deobfuscate stack traces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deobfuscationConfigurationProperty := &DeobfuscationConfigurationProperty{
//   	JavaScriptSourceMaps: &JavaScriptSourceMapsProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-deobfuscationconfiguration.html
//
type CfnAppMonitorPropsMixin_DeobfuscationConfigurationProperty struct {
	// A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-deobfuscationconfiguration.html#cfn-rum-appmonitor-deobfuscationconfiguration-javascriptsourcemaps
	//
	JavaScriptSourceMaps interface{} `field:"optional" json:"javaScriptSourceMaps" yaml:"javaScriptSourceMaps"`
}

