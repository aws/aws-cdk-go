package previewawscleanroomsmixins


// Configuration for payment for synthetic data generation in a membership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipSyntheticDataGenerationPaymentConfigProperty := &MembershipSyntheticDataGenerationPaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipsyntheticdatagenerationpaymentconfig.html
//
type CfnMembershipPropsMixin_MembershipSyntheticDataGenerationPaymentConfigProperty struct {
	// Indicates if this membership is responsible for paying for synthetic data generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipsyntheticdatagenerationpaymentconfig.html#cfn-cleanrooms-membership-membershipsyntheticdatagenerationpaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"optional" json:"isResponsible" yaml:"isResponsible"`
}

