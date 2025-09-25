package awsroute53


// A reference to a RecordSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordSetReference := &RecordSetReference{
//   	RecordSetId: jsii.String("recordSetId"),
//   }
//
type RecordSetReference struct {
	// The Id of the RecordSet resource.
	RecordSetId *string `field:"required" json:"recordSetId" yaml:"recordSetId"`
}

