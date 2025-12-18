package awscleanrooms


// Payment configuration for synthetic data generation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syntheticDataGenerationPaymentConfigProperty := &SyntheticDataGenerationPaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-syntheticdatagenerationpaymentconfig.html
//
type CfnCollaboration_SyntheticDataGenerationPaymentConfigProperty struct {
	// Indicates who is responsible for paying for synthetic data generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-syntheticdatagenerationpaymentconfig.html#cfn-cleanrooms-collaboration-syntheticdatagenerationpaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"required" json:"isResponsible" yaml:"isResponsible"`
}

