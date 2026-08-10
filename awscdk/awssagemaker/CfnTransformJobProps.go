package awssagemaker


// Properties for defining a `CfnTransformJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransformJobProps := &CfnTransformJobProps{
//   	ModelName: jsii.String("modelName"),
//   	TransformInput: &TransformInputProperty{
//   		DataSource: &DataSourceProperty{
//   			S3DataSource: &S3DataSourceProperty{
//   				S3DataType: jsii.String("s3DataType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		CompressionType: jsii.String("compressionType"),
//   		ContentType: jsii.String("contentType"),
//   		SplitType: jsii.String("splitType"),
//   	},
//   	TransformOutput: &TransformOutputProperty{
//   		S3OutputPath: jsii.String("s3OutputPath"),
//
//   		// the properties below are optional
//   		Accept: jsii.String("accept"),
//   		AssembleWith: jsii.String("assembleWith"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	TransformResources: &TransformResourcesProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//
//   	// the properties below are optional
//   	BatchStrategy: jsii.String("batchStrategy"),
//   	DataCaptureConfig: &DataCaptureConfigProperty{
//   		DestinationS3Uri: jsii.String("destinationS3Uri"),
//
//   		// the properties below are optional
//   		GenerateInferenceId: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	DataProcessing: &DataProcessingProperty{
//   		InputFilter: jsii.String("inputFilter"),
//   		JoinSource: jsii.String("joinSource"),
//   		OutputFilter: jsii.String("outputFilter"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExperimentConfig: &ExperimentConfigProperty{
//   		ExperimentName: jsii.String("experimentName"),
//   		TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   		TrialName: jsii.String("trialName"),
//   	},
//   	MaxConcurrentTransforms: jsii.Number(123),
//   	MaxPayloadInMb: jsii.Number(123),
//   	ModelClientConfig: &ModelClientConfigProperty{
//   		InvocationsMaxRetries: jsii.Number(123),
//   		InvocationsTimeoutInSeconds: jsii.Number(123),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html
//
type CfnTransformJobProps struct {
	// The name of the model that you want to use for the transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-modelname
	//
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// Describes the input source and the way the transform job consumes it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-transforminput
	//
	TransformInput interface{} `field:"required" json:"transformInput" yaml:"transformInput"`
	// Describes the results of the transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-transformoutput
	//
	TransformOutput interface{} `field:"required" json:"transformOutput" yaml:"transformOutput"`
	// Describes the resources, including ML instance types and ML instance count, to use for the transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-transformresources
	//
	TransformResources interface{} `field:"required" json:"transformResources" yaml:"transformResources"`
	// Specifies the number of records to include in a mini-batch for an HTTP inference request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-batchstrategy
	//
	BatchStrategy *string `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// Configuration to control how SageMaker captures inference data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-datacaptureconfig
	//
	DataCaptureConfig interface{} `field:"optional" json:"dataCaptureConfig" yaml:"dataCaptureConfig"`
	// The data structure used to specify the data to be used for inference in a batch transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-dataprocessing
	//
	DataProcessing interface{} `field:"optional" json:"dataProcessing" yaml:"dataProcessing"`
	// The environment variables to set in the Docker container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Associates a SageMaker job as a trial component with an experiment and trial.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-experimentconfig
	//
	ExperimentConfig interface{} `field:"optional" json:"experimentConfig" yaml:"experimentConfig"`
	// The maximum number of parallel requests that can be sent to each instance in a transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-maxconcurrenttransforms
	//
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// The maximum allowed size of the payload, in MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-maxpayloadinmb
	//
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// Configures the timeout and maximum number of retries for processing a transform job invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-modelclientconfig
	//
	ModelClientConfig interface{} `field:"optional" json:"modelClientConfig" yaml:"modelClientConfig"`
	// An array of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html#cfn-sagemaker-transformjob-tags
	//
	Tags *[]*CfnTransformJob_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

