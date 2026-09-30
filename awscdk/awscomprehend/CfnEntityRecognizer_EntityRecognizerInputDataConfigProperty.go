package awscomprehend


// Specifies the format and location of the input data for an entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityRecognizerInputDataConfigProperty := &EntityRecognizerInputDataConfigProperty{
//   	EntityTypes: []interface{}{
//   		&EntityTypesListItemProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Annotations: &EntityRecognizerAnnotationsProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	AugmentedManifests: []interface{}{
//   		&AugmentedManifestsListItemProperty{
//   			AttributeNames: []*string{
//   				jsii.String("attributeNames"),
//   			},
//   			S3Uri: jsii.String("s3Uri"),
//
//   			// the properties below are optional
//   			AnnotationDataS3Uri: jsii.String("annotationDataS3Uri"),
//   			DocumentType: jsii.String("documentType"),
//   			SourceDocumentsS3Uri: jsii.String("sourceDocumentsS3Uri"),
//   			Split: jsii.String("split"),
//   		},
//   	},
//   	DataFormat: jsii.String("dataFormat"),
//   	Documents: &EntityRecognizerDocumentsProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		InputFormat: jsii.String("inputFormat"),
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	EntityList: &EntityRecognizerEntityListProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html
//
type CfnEntityRecognizer_EntityRecognizerInputDataConfigProperty struct {
	// The entity types in the labeled training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-entitytypes
	//
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// Describes the annotations associated with an entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-annotations
	//
	Annotations interface{} `field:"optional" json:"annotations" yaml:"annotations"`
	// A list of augmented manifest files that provide training data for a custom model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-augmentedmanifests
	//
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// The format of your training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-dataformat
	//
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Describes the training documents submitted with an entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-documents
	//
	Documents interface{} `field:"optional" json:"documents" yaml:"documents"`
	// Describes the entity list submitted with an entity recognizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-entitylist
	//
	EntityList interface{} `field:"optional" json:"entityList" yaml:"entityList"`
}

