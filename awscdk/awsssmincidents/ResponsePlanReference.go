package awsssmincidents


// A reference to a ResponsePlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responsePlanReference := &ResponsePlanReference{
//   	ResponsePlanArn: jsii.String("responsePlanArn"),
//   }
//
type ResponsePlanReference struct {
	// The Arn of the ResponsePlan resource.
	ResponsePlanArn *string `field:"required" json:"responsePlanArn" yaml:"responsePlanArn"`
}

