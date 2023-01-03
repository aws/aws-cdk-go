package awsstepfunctionstasks


// Stores information about the location of an object in Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationConfig := &s3LocationConfig{
//   	uri: jsii.String("uri"),
//   }
//
type S3LocationConfig struct {
	// Uniquely identifies the resource in Amazon S3.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

