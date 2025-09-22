package awsbatch


// A reference to a SchedulingPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schedulingPolicyReference := &SchedulingPolicyReference{
//   	SchedulingPolicyArn: jsii.String("schedulingPolicyArn"),
//   }
//
type SchedulingPolicyReference struct {
	// The Arn of the SchedulingPolicy resource.
	SchedulingPolicyArn *string `field:"required" json:"schedulingPolicyArn" yaml:"schedulingPolicyArn"`
}

