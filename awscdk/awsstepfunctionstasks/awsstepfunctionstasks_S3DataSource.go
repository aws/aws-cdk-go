package awsstepfunctionstasks


// S3 location of the channel data.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &sageMakerCreateTrainingJobProps{
//   	trainingJobName: sfn.jsonPath.stringAt(jsii.String("$.JobName")),
//   	algorithmSpecification: &algorithmSpecification{
//   		algorithmName: jsii.String("BlazingText"),
//   		trainingInputMode: tasks.inputMode_FILE,
//   	},
//   	inputDataConfig: []channel{
//   		&channel{
//   			channelName: jsii.String("train"),
//   			dataSource: &dataSource{
//   				s3DataSource: &s3DataSource{
//   					s3DataType: tasks.s3DataType_S3_PREFIX,
//   					s3Location: tasks.s3Location.fromJsonExpression(jsii.String("$.S3Bucket")),
//   				},
//   			},
//   		},
//   	},
//   	outputDataConfig: &outputDataConfig{
//   		s3OutputLocation: tasks.*s3Location.fromBucket(s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
//   	},
//   	resourceConfig: &resourceConfig{
//   		instanceCount: jsii.Number(1),
//   		instanceType: ec2.NewInstanceType(sfn.*jsonPath.stringAt(jsii.String("$.InstanceType"))),
//   		volumeSize: awscdk.Size.gibibytes(jsii.Number(50)),
//   	},
//   	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
//   	stoppingCondition: &stoppingCondition{
//   		maxRuntime: awscdk.Duration.hours(jsii.Number(2)),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_S3DataSource.html
//
type S3DataSource struct {
	// S3 Uri.
	S3Location S3Location `field:"required" json:"s3Location" yaml:"s3Location"`
	// List of one or more attribute names to use that are found in a specified augmented manifest file.
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// S3 Data Distribution Type.
	S3DataDistributionType S3DataDistributionType `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// S3 Data Type.
	S3DataType S3DataType `field:"optional" json:"s3DataType" yaml:"s3DataType"`
}

