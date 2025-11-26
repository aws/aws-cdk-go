package previewawsservicecatalogmixins


// Properties for CfnTagOptionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTagOptionMixinProps := &CfnTagOptionMixinProps{
//   	Active: jsii.Boolean(false),
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html
//
type CfnTagOptionMixinProps struct {
	// The TagOption active state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html#cfn-servicecatalog-tagoption-active
	//
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// The TagOption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html#cfn-servicecatalog-tagoption-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The TagOption value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html#cfn-servicecatalog-tagoption-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

