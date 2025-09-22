package awselasticloadbalancingv2


// A reference to a Listener resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerReference := &ListenerReference{
//   	ListenerArn: jsii.String("listenerArn"),
//   }
//
type ListenerReference struct {
	// The ListenerArn of the Listener resource.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
}

