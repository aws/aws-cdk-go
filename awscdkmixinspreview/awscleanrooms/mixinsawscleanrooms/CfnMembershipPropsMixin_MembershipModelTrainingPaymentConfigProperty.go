package mixinsawscleanrooms


// An object representing the collaboration member's model training payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipModelTrainingPaymentConfigProperty := &MembershipModelTrainingPaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmodeltrainingpaymentconfig.html
//
type CfnMembershipPropsMixin_MembershipModelTrainingPaymentConfigProperty struct {
	// Indicates whether the collaboration member has accepted to pay for model training costs ( `TRUE` ) or has not accepted to pay for model training costs ( `FALSE` ).
	//
	// If the collaboration creator has not specified anyone to pay for model training costs, then the member who can query is the default payer.
	//
	// An error message is returned for the following reasons:
	//
	// - If you set the value to `FALSE` but you are responsible to pay for model training costs.
	// - If you set the value to `TRUE` but you are not responsible to pay for model training costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmodeltrainingpaymentconfig.html#cfn-cleanrooms-membership-membershipmodeltrainingpaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"optional" json:"isResponsible" yaml:"isResponsible"`
}

