package previewawssyntheticsmixins


// A structure that specifies the browser type to use for a canary run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   browserConfigProperty := &BrowserConfigProperty{
//   	BrowserType: jsii.String("browserType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-browserconfig.html
//
type CfnCanaryPropsMixin_BrowserConfigProperty struct {
	// The browser type associated with this browser configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-browserconfig.html#cfn-synthetics-canary-browserconfig-browsertype
	//
	BrowserType *string `field:"optional" json:"browserType" yaml:"browserType"`
}

