package awscloudformation


// A reference to a Publisher resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publisherReference := &PublisherReference{
//   	PublisherId: jsii.String("publisherId"),
//   }
//
type PublisherReference struct {
	// The PublisherId of the Publisher resource.
	PublisherId *string `field:"required" json:"publisherId" yaml:"publisherId"`
}

