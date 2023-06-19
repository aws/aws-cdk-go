package awsstepfunctionstasks


// Dataset to be transformed and the Amazon S3 location where it is stored.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &SageMakerCreateTransformJobProps{
//   	TransformJobName: jsii.String("MyTransformJob"),
//   	ModelName: jsii.String("MyModelName"),
//   	ModelClientOptions: &ModelClientOptions{
//   		InvocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		InvocationsTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	},
//   	TransformInput: &TransformInput{
//   		TransformDataSource: &TransformDataSource{
//   			S3DataSource: &TransformS3DataSource{
//   				S3Uri: jsii.String("s3://inputbucket/train"),
//   				S3DataType: tasks.S3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	TransformOutput: &TransformOutput{
//   		S3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	TransformResources: &TransformResources{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M4, ec2.InstanceSize_XLARGE),
//   	},
//   })
//
// Experimental.
type TransformInput struct {
	// S3 location of the channel data.
	// Experimental.
	TransformDataSource *TransformDataSource `field:"required" json:"transformDataSource" yaml:"transformDataSource"`
	// The compression type of the transform data.
	// Experimental.
	CompressionType CompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Multipurpose internet mail extension (MIME) type of the data.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Method to use to split the transform job's data files into smaller batches.
	// Experimental.
	SplitType SplitType `field:"optional" json:"splitType" yaml:"splitType"`
}

