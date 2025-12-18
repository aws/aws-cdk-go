package previewawscleanroomsmixins


// An object representing the collaboration member's machine learning payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mLPaymentConfigProperty := &MLPaymentConfigProperty{
//   	ModelInference: &ModelInferencePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   	ModelTraining: &ModelTrainingPaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   	SyntheticDataGeneration: &SyntheticDataGenerationPaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html
//
type CfnCollaborationPropsMixin_MLPaymentConfigProperty struct {
	// The payment responsibilities accepted by the member for model inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html#cfn-cleanrooms-collaboration-mlpaymentconfig-modelinference
	//
	ModelInference interface{} `field:"optional" json:"modelInference" yaml:"modelInference"`
	// The payment responsibilities accepted by the member for model training.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html#cfn-cleanrooms-collaboration-mlpaymentconfig-modeltraining
	//
	ModelTraining interface{} `field:"optional" json:"modelTraining" yaml:"modelTraining"`
	// The payment configuration for machine learning synthetic data generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html#cfn-cleanrooms-collaboration-mlpaymentconfig-syntheticdatageneration
	//
	SyntheticDataGeneration interface{} `field:"optional" json:"syntheticDataGeneration" yaml:"syntheticDataGeneration"`
}

