package previewawscomprehendmixins


// Configuration required for a document classification model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentClassificationConfigProperty := &DocumentClassificationConfigProperty{
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-documentclassificationconfig.html
//
type CfnFlywheelPropsMixin_DocumentClassificationConfigProperty struct {
	// One or more labels to associate with the custom classifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-documentclassificationconfig.html#cfn-comprehend-flywheel-documentclassificationconfig-labels
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Classification mode indicates whether the documents are `MULTI_CLASS` or `MULTI_LABEL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-documentclassificationconfig.html#cfn-comprehend-flywheel-documentclassificationconfig-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

