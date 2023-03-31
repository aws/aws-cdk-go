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
//   activeDirectoryConfigurationProperty := &activeDirectoryConfigurationProperty{
//   	netBiosName: jsii.String("netBiosName"),
//   	selfManagedActiveDirectoryConfiguration: &selfManagedActiveDirectoryConfigurationProperty{
//   		dnsIps: []*string{
//   			jsii.String("dnsIps"),
//   		},
//   		domainName: jsii.String("domainName"),
//   		fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		password: jsii.String("password"),
//   		userName: jsii.String("userName"),
//   	},
//   }
//
type CfnStorageVirtualMachine_ActiveDirectoryConfigurationProperty struct {
	// The NetBIOS name of the Active Directory computer object that will be created for your SVM.
	NetBiosName *string `field:"optional" json:"netBiosName" yaml:"netBiosName"`
	// The configuration that Amazon FSx uses to join the ONTAP storage virtual machine (SVM) to your self-managed (including on-premises) Microsoft Active Directory (AD) directory.
	SelfManagedActiveDirectoryConfiguration interface{} `field:"optional" json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

