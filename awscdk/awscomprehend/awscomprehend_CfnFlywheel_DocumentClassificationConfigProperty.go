package awscomprehend


// Configuration required for a custom classification model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentClassificationConfigProperty := &documentClassificationConfigProperty{
//   	mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   }
//
type CfnFlywheel_DocumentClassificationConfigProperty struct {
	// Classification mode indicates whether the documents are `MULTI_CLASS` or `MULTI_LABEL` .
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// One or more labels to associate with the custom classifier.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

