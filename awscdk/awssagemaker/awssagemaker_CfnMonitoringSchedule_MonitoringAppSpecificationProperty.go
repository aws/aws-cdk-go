package awssagemaker


// Container image configuration object for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringAppSpecificationProperty := &monitoringAppSpecificationProperty{
//   	imageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	containerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	containerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   	recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   }
//
type CfnMonitoringSchedule_MonitoringAppSpecificationProperty struct {
	// The container image to be run by the monitoring job.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// An array of arguments for the container used to run the monitoring job.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Specifies the entrypoint for a container used to run the monitoring job.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

