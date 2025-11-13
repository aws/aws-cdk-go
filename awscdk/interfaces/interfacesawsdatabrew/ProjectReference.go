package interfacesawsdatabrew


// A reference to a Project resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectReference := &ProjectReference{
//   	ProjectName: jsii.String("projectName"),
//   }
//
type ProjectReference struct {
	// The Name of the Project resource.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
}

