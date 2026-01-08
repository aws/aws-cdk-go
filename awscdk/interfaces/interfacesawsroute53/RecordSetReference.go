package interfacesawsroute53


// A reference to a RecordSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordSetReference := &RecordSetReference{
//   	RecordSetName: jsii.String("recordSetName"),
//   }
//
type RecordSetReference struct {
	// The Name of the RecordSet resource.
	RecordSetName *string `field:"required" json:"recordSetName" yaml:"recordSetName"`
}

