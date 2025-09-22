package awsappstream


// A reference to a DirectoryConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directoryConfigReference := &DirectoryConfigReference{
//   	DirectoryName: jsii.String("directoryName"),
//   }
//
type DirectoryConfigReference struct {
	// The DirectoryName of the DirectoryConfig resource.
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
}

