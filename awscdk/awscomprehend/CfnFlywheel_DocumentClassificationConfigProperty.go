package awscomprehend


// Configuration required for a document classification model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassificationConfigProperty := &DocumentClassificationConfigProperty{
//   	Mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-documentclassificationconfig.html
//
type CfnFlywheel_DocumentClassificationConfigProperty struct {
	// Classification mode indicates whether the documents are `MULTI_CLASS` or `MULTI_LABEL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-documentclassificationconfig.html#cfn-comprehend-flywheel-documentclassificationconfig-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// One or more labels to associate with the custom classifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-documentclassificationconfig.html#cfn-comprehend-flywheel-documentclassificationconfig-labels
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

