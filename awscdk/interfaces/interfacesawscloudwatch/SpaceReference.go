package interfacesawscloudwatch


// A reference to a Space resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceReference := &SpaceReference{
//   	SpaceArn: jsii.String("spaceArn"),
//   }
//
type SpaceReference struct {
	// The Arn of the Space resource.
	SpaceArn *string `field:"required" json:"spaceArn" yaml:"spaceArn"`
}

