package awscomprehend


// An augmented manifest file that provides training data for your custom model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   augmentedManifestsListItemProperty := &AugmentedManifestsListItemProperty{
//   	AnnotationDataS3Uri: jsii.String("annotationDataS3Uri"),
//   	AttributeNames: []*string{
//   		jsii.String("attributeNames"),
//   	},
//   	DocumentType: jsii.String("documentType"),
//   	S3Uri: jsii.String("s3Uri"),
//   	SourceDocumentsS3Uri: jsii.String("sourceDocumentsS3Uri"),
//   	Split: jsii.String("split"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html
//
type CfnEntityRecognizerPropsMixin_AugmentedManifestsListItemProperty struct {
	// The S3 prefix to the annotation files that are referred in the augmented manifest file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html#cfn-comprehend-entityrecognizer-augmentedmanifestslistitem-annotationdatas3uri
	//
	AnnotationDataS3Uri *string `field:"optional" json:"annotationDataS3Uri" yaml:"annotationDataS3Uri"`
	// The JSON attribute that contains the annotations for your training documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html#cfn-comprehend-entityrecognizer-augmentedmanifestslistitem-attributenames
	//
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// The type of augmented manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html#cfn-comprehend-entityrecognizer-augmentedmanifestslistitem-documenttype
	//
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// The Amazon S3 location of the augmented manifest file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html#cfn-comprehend-entityrecognizer-augmentedmanifestslistitem-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// The S3 prefix to the source files (PDFs) that are referred to in the augmented manifest file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html#cfn-comprehend-entityrecognizer-augmentedmanifestslistitem-sourcedocumentss3uri
	//
	SourceDocumentsS3Uri *string `field:"optional" json:"sourceDocumentsS3Uri" yaml:"sourceDocumentsS3Uri"`
	// The purpose of the data you've provided in the augmented manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-augmentedmanifestslistitem.html#cfn-comprehend-entityrecognizer-augmentedmanifestslistitem-split
	//
	Split *string `field:"optional" json:"split" yaml:"split"`
}

