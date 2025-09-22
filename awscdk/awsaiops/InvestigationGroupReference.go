package awsaiops


// A reference to a InvestigationGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   investigationGroupReference := &InvestigationGroupReference{
//   	InvestigationGroupArn: jsii.String("investigationGroupArn"),
//   }
//
type InvestigationGroupReference struct {
	// The Arn of the InvestigationGroup resource.
	InvestigationGroupArn *string `field:"required" json:"investigationGroupArn" yaml:"investigationGroupArn"`
}

