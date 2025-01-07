package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mLPaymentConfigProperty := &MLPaymentConfigProperty{
//   	ModelInference: &ModelInferencePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   	ModelTraining: &ModelTrainingPaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html
//
type CfnCollaboration_MLPaymentConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html#cfn-cleanrooms-collaboration-mlpaymentconfig-modelinference
	//
	ModelInference interface{} `field:"optional" json:"modelInference" yaml:"modelInference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html#cfn-cleanrooms-collaboration-mlpaymentconfig-modeltraining
	//
	ModelTraining interface{} `field:"optional" json:"modelTraining" yaml:"modelTraining"`
}

