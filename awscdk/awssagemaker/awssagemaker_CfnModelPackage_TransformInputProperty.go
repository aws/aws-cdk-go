package awssagemaker


// Describes the input source of a transform job and the way the transform job consumes it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformInputProperty := &transformInputProperty{
//   	dataSource: &dataSourceProperty{
//   		s3DataSource: &s3DataSourceProperty{
//   			s3DataType: jsii.String("s3DataType"),
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	compressionType: jsii.String("compressionType"),
//   	contentType: jsii.String("contentType"),
//   	splitType: jsii.String("splitType"),
//   }
//
type CfnModelPackage_TransformInputProperty struct {
	// Describes the location of the channel data, which is, the S3 location of the input data that the model can consume.
	DataSource interface{} `field:"required" json:"dataSource" yaml:"dataSource"`
	// If your transform data is compressed, specify the compression type.
	//
	// Amazon SageMaker automatically decompresses the data for the transform job accordingly. The default value is `None` .
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The multipurpose internet mail extension (MIME) type of the data.
	//
	// Amazon SageMaker uses the MIME type with each http call to transfer data to the transform job.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The method to use to split the transform job's data files into smaller batches.
	//
	// Splitting is necessary when the total size of each object is too large to fit in a single request. You can also use data splitting to improve performance by processing multiple concurrent mini-batches. The default value for `SplitType` is `None` , which indicates that input data files are not split, and request payloads contain the entire contents of an input object. Set the value of this parameter to `Line` to split records on a newline character boundary. `SplitType` also supports a number of record-oriented binary data formats. Currently, the supported record formats are:
	//
	// - RecordIO
	// - TFRecord
	//
	// When splitting is enabled, the size of a mini-batch depends on the values of the `BatchStrategy` and `MaxPayloadInMB` parameters. When the value of `BatchStrategy` is `MultiRecord` , Amazon SageMaker sends the maximum number of records in each request, up to the `MaxPayloadInMB` limit. If the value of `BatchStrategy` is `SingleRecord` , Amazon SageMaker sends individual records in each request.
	//
	// > Some data formats represent a record as a binary payload wrapped with extra padding bytes. When splitting is applied to a binary data format, padding is removed if the value of `BatchStrategy` is set to `SingleRecord` . Padding is not removed if the value of `BatchStrategy` is set to `MultiRecord` .
	// >
	// > For more information about `RecordIO` , see [Create a Dataset Using RecordIO](https://docs.aws.amazon.com/https://mxnet.apache.org/api/faq/recordio) in the MXNet documentation. For more information about `TFRecord` , see [Consuming TFRecord data](https://docs.aws.amazon.com/https://www.tensorflow.org/guide/data#consuming_tfrecord_data) in the TensorFlow documentation.
	SplitType *string `field:"optional" json:"splitType" yaml:"splitType"`
}

