package previewawsbedrockagentcoremixins


// Browser signing configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   browserSigningProperty := &BrowserSigningProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsersigning.html
//
type CfnBrowserCustomPropsMixin_BrowserSigningProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsersigning.html#cfn-bedrockagentcore-browsercustom-browsersigning-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

