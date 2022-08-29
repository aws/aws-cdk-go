package awscodebuild


// Properties common to all Artifacts classes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactsProps := &artifactsProps{
//   	identifier: jsii.String("identifier"),
//   }
//
type ArtifactsProps struct {
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

