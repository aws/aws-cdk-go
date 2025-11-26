package previewawsrummixins


// A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   javaScriptSourceMapsProperty := &JavaScriptSourceMapsProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-javascriptsourcemaps.html
//
type CfnAppMonitorPropsMixin_JavaScriptSourceMapsProperty struct {
	// The S3Uri of the bucket or folder that stores the source map files.
	//
	// It is required if status is ENABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-javascriptsourcemaps.html#cfn-rum-appmonitor-javascriptsourcemaps-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// Specifies whether JavaScript error stack traces should be unminified for this app monitor.
	//
	// The default is for JavaScript error stack trace unminification to be `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-javascriptsourcemaps.html#cfn-rum-appmonitor-javascriptsourcemaps-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

