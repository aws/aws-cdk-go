package interfacesawslambda


// A reference to a WebFunctionRevision resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionRevisionReference := &WebFunctionRevisionReference{
//   	RevisionArn: jsii.String("revisionArn"),
//   }
//
type WebFunctionRevisionReference struct {
	// The RevisionArn of the WebFunctionRevision resource.
	RevisionArn *string `field:"required" json:"revisionArn" yaml:"revisionArn"`
}

