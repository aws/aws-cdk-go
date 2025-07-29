package awsrum


// A structure that contains the configuration for how an app monitor can deobfuscate stack traces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deobfuscationConfigurationProperty := &DeobfuscationConfigurationProperty{
//   	JavaScriptSourceMaps: &JavaScriptSourceMapsProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-deobfuscationconfiguration.html
//
type CfnAppMonitor_DeobfuscationConfigurationProperty struct {
	// A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-deobfuscationconfiguration.html#cfn-rum-appmonitor-deobfuscationconfiguration-javascriptsourcemaps
	//
	JavaScriptSourceMaps interface{} `field:"optional" json:"javaScriptSourceMaps" yaml:"javaScriptSourceMaps"`
}

