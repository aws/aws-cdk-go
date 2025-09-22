package awsstepfunctionstasks


// S3 location of the channel data.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &SageMakerCreateTrainingJobProps{
//   	TrainingJobName: sfn.JsonPath_StringAt(jsii.String("$.JobName")),
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		AlgorithmName: jsii.String("BlazingText"),
//   		TrainingInputMode: tasks.InputMode_FILE,
//   	},
//   	InputDataConfig: []channel{
//   		&channel{
//   			ChannelName: jsii.String("train"),
//   			DataSource: &DataSource{
//   				S3DataSource: &S3DataSource{
//   					S3DataType: tasks.S3DataType_S3_PREFIX,
//   					S3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.S3Bucket")),
//   				},
//   			},
//   		},
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		S3OutputLocation: tasks.S3Location_FromBucket(s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket")), jsii.String("myoutputpath")),
//   	},
//   	ResourceConfig: &ResourceConfig{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.NewInstanceType(sfn.JsonPath_*StringAt(jsii.String("$.InstanceType"))),
//   		VolumeSize: awscdk.Size_Gibibytes(jsii.Number(50)),
//   	},
//   	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntime: awscdk.Duration_Hours(jsii.Number(2)),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_S3DataSource.html
//
type S3DataSource struct {
	// S3 Uri.
	S3Location S3Location `field:"required" json:"s3Location" yaml:"s3Location"`
	// List of one or more attribute names to use that are found in a specified augmented manifest file.
	// Default: - No attribute names.
	//
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// S3 Data Distribution Type.
	// Default: - None.
	//
	S3DataDistributionType S3DataDistributionType `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// S3 Data Type.
	// Default: S3_PREFIX.
	//
	S3DataType S3DataType `field:"optional" json:"s3DataType" yaml:"s3DataType"`
}

