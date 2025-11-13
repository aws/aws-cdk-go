package interfacesawsfrauddetector


// A reference to a List resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listReference := &ListReference{
//   	ListArn: jsii.String("listArn"),
//   }
//
type ListReference struct {
	// The Arn of the List resource.
	ListArn *string `field:"required" json:"listArn" yaml:"listArn"`
}

