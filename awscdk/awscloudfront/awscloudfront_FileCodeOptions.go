package awscloudfront


// Options when reading the function's code from an external file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileCodeOptions := &fileCodeOptions{
//   	filePath: jsii.String("filePath"),
//   }
//
type FileCodeOptions struct {
	// The path of the file to read the code from.
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
}

