package awssagemaker


// The inputs for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringInputProperty := &MonitoringInputProperty{
//   	BatchTransformInput: &BatchTransformInputProperty{
//   		DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   		DatasetFormat: &DatasetFormatProperty{
//   			Csv: &CsvProperty{
//   				Header: jsii.Boolean(false),
//   			},
//   			Json: &JsonProperty{
//   				Line: jsii.Boolean(false),
//   			},
//   			Parquet: jsii.Boolean(false),
//   		},
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   	EndpointInput: &EndpointInputProperty{
//   		EndpointName: jsii.String("endpointName"),
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringInputProperty struct {
	// `CfnMonitoringSchedule.MonitoringInputProperty.BatchTransformInput`.
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// The endpoint for a monitoring job.
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

