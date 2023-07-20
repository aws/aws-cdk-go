package awssagemaker


// The metadata of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointMetadataProperty := &EndpointMetadataProperty{
//   	EndpointName: jsii.String("endpointName"),
//
//   	// the properties below are optional
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//   	EndpointStatus: jsii.String("endpointStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html
//
type CfnInferenceExperiment_EndpointMetadataProperty struct {
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The name of the endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointconfigname
	//
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The status of the endpoint.
	//
	// For possible values of the status of an endpoint, see [](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointstatus) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointstatus
	//
	EndpointStatus *string `field:"optional" json:"endpointStatus" yaml:"endpointStatus"`
}

