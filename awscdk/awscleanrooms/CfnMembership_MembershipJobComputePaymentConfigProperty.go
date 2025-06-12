package awscleanrooms


// An object representing the payment responsibilities accepted by the collaboration member for query and job compute costs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipJobComputePaymentConfigProperty := &MembershipJobComputePaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipjobcomputepaymentconfig.html
//
type CfnMembership_MembershipJobComputePaymentConfigProperty struct {
	// Indicates whether the collaboration member has accepted to pay for job compute costs ( `TRUE` ) or has not accepted to pay for query and job compute costs ( `FALSE` ).
	//
	// There is only one member who pays for queries and jobs.
	//
	// An error message is returned for the following reasons:
	//
	// - If you set the value to `FALSE` but you are responsible to pay for query and job compute costs.
	// - If you set the value to `TRUE` but you are not responsible to pay for query and job compute costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipjobcomputepaymentconfig.html#cfn-cleanrooms-membership-membershipjobcomputepaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"required" json:"isResponsible" yaml:"isResponsible"`
}

