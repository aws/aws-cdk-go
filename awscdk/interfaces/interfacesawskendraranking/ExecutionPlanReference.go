package interfacesawskendraranking


// A reference to a ExecutionPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionPlanReference := &ExecutionPlanReference{
//   	ExecutionPlanArn: jsii.String("executionPlanArn"),
//   	ExecutionPlanId: jsii.String("executionPlanId"),
//   }
//
type ExecutionPlanReference struct {
	// The ARN of the ExecutionPlan resource.
	ExecutionPlanArn *string `field:"required" json:"executionPlanArn" yaml:"executionPlanArn"`
	// The Id of the ExecutionPlan resource.
	ExecutionPlanId *string `field:"required" json:"executionPlanId" yaml:"executionPlanId"`
}

