package mixinsawscleanrooms


// An object representing the payment responsibilities accepted by the collaboration member.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipPaymentConfigurationProperty := &MembershipPaymentConfigurationProperty{
//   	JobCompute: &MembershipJobComputePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   	MachineLearning: &MembershipMLPaymentConfigProperty{
//   		ModelInference: &MembershipModelInferencePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   		ModelTraining: &MembershipModelTrainingPaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   	},
//   	QueryCompute: &MembershipQueryComputePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html
//
type CfnMembershipPropsMixin_MembershipPaymentConfigurationProperty struct {
	// The payment responsibilities accepted by the collaboration member for job compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html#cfn-cleanrooms-membership-membershippaymentconfiguration-jobcompute
	//
	JobCompute interface{} `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// The payment responsibilities accepted by the collaboration member for machine learning costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html#cfn-cleanrooms-membership-membershippaymentconfiguration-machinelearning
	//
	MachineLearning interface{} `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// The payment responsibilities accepted by the collaboration member for query compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html#cfn-cleanrooms-membership-membershippaymentconfiguration-querycompute
	//
	QueryCompute interface{} `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

