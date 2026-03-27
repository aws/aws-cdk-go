package awsbatch


// Quota Share Policy for the Job Queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quotaSharePolicyProperty := &QuotaSharePolicyProperty{
//   	IdleResourceAssignmentStrategy: jsii.String("idleResourceAssignmentStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-quotasharepolicy.html
//
type CfnSchedulingPolicy_QuotaSharePolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-schedulingpolicy-quotasharepolicy.html#cfn-batch-schedulingpolicy-quotasharepolicy-idleresourceassignmentstrategy
	//
	IdleResourceAssignmentStrategy *string `field:"optional" json:"idleResourceAssignmentStrategy" yaml:"idleResourceAssignmentStrategy"`
}

