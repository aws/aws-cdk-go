package awsstepfunctionstasks


// Specify the training algorithm and algorithm-specific metadata.
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
type AlgorithmSpecification struct {
	// Name of the algorithm resource to use for the training job.
	//
	// This must be an algorithm resource that you created or subscribe to on AWS Marketplace.
	// If you specify a value for this parameter, you can't specify a value for TrainingImage.
	AlgorithmName *string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// List of metric definition objects.
	//
	// Each object specifies the metric name and regular expressions used to parse algorithm logs.
	MetricDefinitions *[]*MetricDefinition `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Registry path of the Docker image that contains the training algorithm.
	TrainingImage DockerImage `field:"optional" json:"trainingImage" yaml:"trainingImage"`
	// Input mode that the algorithm supports.
	TrainingInputMode InputMode `field:"optional" json:"trainingInputMode" yaml:"trainingInputMode"`
}

