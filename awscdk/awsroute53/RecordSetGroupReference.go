package awsroute53


// A reference to a RecordSetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordSetGroupReference := &RecordSetGroupReference{
//   	RecordSetGroupId: jsii.String("recordSetGroupId"),
//   }
//
type RecordSetGroupReference struct {
	// The Id of the RecordSetGroup resource.
	RecordSetGroupId *string `field:"required" json:"recordSetGroupId" yaml:"recordSetGroupId"`
}

