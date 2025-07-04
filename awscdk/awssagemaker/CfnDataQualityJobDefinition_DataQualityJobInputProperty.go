package awssagemaker


// The input for the data quality monitoring job.
//
// Currently endpoints are supported for input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityJobInputProperty := &DataQualityJobInputProperty{
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
//   		ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   	EndpointInput: &EndpointInputProperty{
//   		EndpointName: jsii.String("endpointName"),
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualityjobinput.html
//
type CfnDataQualityJobDefinition_DataQualityJobInputProperty struct {
	// Input object for the batch transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualityjobinput.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityjobinput-batchtransforminput
	//
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// Input object for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-dataqualityjobinput.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityjobinput-endpointinput
	//
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

