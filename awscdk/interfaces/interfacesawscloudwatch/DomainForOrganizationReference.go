package interfacesawscloudwatch


// A reference to a DomainForOrganization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainForOrganizationReference := &DomainForOrganizationReference{
//   	DomainForOrganizationArn: jsii.String("domainForOrganizationArn"),
//   }
//
type DomainForOrganizationReference struct {
	// The Arn of the DomainForOrganization resource.
	DomainForOrganizationArn *string `field:"required" json:"domainForOrganizationArn" yaml:"domainForOrganizationArn"`
}

