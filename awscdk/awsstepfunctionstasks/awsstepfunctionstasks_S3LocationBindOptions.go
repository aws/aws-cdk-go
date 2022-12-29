package awsstepfunctionstasks


// Options for binding an S3 Location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationBindOptions := &s3LocationBindOptions{
//   	forReading: jsii.Boolean(false),
//   	forWriting: jsii.Boolean(false),
//   }
//
type S3LocationBindOptions struct {
	// Allow reading from the S3 Location.
	ForReading *bool `field:"optional" json:"forReading" yaml:"forReading"`
	// Allow writing to the S3 Location.
	ForWriting *bool `field:"optional" json:"forWriting" yaml:"forWriting"`
}

