package awsarcregionswitch


// A reference to a Plan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   planReference := &PlanReference{
//   	PlanArn: jsii.String("planArn"),
//   }
//
type PlanReference struct {
	// The Arn of the Plan resource.
	PlanArn *string `field:"required" json:"planArn" yaml:"planArn"`
}

