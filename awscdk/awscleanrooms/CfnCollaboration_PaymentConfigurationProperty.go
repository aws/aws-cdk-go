package awscleanrooms


// An object representing the collaboration member's payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paymentConfigurationProperty := &PaymentConfigurationProperty{
//   	QueryCompute: &QueryComputePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
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
//   		SyntheticDataGeneration: &SyntheticDataGenerationPaymentConfigProperty{
//   			IsResponsible: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html
//
type CfnCollaboration_PaymentConfigurationProperty struct {
	// The collaboration member's payment responsibilities set by the collaboration creator for query compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-querycompute
	//
	QueryCompute interface{} `field:"required" json:"queryCompute" yaml:"queryCompute"`
	// The compute configuration for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-jobcompute
	//
	JobCompute interface{} `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// An object representing the collaboration member's machine learning payment responsibilities set by the collaboration creator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-machinelearning
	//
	MachineLearning interface{} `field:"optional" json:"machineLearning" yaml:"machineLearning"`
}

