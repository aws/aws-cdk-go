package awsrum


// A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   javaScriptSourceMapsProperty := &JavaScriptSourceMapsProperty{
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-javascriptsourcemaps.html
//
type CfnAppMonitor_JavaScriptSourceMapsProperty struct {
	// Specifies whether JavaScript error stack traces should be unminified for this app monitor.
	//
	// The default is for JavaScript error stack trace unminification to be DISABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-javascriptsourcemaps.html#cfn-rum-appmonitor-javascriptsourcemaps-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// The S3Uri of the bucket or folder that stores the source map files.
	//
	// It is required if status is ENABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-javascriptsourcemaps.html#cfn-rum-appmonitor-javascriptsourcemaps-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

