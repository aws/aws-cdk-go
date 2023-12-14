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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html
//
type CfnCollaboration_PaymentConfigurationProperty struct {
	// The collaboration member's payment responsibilities set by the collaboration creator for query compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-paymentconfiguration.html#cfn-cleanrooms-collaboration-paymentconfiguration-querycompute
	//
	QueryCompute interface{} `field:"required" json:"queryCompute" yaml:"queryCompute"`
}

