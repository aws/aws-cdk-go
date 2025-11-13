package interfacesawsconfig


// A reference to a OrganizationConformancePack resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationConformancePackReference := &OrganizationConformancePackReference{
//   	OrganizationConformancePackName: jsii.String("organizationConformancePackName"),
//   }
//
type OrganizationConformancePackReference struct {
	// The OrganizationConformancePackName of the OrganizationConformancePack resource.
	OrganizationConformancePackName *string `field:"required" json:"organizationConformancePackName" yaml:"organizationConformancePackName"`
}

