package awsstepfunctionstasks


// Configuration for a using Docker image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageConfig := &dockerImageConfig{
//   	imageUri: jsii.String("imageUri"),
//   }
//
type DockerImageConfig struct {
	// The fully qualified URI of the Docker image.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

