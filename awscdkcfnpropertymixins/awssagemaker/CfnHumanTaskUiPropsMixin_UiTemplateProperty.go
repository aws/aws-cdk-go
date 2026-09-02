package awssagemaker


// The Liquid template for the worker user interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   uiTemplateProperty := &UiTemplateProperty{
//   	Content: jsii.String("content"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-humantaskui-uitemplate.html
//
type CfnHumanTaskUiPropsMixin_UiTemplateProperty struct {
	// The content of the Liquid template for the worker user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-humantaskui-uitemplate.html#cfn-sagemaker-humantaskui-uitemplate-content
	//
	Content *string `field:"optional" json:"content" yaml:"content"`
}

