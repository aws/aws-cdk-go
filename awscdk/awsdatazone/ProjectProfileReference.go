package awsdatazone


// A reference to a ProjectProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectProfileReference := &ProjectProfileReference{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Identifier: jsii.String("identifier"),
//   }
//
type ProjectProfileReference struct {
	// The DomainIdentifier of the ProjectProfile resource.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The Identifier of the ProjectProfile resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

