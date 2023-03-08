package awsfsx


// The configuration that Amazon FSx uses to join a FSx for Windows File Server file system or an ONTAP storage virtual machine (SVM) to a self-managed (including on-premises) Microsoft Active Directory (AD) directory.
//
// For more information, see [Using Amazon FSx with your self-managed Microsoft Active Directory](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html) or [Managing SVMs](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedActiveDirectoryConfigurationProperty := &selfManagedActiveDirectoryConfigurationProperty{
//   	dnsIps: []*string{
//   		jsii.String("dnsIps"),
//   	},
//   	domainName: jsii.String("domainName"),
//   	fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   	organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	password: jsii.String("password"),
//   	userName: jsii.String("userName"),
//   }
//
type CfnStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty struct {
	// A list of up to three IP addresses of DNS servers or domain controllers in the self-managed AD directory.
	DnsIps *[]*string `field:"optional" json:"dnsIps" yaml:"dnsIps"`
	// The fully qualified domain name of the self-managed AD directory, such as `corp.example.com` .
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// (Optional) The name of the domain group whose members are granted administrative privileges for the file system.
	//
	// Administrative privileges include taking ownership of files and folders, setting audit controls (audit ACLs) on files and folders, and administering the file system remotely by using the FSx Remote PowerShell. The group that you specify must already exist in your domain. If you don't provide one, your AD domain's Domain Admins group is used.
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// (Optional) The fully qualified distinguished name of the organizational unit within your self-managed AD directory.
	//
	// Amazon FSx only accepts OU as the direct parent of the file system. An example is `OU=FSx,DC=yourdomain,DC=corp,DC=com` . To learn more, see [RFC 2253](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc2253) . If none is provided, the FSx file system is created in the default location of your self-managed AD directory.
	//
	// > Only Organizational Unit (OU) objects can be the direct parent of the file system that you're creating.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain.
	//
	// This account must have the permission to join computers to the domain in the organizational unit provided in `OrganizationalUnitDistinguishedName` , or in the default location of your AD domain.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

