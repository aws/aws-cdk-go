package previewawscleanroomsmixins


// An object representing the payment responsibilities accepted by the collaboration member for query compute costs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipQueryComputePaymentConfigProperty := &MembershipQueryComputePaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipquerycomputepaymentconfig.html
//
type CfnMembershipPropsMixin_MembershipQueryComputePaymentConfigProperty struct {
	// Indicates whether the collaboration member has accepted to pay for query compute costs ( `TRUE` ) or has not accepted to pay for query compute costs ( `FALSE` ).
	//
	// If the collaboration creator has not specified anyone to pay for query compute costs, then the member who can query is the default payer.
	//
	// An error message is returned for the following reasons:
	//
	// - If you set the value to `FALSE` but you are responsible to pay for query compute costs.
	// - If you set the value to `TRUE` but you are not responsible to pay for query compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipquerycomputepaymentconfig.html#cfn-cleanrooms-membership-membershipquerycomputepaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"optional" json:"isResponsible" yaml:"isResponsible"`
}

