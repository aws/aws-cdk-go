package interfacesawsfsx


// A reference to a FileCache resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileCacheReference := &FileCacheReference{
//   	FileCacheId: jsii.String("fileCacheId"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type FileCacheReference struct {
	// The FileCacheId of the FileCache resource.
	FileCacheId *string `field:"required" json:"fileCacheId" yaml:"fileCacheId"`
	// The ResourceARN of the FileCache resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

