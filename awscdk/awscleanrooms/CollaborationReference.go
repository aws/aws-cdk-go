package awscleanrooms


// A reference to a Collaboration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collaborationReference := &CollaborationReference{
//   	CollaborationArn: jsii.String("collaborationArn"),
//   	CollaborationIdentifier: jsii.String("collaborationIdentifier"),
//   }
//
type CollaborationReference struct {
	// The ARN of the Collaboration resource.
	CollaborationArn *string `field:"required" json:"collaborationArn" yaml:"collaborationArn"`
	// The CollaborationIdentifier of the Collaboration resource.
	CollaborationIdentifier *string `field:"required" json:"collaborationIdentifier" yaml:"collaborationIdentifier"`
}

