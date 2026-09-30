package awscomprehend


// Specifies the format and location of the input data for an entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   entityRecognizerInputDataConfigProperty := &EntityRecognizerInputDataConfigProperty{
//   	Annotations: &EntityRecognizerAnnotationsProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	AugmentedManifests: []interface{}{
//   		&AugmentedManifestsListItemProperty{
//   			AnnotationDataS3Uri: jsii.String("annotationDataS3Uri"),
//   			AttributeNames: []*string{
//   				jsii.String("attributeNames"),
//   			},
//   			DocumentType: jsii.String("documentType"),
//   			S3Uri: jsii.String("s3Uri"),
//   			SourceDocumentsS3Uri: jsii.String("sourceDocumentsS3Uri"),
//   			Split: jsii.String("split"),
//   		},
//   	},
//   	DataFormat: jsii.String("dataFormat"),
//   	Documents: &EntityRecognizerDocumentsProperty{
//   		InputFormat: jsii.String("inputFormat"),
//   		S3Uri: jsii.String("s3Uri"),
//   		TestS3Uri: jsii.String("testS3Uri"),
//   	},
//   	EntityList: &EntityRecognizerEntityListProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	EntityTypes: []interface{}{
//   		&EntityTypesListItemProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html
//
type CfnEntityRecognizerPropsMixin_EntityRecognizerInputDataConfigProperty struct {
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
	// The entity types in the labeled training data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entityrecognizerinputdataconfig.html#cfn-comprehend-entityrecognizer-entityrecognizerinputdataconfig-entitytypes
	//
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
}

