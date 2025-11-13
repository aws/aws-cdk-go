package interfacesawsmediatailor


// A reference to a SourceLocation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceLocationReference := &SourceLocationReference{
//   	SourceLocationArn: jsii.String("sourceLocationArn"),
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   }
//
type SourceLocationReference struct {
	// The ARN of the SourceLocation resource.
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// The SourceLocationName of the SourceLocation resource.
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
}

