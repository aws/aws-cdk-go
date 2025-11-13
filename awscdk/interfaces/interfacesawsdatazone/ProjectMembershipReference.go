package interfacesawsdatazone


// A reference to a ProjectMembership resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectMembershipReference := &ProjectMembershipReference{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	MemberIdentifier: jsii.String("memberIdentifier"),
//   	MemberIdentifierType: jsii.String("memberIdentifierType"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   }
//
type ProjectMembershipReference struct {
	// The DomainIdentifier of the ProjectMembership resource.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The MemberIdentifier of the ProjectMembership resource.
	MemberIdentifier *string `field:"required" json:"memberIdentifier" yaml:"memberIdentifier"`
	// The MemberIdentifierType of the ProjectMembership resource.
	MemberIdentifierType *string `field:"required" json:"memberIdentifierType" yaml:"memberIdentifierType"`
	// The ProjectIdentifier of the ProjectMembership resource.
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
}

