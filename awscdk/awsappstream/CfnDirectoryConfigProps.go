package awsappstream


// Properties for defining a `CfnDirectoryConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDirectoryConfigProps := &CfnDirectoryConfigProps{
//   	DirectoryName: jsii.String("directoryName"),
//   	OrganizationalUnitDistinguishedNames: []*string{
//   		jsii.String("organizationalUnitDistinguishedNames"),
//   	},
//   	ServiceAccountCredentials: &ServiceAccountCredentialsProperty{
//   		AccountName: jsii.String("accountName"),
//   		AccountPassword: jsii.String("accountPassword"),
//   	},
//
//   	// the properties below are optional
//   	CertificateBasedAuthProperties: &CertificateBasedAuthPropertiesProperty{
//   		CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html
//
type CfnDirectoryConfigProps struct {
	// The fully qualified name of the directory (for example, corp.example.com).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-directoryname
	//
	DirectoryName *string `field:"required" json:"directoryName" yaml:"directoryName"`
	// The distinguished names of the organizational units for computer accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-organizationalunitdistinguishednames
	//
	OrganizationalUnitDistinguishedNames *[]*string `field:"required" json:"organizationalUnitDistinguishedNames" yaml:"organizationalUnitDistinguishedNames"`
	// The credentials for the service account used by the streaming instance to connect to the directory.
	//
	// Do not use this parameter directly. Use `ServiceAccountCredentials` as an input parameter with `noEcho` as shown in the [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html) . For best practices information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-serviceaccountcredentials
	//
	ServiceAccountCredentials interface{} `field:"required" json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
	// The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-certificatebasedauthproperties
	//
	CertificateBasedAuthProperties interface{} `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
}

