package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipPaymentConfigurationProperty := &MembershipPaymentConfigurationProperty{
//   	QueryCompute: &MembershipQueryComputePaymentConfigProperty{
//   		IsResponsible: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html
//
type CfnMembership_MembershipPaymentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html#cfn-cleanrooms-membership-membershippaymentconfiguration-querycompute
	//
	QueryCompute interface{} `field:"required" json:"queryCompute" yaml:"queryCompute"`
}

