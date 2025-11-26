package previewawswisdommixins


// The container quick response content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   quickResponseContentProviderProperty := &QuickResponseContentProviderProperty{
//   	Content: jsii.String("content"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-quickresponse-quickresponsecontentprovider.html
//
type CfnQuickResponsePropsMixin_QuickResponseContentProviderProperty struct {
	// The content of the quick response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-quickresponse-quickresponsecontentprovider.html#cfn-wisdom-quickresponse-quickresponsecontentprovider-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
}

