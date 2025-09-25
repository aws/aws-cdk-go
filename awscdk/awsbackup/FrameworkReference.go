package awsbackup


// A reference to a Framework resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameworkReference := &FrameworkReference{
//   	FrameworkArn: jsii.String("frameworkArn"),
//   }
//
type FrameworkReference struct {
	// The FrameworkArn of the Framework resource.
	FrameworkArn *string `field:"required" json:"frameworkArn" yaml:"frameworkArn"`
}

