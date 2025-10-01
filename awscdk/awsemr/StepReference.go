package awsemr


// A reference to a Step resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepReference := &StepReference{
//   	StepId: jsii.String("stepId"),
//   }
//
type StepReference struct {
	// The Id of the Step resource.
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
}

