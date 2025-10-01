package awsfrauddetector


// A reference to a EntityType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityTypeReference := &EntityTypeReference{
//   	EntityTypeArn: jsii.String("entityTypeArn"),
//   }
//
type EntityTypeReference struct {
	// The Arn of the EntityType resource.
	EntityTypeArn *string `field:"required" json:"entityTypeArn" yaml:"entityTypeArn"`
}

