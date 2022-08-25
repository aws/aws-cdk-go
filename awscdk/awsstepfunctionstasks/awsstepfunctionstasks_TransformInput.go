package awsstepfunctionstasks


// Dataset to be transformed and the Amazon S3 location where it is stored.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &sageMakerCreateTransformJobProps{
//   	transformJobName: jsii.String("MyTransformJob"),
//   	modelName: jsii.String("MyModelName"),
//   	modelClientOptions: &modelClientOptions{
//   		invocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		invocationsTimeout: awscdk.Duration.minutes(jsii.Number(5)),
//   	},
//   	transformInput: &transformInput{
//   		transformDataSource: &transformDataSource{
//   			s3DataSource: &transformS3DataSource{
//   				s3Uri: jsii.String("s3://inputbucket/train"),
//   				s3DataType: tasks.s3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	transformOutput: &transformOutput{
//   		s3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	transformResources: &transformResources{
//   		instanceCount: jsii.Number(1),
//   		instanceType: ec2.instanceType.of(ec2.instanceClass_M4, ec2.instanceSize_XLARGE),
//   	},
//   })
//
type TransformInput struct {
	// S3 location of the channel data.
	TransformDataSource *TransformDataSource `field:"required" json:"transformDataSource" yaml:"transformDataSource"`
	// The compression type of the transform data.
	CompressionType CompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Multipurpose internet mail extension (MIME) type of the data.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Method to use to split the transform job's data files into smaller batches.
	SplitType SplitType `field:"optional" json:"splitType" yaml:"splitType"`
}

