package awscloudformation


// A reference to a PublicTypeVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicTypeVersionReference := &PublicTypeVersionReference{
//   	PublicTypeArn: jsii.String("publicTypeArn"),
//   }
//
type PublicTypeVersionReference struct {
	// The PublicTypeArn of the PublicTypeVersion resource.
	PublicTypeArn *string `field:"required" json:"publicTypeArn" yaml:"publicTypeArn"`
}

