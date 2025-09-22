package awsomics


// A reference to a RunGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runGroupReference := &RunGroupReference{
//   	RunGroupArn: jsii.String("runGroupArn"),
//   	RunGroupId: jsii.String("runGroupId"),
//   }
//
type RunGroupReference struct {
	// The ARN of the RunGroup resource.
	RunGroupArn *string `field:"required" json:"runGroupArn" yaml:"runGroupArn"`
	// The Id of the RunGroup resource.
	RunGroupId *string `field:"required" json:"runGroupId" yaml:"runGroupId"`
}

