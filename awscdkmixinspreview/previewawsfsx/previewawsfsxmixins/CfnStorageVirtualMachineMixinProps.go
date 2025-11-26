package previewawsfsxmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStorageVirtualMachinePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStorageVirtualMachineMixinProps := &CfnStorageVirtualMachineMixinProps{
//   	ActiveDirectoryConfiguration: &ActiveDirectoryConfigurationProperty{
//   		NetBiosName: jsii.String("netBiosName"),
//   		SelfManagedActiveDirectoryConfiguration: &SelfManagedActiveDirectoryConfigurationProperty{
//   			DnsIps: []*string{
//   				jsii.String("dnsIps"),
//   			},
//   			DomainJoinServiceAccountSecret: jsii.String("domainJoinServiceAccountSecret"),
//   			DomainName: jsii.String("domainName"),
//   			FileSystemAdministratorsGroup: jsii.String("fileSystemAdministratorsGroup"),
//   			OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   			Password: jsii.String("password"),
//   			UserName: jsii.String("userName"),
//   		},
//   	},
//   	FileSystemId: jsii.String("fileSystemId"),
//   	Name: jsii.String("name"),
//   	RootVolumeSecurityStyle: jsii.String("rootVolumeSecurityStyle"),
//   	SvmAdminPassword: jsii.String("svmAdminPassword"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html
//
type CfnStorageVirtualMachineMixinProps struct {
	// Describes the Microsoft Active Directory configuration to which the SVM is joined, if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html#cfn-fsx-storagevirtualmachine-activedirectoryconfiguration
	//
	ActiveDirectoryConfiguration interface{} `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// Specifies the FSx for ONTAP file system on which to create the SVM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html#cfn-fsx-storagevirtualmachine-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The name of the SVM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html#cfn-fsx-storagevirtualmachine-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The security style of the root volume of the SVM. Specify one of the following values:.
	//
	// - `UNIX` if the file system is managed by a UNIX administrator, the majority of users are NFS clients, and an application accessing the data uses a UNIX user as the service account.
	// - `NTFS` if the file system is managed by a Microsoft Windows administrator, the majority of users are SMB clients, and an application accessing the data uses a Microsoft Windows user as the service account.
	// - `MIXED` This is an advanced setting. For more information, see [Volume security style](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-security-style.html) in the Amazon FSx for NetApp ONTAP User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html#cfn-fsx-storagevirtualmachine-rootvolumesecuritystyle
	//
	RootVolumeSecurityStyle *string `field:"optional" json:"rootVolumeSecurityStyle" yaml:"rootVolumeSecurityStyle"`
	// Specifies the password to use when logging on to the SVM using a secure shell (SSH) connection to the SVM's management endpoint.
	//
	// Doing so enables you to manage the SVM using the NetApp ONTAP CLI or REST API. If you do not specify a password, you can still use the file system's `fsxadmin` user to manage the SVM. For more information, see [Managing SVMs using the NetApp ONTAP CLI](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-resources-ontap-apps.html#vsadmin-ontap-cli) in the *FSx for ONTAP User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html#cfn-fsx-storagevirtualmachine-svmadminpassword
	//
	SvmAdminPassword *string `field:"optional" json:"svmAdminPassword" yaml:"svmAdminPassword"`
	// A list of `Tag` values, with a maximum of 50 elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-storagevirtualmachine.html#cfn-fsx-storagevirtualmachine-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

