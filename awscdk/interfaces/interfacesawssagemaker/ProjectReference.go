package interfacesawssagemaker


// A reference to a Project resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectReference := &ProjectReference{
//   	ProjectArn: jsii.String("projectArn"),
//   }
//
type ProjectReference struct {
	// The ProjectArn of the Project resource.
	ProjectArn *string `field:"required" json:"projectArn" yaml:"projectArn"`
}

