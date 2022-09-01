package awsstepfunctionstasks


// S3 Data Type.
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
type S3DataType string

const (
	// Manifest File Data Type.
	S3DataType_MANIFEST_FILE S3DataType = "MANIFEST_FILE"
	// S3 Prefix Data Type.
	S3DataType_S3_PREFIX S3DataType = "S3_PREFIX"
	// Augmented Manifest File Data Type.
	S3DataType_AUGMENTED_MANIFEST_FILE S3DataType = "AUGMENTED_MANIFEST_FILE"
)

