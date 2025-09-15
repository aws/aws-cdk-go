package awslogs


// A reference to a AccountPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountPolicyReference := &AccountPolicyReference{
//   	AccountId: jsii.String("accountId"),
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//   }
//
type AccountPolicyReference struct {
	// The AccountId of the AccountPolicy resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The PolicyName of the AccountPolicy resource.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The PolicyType of the AccountPolicy resource.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
}

