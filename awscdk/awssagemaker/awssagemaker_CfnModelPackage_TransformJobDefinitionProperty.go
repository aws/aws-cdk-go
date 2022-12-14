package awssagemaker


// Defines the input needed to run a transform job using the inference specification specified in the algorithm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformJobDefinitionProperty := &transformJobDefinitionProperty{
//   	transformInput: &transformInputProperty{
//   		dataSource: &dataSourceProperty{
//   			s3DataSource: &s3DataSourceProperty{
//   				s3DataType: jsii.String("s3DataType"),
//   				s3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		compressionType: jsii.String("compressionType"),
//   		contentType: jsii.String("contentType"),
//   		splitType: jsii.String("splitType"),
//   	},
//   	transformOutput: &transformOutputProperty{
//   		s3OutputPath: jsii.String("s3OutputPath"),
//
//   		// the properties below are optional
//   		accept: jsii.String("accept"),
//   		assembleWith: jsii.String("assembleWith"),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	transformResources: &transformResourcesProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//
//   	// the properties below are optional
//   	batchStrategy: jsii.String("batchStrategy"),
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	maxConcurrentTransforms: jsii.Number(123),
//   	maxPayloadInMb: jsii.Number(123),
//   }
//
type CfnModelPackage_TransformJobDefinitionProperty struct {
	// A description of the input source and the way the transform job consumes it.
	TransformInput interface{} `field:"required" json:"transformInput" yaml:"transformInput"`
	// Identifies the Amazon S3 location where you want Amazon SageMaker to save the results from the transform job.
	TransformOutput interface{} `field:"required" json:"transformOutput" yaml:"transformOutput"`
	// Identifies the ML compute instances for the transform job.
	TransformResources interface{} `field:"required" json:"transformResources" yaml:"transformResources"`
	// A string that determines the number of records included in a single mini-batch.
	//
	// `SingleRecord` means only one record is used per mini-batch. `MultiRecord` means a mini-batch is set to contain as many records that can fit within the `MaxPayloadInMB` limit.
	BatchStrategy *string `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// The environment variables to set in the Docker container.
	//
	// We support up to 16 key and values entries in the map.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The maximum number of parallel requests that can be sent to each instance in a transform job.
	//
	// The default value is 1.
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// The maximum payload size allowed, in MB.
	//
	// A payload is the data portion of a record (without metadata).
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
}

