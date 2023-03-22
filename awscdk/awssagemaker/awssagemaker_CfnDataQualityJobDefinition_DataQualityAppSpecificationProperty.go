package awssagemaker


// Information about the container that a data quality monitoring job runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityAppSpecificationProperty := &dataQualityAppSpecificationProperty{
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	containerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	containerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   	recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   }
//
type CfnDataQualityJobDefinition_DataQualityAppSpecificationProperty struct {
	// The container image that the data quality monitoring job runs.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The arguments to send to the container that the monitoring job runs.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// The entrypoint for a container used to run a monitoring job.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Sets the environment variables in the container that the monitoring job runs.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

