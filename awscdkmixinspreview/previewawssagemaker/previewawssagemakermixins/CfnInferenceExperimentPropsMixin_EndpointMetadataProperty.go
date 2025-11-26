package previewawssagemakermixins


// The metadata of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointMetadataProperty := &EndpointMetadataProperty{
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//   	EndpointName: jsii.String("endpointName"),
//   	EndpointStatus: jsii.String("endpointStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html
//
type CfnInferenceExperimentPropsMixin_EndpointMetadataProperty struct {
	// The name of the endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointconfigname
	//
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The name of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The status of the endpoint.
	//
	// For possible values of the status of an endpoint, see [](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointstatus) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-endpointmetadata.html#cfn-sagemaker-inferenceexperiment-endpointmetadata-endpointstatus
	//
	EndpointStatus *string `field:"optional" json:"endpointStatus" yaml:"endpointStatus"`
}

