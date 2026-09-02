package interfacesawsquicksight


// A reference to a ApprovalPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   approvalPolicyReference := &ApprovalPolicyReference{
//   	PolicyArn: jsii.String("policyArn"),
//   	PolicyId: jsii.String("policyId"),
//   }
//
type ApprovalPolicyReference struct {
	// The ARN of the ApprovalPolicy resource.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// The PolicyId of the ApprovalPolicy resource.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

