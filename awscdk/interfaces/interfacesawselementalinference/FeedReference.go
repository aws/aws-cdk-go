package interfacesawselementalinference


// A reference to a Feed resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   feedReference := &FeedReference{
//   	FeedArn: jsii.String("feedArn"),
//   	FeedId: jsii.String("feedId"),
//   }
//
type FeedReference struct {
	// The ARN of the Feed resource.
	FeedArn *string `field:"required" json:"feedArn" yaml:"feedArn"`
	// The Id of the Feed resource.
	FeedId *string `field:"required" json:"feedId" yaml:"feedId"`
}

