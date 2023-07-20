package awsfsx


// Describes the self-managed Microsoft Active Directory to which you want to join the SVM.
//
// Joining an Active Directory provides user authentication and access control for SMB clients, including Microsoft Windows and macOS client accessing the file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryConfigurationProperty := &ActiveDirectoryConfigurationProperty{
//   	NetBiosName: jsii.String("netBiosName"),
//   	SelfManagedActiveDirectoryConfiguration: &SelfManagedActiveDirectoryConfigurationProperty{
//   		DnsIps: []*string{
//   			jsii.String("dnsIps"),
//   		},
//   		DomainName: jsii.String("domainName"),
//   		FileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   		OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		Password: jsii.String("password"),
//   		UserName: jsii.String("userName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html
//
type CfnStorageVirtualMachine_ActiveDirectoryConfigurationProperty struct {
	// The NetBIOS name of the Active Directory computer object that will be created for your SVM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-activedirectoryconfiguration-netbiosname
	//
	NetBiosName *string `field:"optional" json:"netBiosName" yaml:"netBiosName"`
	// The configuration that Amazon FSx uses to join the ONTAP storage virtual machine (SVM) to your self-managed (including on-premises) Microsoft Active Directory (AD) directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-activedirectoryconfiguration-selfmanagedactivedirectoryconfiguration
	//
	SelfManagedActiveDirectoryConfiguration interface{} `field:"optional" json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

