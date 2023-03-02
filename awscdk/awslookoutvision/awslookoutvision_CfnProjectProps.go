package awslookoutvision


// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
//   	projectName: jsii.String("projectName"),
//   }
//
type CfnProjectProps struct {
	// The name of the project.
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
}

