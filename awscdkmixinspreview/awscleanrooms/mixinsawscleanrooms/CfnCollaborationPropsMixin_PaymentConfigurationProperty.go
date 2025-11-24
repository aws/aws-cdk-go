package mixinsawscleanrooms


// An object representing the collaboration member's payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   paymentConfigurationProperty := &PaymentConfigurationProperty{
//   	JobCompute: &JobComputePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   	MachineLearning: &MLPaymentConfigProperty{
//   		ModelInference: &ModelInferencePaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   		ModelTraining: &ModelTrainingPaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   	},
//   	QueryCompute: &QueryComputePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html
//
type CfnCollaborationPropsMixin_PaymentConfigurationProperty struct {
	// The compute configuration for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-jobcompute
	//
	JobCompute interface{} `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// An object representing the collaboration member's machine learning payment responsibilities set by the collaboration creator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-machinelearning
	//
	MachineLearning interface{} `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// The collaboration member's payment responsibilities set by the collaboration creator for query compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-querycompute
	//
	QueryCompute interface{} `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

