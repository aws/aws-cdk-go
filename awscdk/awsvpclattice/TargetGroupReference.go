package awsvpclattice


// A reference to a TargetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupReference := &TargetGroupReference{
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type TargetGroupReference struct {
	// The Arn of the TargetGroup resource.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
}

