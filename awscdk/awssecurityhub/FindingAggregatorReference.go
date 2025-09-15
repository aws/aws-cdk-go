package awssecurityhub


// A reference to a FindingAggregator resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findingAggregatorReference := &FindingAggregatorReference{
//   	FindingAggregatorArn: jsii.String("findingAggregatorArn"),
//   }
//
type FindingAggregatorReference struct {
	// The FindingAggregatorArn of the FindingAggregator resource.
	FindingAggregatorArn *string `field:"required" json:"findingAggregatorArn" yaml:"findingAggregatorArn"`
}

