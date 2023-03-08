package awsappstream


// The name of the directory and organizational unit (OU) to use to join a fleet to a Microsoft Active Directory domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainJoinInfoProperty := &domainJoinInfoProperty{
//   	directoryName: jsii.String("directoryName"),
//   	organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   }
//
type CfnFleet_DomainJoinInfoProperty struct {
	// The fully qualified name of the directory (for example, corp.example.com).
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// The distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

