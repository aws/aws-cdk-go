package awssagemaker


// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointInputProperty := &endpointInputProperty{
//   	endpointName: jsii.String("endpointName"),
//   	localPath: jsii.String("localPath"),
//
//   	// the properties below are optional
//   	s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   	s3InputMode: jsii.String("s3InputMode"),
//   }
//
type CfnMonitoringSchedule_EndpointInputProperty struct {
	// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Path to the filesystem where the endpoint data is available to the container.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key.
	//
	// Defaults to `FullyReplicated`.
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the `Pipe` or `File` is used as the input mode for transferring data for the monitoring job.
	//
	// `Pipe` mode is recommended for large datasets. `File` mode is useful for small files that fit in memory. Defaults to `File` .
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

