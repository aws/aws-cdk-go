package interfacesawssagemaker


// A reference to a Artifact resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactReference := &ArtifactReference{
//   	ArtifactArn: jsii.String("artifactArn"),
//   }
//
type ArtifactReference struct {
	// The Arn of the Artifact resource.
	ArtifactArn *string `field:"required" json:"artifactArn" yaml:"artifactArn"`
}

