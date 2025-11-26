package previewawspipesmixins


// Name/Value pair of a parameter to start execution of a SageMaker AI Model Building Pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerPipelineParameterProperty := &SageMakerPipelineParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-sagemakerpipelineparameter.html
//
type CfnPipePropsMixin_SageMakerPipelineParameterProperty struct {
	// Name of parameter to start execution of a SageMaker AI Model Building Pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-sagemakerpipelineparameter.html#cfn-pipes-pipe-sagemakerpipelineparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker AI Model Building Pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-sagemakerpipelineparameter.html#cfn-pipes-pipe-sagemakerpipelineparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

