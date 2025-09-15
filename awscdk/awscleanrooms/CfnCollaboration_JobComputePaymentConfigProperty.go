package awscleanrooms


// An object representing the collaboration member's payment responsibilities set by the collaboration creator for query and job compute costs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobComputePaymentConfigProperty := &JobComputePaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-jobcomputepaymentconfig.html
//
type CfnCollaboration_JobComputePaymentConfigProperty struct {
	// Indicates whether the collaboration creator has configured the collaboration member to pay for query and job compute costs ( `TRUE` ) or has not configured the collaboration member to pay for query and job compute costs ( `FALSE` ).
	//
	// Exactly one member can be configured to pay for query and job compute costs. An error is returned if the collaboration creator sets a `TRUE` value for more than one member in the collaboration.
	//
	// An error is returned if the collaboration creator sets a `FALSE` value for the member who can run queries and jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-jobcomputepaymentconfig.html#cfn-cleanrooms-collaboration-jobcomputepaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"required" json:"isResponsible" yaml:"isResponsible"`
}

