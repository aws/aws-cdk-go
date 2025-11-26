package previewawsfsxmixins


// The configuration that Amazon FSx uses to join the ONTAP storage virtual machine (SVM) to your self-managed (including on-premises) Microsoft Active Directory directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   selfManagedActiveDirectoryConfigurationProperty := &SelfManagedActiveDirectoryConfigurationProperty{
//   	DnsIps: []*string{
//   		jsii.String("dnsIps"),
//   	},
//   	DomainJoinServiceAccountSecret: jsii.String("domainJoinServiceAccountSecret"),
//   	DomainName: jsii.String("domainName"),
//   	FileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   	OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	Password: jsii.String("password"),
//   	UserName: jsii.String("userName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html
//
type CfnStorageVirtualMachinePropsMixin_SelfManagedActiveDirectoryConfigurationProperty struct {
	// A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-dnsips
	//
	DnsIps *[]*string `field:"optional" json:"dnsIps" yaml:"dnsIps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-domainjoinserviceaccountsecret
	//
	DomainJoinServiceAccountSecret *string `field:"optional" json:"domainJoinServiceAccountSecret" yaml:"domainJoinServiceAccountSecret"`
	// The fully qualified domain name of the self-managed AD directory, such as `corp.example.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// (Optional) The name of the domain group whose members are granted administrative privileges for the file system.
	//
	// Administrative privileges include taking ownership of files and folders, setting audit controls (audit ACLs) on files and folders, and administering the file system remotely by using the FSx Remote PowerShell. The group that you specify must already exist in your domain. If you don't provide one, your AD domain's Domain Admins group is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-filesystemadministratorsgroup
	//
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// (Optional) The fully qualified distinguished name of the organizational unit within your self-managed AD directory.
	//
	// Amazon FSx only accepts OU as the direct parent of the file system. An example is `OU=FSx,DC=yourdomain,DC=corp,DC=com` . To learn more, see [RFC 2253](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc2253) . If none is provided, the FSx file system is created in the default location of your self-managed AD directory.
	//
	// > Only Organizational Unit (OU) objects can be the direct parent of the file system that you're creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-organizationalunitdistinguishedname
	//
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// This account must have the permission to join computers to the domain in the organizational unit provided in `OrganizationalUnitDistinguishedName` , or in the default location of your AD domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration-username
	//
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

