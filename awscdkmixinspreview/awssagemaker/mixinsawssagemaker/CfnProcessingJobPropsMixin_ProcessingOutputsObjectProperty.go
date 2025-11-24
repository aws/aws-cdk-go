package mixinsawssagemaker


// Describes the results of a processing job.
//
// The processing output must specify exactly one of either `S3Output` or `FeatureStoreOutput` types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   processingOutputsObjectProperty := &ProcessingOutputsObjectProperty{
//   	AppManaged: jsii.Boolean(false),
//   	FeatureStoreOutput: &FeatureStoreOutputProperty{
//   		FeatureGroupName: jsii.String("featureGroupName"),
//   	},
//   	OutputName: jsii.String("outputName"),
//   	S3Output: &S3OutputProperty{
//   		LocalPath: jsii.String("localPath"),
//   		S3UploadMode: jsii.String("s3UploadMode"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputsobject.html
//
type CfnProcessingJobPropsMixin_ProcessingOutputsObjectProperty struct {
	// When `True` , output operations such as data upload are managed natively by the processing job application.
	//
	// When `False` (default), output operations are managed by Amazon SageMaker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputsobject.html#cfn-sagemaker-processingjob-processingoutputsobject-appmanaged
	//
	AppManaged interface{} `field:"optional" json:"appManaged" yaml:"appManaged"`
	// Configuration for processing job outputs in Amazon SageMaker Feature Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputsobject.html#cfn-sagemaker-processingjob-processingoutputsobject-featurestoreoutput
	//
	FeatureStoreOutput interface{} `field:"optional" json:"featureStoreOutput" yaml:"featureStoreOutput"`
	// The name for the processing job output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputsobject.html#cfn-sagemaker-processingjob-processingoutputsobject-outputname
	//
	OutputName *string `field:"optional" json:"outputName" yaml:"outputName"`
	// Configuration for uploading output data to Amazon S3 from the processing container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputsobject.html#cfn-sagemaker-processingjob-processingoutputsobject-s3output
	//
	S3Output interface{} `field:"optional" json:"s3Output" yaml:"s3Output"`
}

