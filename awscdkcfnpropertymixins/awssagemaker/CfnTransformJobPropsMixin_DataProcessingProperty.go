package awssagemaker


// The data structure used to specify the data to be used for inference in a batch transform job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataProcessingProperty := &DataProcessingProperty{
//   	InputFilter: jsii.String("inputFilter"),
//   	JoinSource: jsii.String("joinSource"),
//   	OutputFilter: jsii.String("outputFilter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-dataprocessing.html
//
type CfnTransformJobPropsMixin_DataProcessingProperty struct {
	// A JSONPath expression used to select a portion of the input data to pass to the algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-dataprocessing.html#cfn-sagemaker-transformjob-dataprocessing-inputfilter
	//
	InputFilter *string `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// Specifies the source of the data to join with the transformed data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-dataprocessing.html#cfn-sagemaker-transformjob-dataprocessing-joinsource
	//
	JoinSource *string `field:"optional" json:"joinSource" yaml:"joinSource"`
	// A JSONPath expression used to select a portion of the joined dataset to save in the output file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-dataprocessing.html#cfn-sagemaker-transformjob-dataprocessing-outputfilter
	//
	OutputFilter *string `field:"optional" json:"outputFilter" yaml:"outputFilter"`
}

