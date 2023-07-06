package awscomprehend


// An augmented manifest file that provides training data for your custom model.
//
// An augmented manifest file is a labeled dataset that is produced by Amazon SageMaker Ground Truth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   augmentedManifestsListItemProperty := &AugmentedManifestsListItemProperty{
//   	AttributeNames: []*string{
//   		jsii.String("attributeNames"),
//   	},
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	Split: jsii.String("split"),
//   }
//
type CfnDocumentClassifier_AugmentedManifestsListItemProperty struct {
	// The JSON attribute that contains the annotations for your training documents.
	//
	// The number of attribute names that you specify depends on whether your augmented manifest file is the output of a single labeling job or a chained labeling job.
	//
	// If your file is the output of a single labeling job, specify the LabelAttributeName key that was used when the job was created in Ground Truth.
	//
	// If your file is the output of a chained labeling job, specify the LabelAttributeName key for one or more jobs in the chain. Each LabelAttributeName key provides the annotations from an individual job.
	AttributeNames *[]*string `field:"required" json:"attributeNames" yaml:"attributeNames"`
	// The Amazon S3 location of the augmented manifest file.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The purpose of the data you've provided in the augmented manifest.
	//
	// You can either train or test this data. If you don't specify, the default is train.
	//
	// TRAIN - all of the documents in the manifest will be used for training. If no test documents are provided, Amazon Comprehend will automatically reserve a portion of the training documents for testing.
	//
	// TEST - all of the documents in the manifest will be used for testing.
	Split *string `field:"optional" json:"split" yaml:"split"`
}

