package awsses


// The reputation policy for the tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reputationPolicyProperty := &ReputationPolicyProperty{
//   	ReputationEntityPolicy: jsii.String("reputationEntityPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-reputationpolicy.html
//
type CfnTenant_ReputationPolicyProperty struct {
	// The ARN of an SES-managed reputation policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-reputationpolicy.html#cfn-ses-tenant-reputationpolicy-reputationentitypolicy
	//
	ReputationEntityPolicy *string `field:"required" json:"reputationEntityPolicy" yaml:"reputationEntityPolicy"`
}

