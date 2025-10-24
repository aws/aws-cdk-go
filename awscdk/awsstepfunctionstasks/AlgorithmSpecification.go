package awsstepfunctionstasks


// Specify the training algorithm and algorithm-specific metadata.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &SageMakerCreateTrainingJobProps{
//   	TrainingJobName: sfn.JsonPath_StringAt(jsii.String("$.JobName")),
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		AlgorithmName: jsii.String("BlazingText"),
//   		TrainingInputMode: tasks.InputMode_FILE,
//   	},
//   	InputDataConfig: []Channel{
//   		&Channel{
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
type AlgorithmSpecification struct {
	// Name of the algorithm resource to use for the training job.
	//
	// This must be an algorithm resource that you created or subscribe to on AWS Marketplace.
	// If you specify a value for this parameter, you can't specify a value for TrainingImage.
	// Default: - No algorithm is specified.
	//
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// List of metric definition objects.
	//
	// Each object specifies the metric name and regular expressions used to parse algorithm logs.
	// Default: - No metrics.
	//
	MetricDefinitions *[]*MetricDefinition `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Registry path of the Docker image that contains the training algorithm.
	// Default: - No Docker image is specified.
	//
	TrainingImage DockerImage `field:"optional" json:"trainingImage" yaml:"trainingImage"`
	// Input mode that the algorithm supports.
	// Default: 'File' mode.
	//
	TrainingInputMode InputMode `field:"optional" json:"trainingInputMode" yaml:"trainingInputMode"`
}

