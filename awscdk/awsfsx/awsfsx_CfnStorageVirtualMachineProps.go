package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStorageVirtualMachine`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStorageVirtualMachineProps := &cfnStorageVirtualMachineProps{
//   	fileSystemId: jsii.String("fileSystemId"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   		netBiosName: jsii.String("netBiosName"),
//   		selfManagedActiveDirectoryConfiguration: &selfManagedActiveDirectoryConfigurationProperty{
//   			dnsIps: []*string{
//   				jsii.String("dnsIps"),
//   			},
//   			domainName: jsii.String("domainName"),
//   			fileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   			organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   			password: jsii.String("password"),
//   			userName: jsii.String("userName"),
//   		},
//   	},
//   	rootVolumeSecurityStyle: jsii.String("rootVolumeSecurityStyle"),
//   	svmAdminPassword: jsii.String("svmAdminPassword"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStorageVirtualMachineProps struct {
	// Specifies the FSx for ONTAP file system on which to create the SVM.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The name of the SVM.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the Microsoft Active Directory configuration to which the SVM is joined, if applicable.
	ActiveDirectoryConfiguration interface{} `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// The security style of the root volume of the SVM. Specify one of the following values:.
	//
	// - `UNIX` if the file system is managed by a UNIX administrator, the majority of users are NFS clients, and an application accessing the data uses a UNIX user as the service account.
	// - `NTFS` if the file system is managed by a Windows administrator, the majority of users are SMB clients, and an application accessing the data uses a Windows user as the service account.
	// - `MIXED` if the file system is managed by both UNIX and Windows administrators and users consist of both NFS and SMB clients.
	RootVolumeSecurityStyle *string `field:"optional" json:"rootVolumeSecurityStyle" yaml:"rootVolumeSecurityStyle"`
	// Specifies the password to use when logging on to the SVM using a secure shell (SSH) connection to the SVM's management endpoint.
	//
	// Doing so enables you to manage the SVM using the NetApp ONTAP CLI or REST API. If you do not specify a password, you can still use the file system's `fsxadmin` user to manage the SVM. For more information, see [Managing SVMs using the NetApp ONTAP CLI](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-resources-ontap-apps.html#vsadmin-ontap-cli) in the *FSx for ONTAP User Guide* .
	SvmAdminPassword *string `field:"optional" json:"svmAdminPassword" yaml:"svmAdminPassword"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

