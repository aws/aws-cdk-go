package awsmediaconnect


// A reference to a FlowSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowSourceReference := &FlowSourceReference{
//   	SourceArn: jsii.String("sourceArn"),
//   }
//
type FlowSourceReference struct {
	// The SourceArn of the FlowSource resource.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
}

