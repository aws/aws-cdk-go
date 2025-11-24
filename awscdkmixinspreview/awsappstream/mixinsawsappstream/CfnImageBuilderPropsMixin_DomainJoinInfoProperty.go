package mixinsawsappstream


// The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   domainJoinInfoProperty := &DomainJoinInfoProperty{
//   	DirectoryName: jsii.String("directoryName"),
//   	OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-domainjoininfo.html
//
type CfnImageBuilderPropsMixin_DomainJoinInfoProperty struct {
	// The fully qualified name of the directory (for example, corp.example.com).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-domainjoininfo.html#cfn-appstream-imagebuilder-domainjoininfo-directoryname
	//
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// The distinguished name of the organizational unit for computer accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-domainjoininfo.html#cfn-appstream-imagebuilder-domainjoininfo-organizationalunitdistinguishedname
	//
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

