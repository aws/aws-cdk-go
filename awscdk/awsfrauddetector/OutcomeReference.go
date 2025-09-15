package awsfrauddetector


// A reference to a Outcome resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outcomeReference := &OutcomeReference{
//   	OutcomeArn: jsii.String("outcomeArn"),
//   }
//
type OutcomeReference struct {
	// The Arn of the Outcome resource.
	OutcomeArn *string `field:"required" json:"outcomeArn" yaml:"outcomeArn"`
}

