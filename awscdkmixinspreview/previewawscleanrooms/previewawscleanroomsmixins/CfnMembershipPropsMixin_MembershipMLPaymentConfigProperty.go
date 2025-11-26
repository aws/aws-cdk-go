package previewawscleanroomsmixins


// An object representing the collaboration member's machine learning payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipMLPaymentConfigProperty := &MembershipMLPaymentConfigProperty{
//   	ModelInference: &MembershipModelInferencePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   	ModelTraining: &MembershipModelTrainingPaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmlpaymentconfig.html
//
type CfnMembershipPropsMixin_MembershipMLPaymentConfigProperty struct {
	// The payment responsibilities accepted by the member for model inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmlpaymentconfig.html#cfn-cleanrooms-membership-membershipmlpaymentconfig-modelinference
	//
	ModelInference interface{} `field:"optional" json:"modelInference" yaml:"modelInference"`
	// The payment responsibilities accepted by the member for model training.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmlpaymentconfig.html#cfn-cleanrooms-membership-membershipmlpaymentconfig-modeltraining
	//
	ModelTraining interface{} `field:"optional" json:"modelTraining" yaml:"modelTraining"`
}

