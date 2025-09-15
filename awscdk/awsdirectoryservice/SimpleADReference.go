package awsdirectoryservice


// A reference to a SimpleAD resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleADReference := &SimpleADReference{
//   	DirectoryId: jsii.String("directoryId"),
//   }
//
type SimpleADReference struct {
	// The DirectoryId of the SimpleAD resource.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
}

