package previewawswisdommixins


// The content of the quick response stored in different media types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   quickResponseContentsProperty := &QuickResponseContentsProperty{
//   	Markdown: &QuickResponseContentProviderProperty{
//   		Content: jsii.String("content"),
//   	},
//   	PlainText: &QuickResponseContentProviderProperty{
//   		Content: jsii.String("content"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-quickresponse-quickresponsecontents.html
//
type CfnQuickResponsePropsMixin_QuickResponseContentsProperty struct {
	// The quick response content in markdown format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-quickresponse-quickresponsecontents.html#cfn-wisdom-quickresponse-quickresponsecontents-markdown
	//
	Markdown interface{} `field:"optional" json:"markdown" yaml:"markdown"`
	// The quick response content in plaintext format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-quickresponse-quickresponsecontents.html#cfn-wisdom-quickresponse-quickresponsecontents-plaintext
	//
	PlainText interface{} `field:"optional" json:"plainText" yaml:"plainText"`
}

