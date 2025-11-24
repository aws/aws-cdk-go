package mixinsawscleanrooms


// An object representing the collaboration member's model inference payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipModelInferencePaymentConfigProperty := &MembershipModelInferencePaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmodelinferencepaymentconfig.html
//
type CfnMembershipPropsMixin_MembershipModelInferencePaymentConfigProperty struct {
	// Indicates whether the collaboration member has accepted to pay for model inference costs ( `TRUE` ) or has not accepted to pay for model inference costs ( `FALSE` ).
	//
	// If the collaboration creator has not specified anyone to pay for model inference costs, then the member who can query is the default payer.
	//
	// An error message is returned for the following reasons:
	//
	// - If you set the value to `FALSE` but you are responsible to pay for model inference costs.
	// - If you set the value to `TRUE` but you are not responsible to pay for model inference costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipmodelinferencepaymentconfig.html#cfn-cleanrooms-membership-membershipmodelinferencepaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"optional" json:"isResponsible" yaml:"isResponsible"`
}

