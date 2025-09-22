package awsdatazone


// A reference to a Project resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectReference := &ProjectReference{
//   	DomainId: jsii.String("domainId"),
//   	ProjectId: jsii.String("projectId"),
//   }
//
type ProjectReference struct {
	// The DomainId of the Project resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The Id of the Project resource.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

