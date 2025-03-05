package awssagemaker


// Container image configuration object for the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelQualityAppSpecificationProperty := &ModelQualityAppSpecificationProperty{
//   	ImageUri: jsii.String("imageUri"),
//   	ProblemType: jsii.String("problemType"),
//
//   	// the properties below are optional
//   	ContainerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	ContainerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   	RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html
//
type CfnModelQualityJobDefinition_ModelQualityAppSpecificationProperty struct {
	// The address of the container image that the monitoring job runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The machine learning problem type of the model that the monitoring job monitors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-problemtype
	//
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// An array of arguments for the container used to run the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-containerarguments
	//
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Specifies the entrypoint for a container that the monitoring job runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-containerentrypoint
	//
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Sets the environment variables in the container that the monitoring job runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-postanalyticsprocessorsourceuri
	//
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flattened JSON so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-modelqualityappspecification.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification-recordpreprocessorsourceuri
	//
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

