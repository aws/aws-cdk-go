package interfacesawscloudwatch


// A reference to a DomainAccessGrantForOrganization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainAccessGrantForOrganizationReference := &DomainAccessGrantForOrganizationReference{
//   	DomainAccessGrantForOrganizationArn: jsii.String("domainAccessGrantForOrganizationArn"),
//   }
//
type DomainAccessGrantForOrganizationReference struct {
	// The Arn of the DomainAccessGrantForOrganization resource.
	DomainAccessGrantForOrganizationArn *string `field:"required" json:"domainAccessGrantForOrganizationArn" yaml:"domainAccessGrantForOrganizationArn"`
}

