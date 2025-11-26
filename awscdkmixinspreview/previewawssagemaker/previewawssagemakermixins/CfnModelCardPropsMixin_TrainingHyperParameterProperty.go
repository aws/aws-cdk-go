package previewawssagemakermixins


// A hyper parameter that was configured in training the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trainingHyperParameterProperty := &TrainingHyperParameterProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-traininghyperparameter.html
//
type CfnModelCardPropsMixin_TrainingHyperParameterProperty struct {
	// The name of the hyper parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-traininghyperparameter.html#cfn-sagemaker-modelcard-traininghyperparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value specified for the hyper parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-traininghyperparameter.html#cfn-sagemaker-modelcard-traininghyperparameter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

