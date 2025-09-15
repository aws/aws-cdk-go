package pipelines


// Location of a FileSet consumed or produced by a ShellStep.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSet fileSet
//
//   fileSetLocation := &FileSetLocation{
//   	Directory: jsii.String("directory"),
//   	FileSet: fileSet,
//   }
//
type FileSetLocation struct {
	// The (relative) directory where the FileSet is found.
	Directory *string `field:"required" json:"directory" yaml:"directory"`
	// The FileSet object.
	FileSet FileSet `field:"required" json:"fileSet" yaml:"fileSet"`
}

