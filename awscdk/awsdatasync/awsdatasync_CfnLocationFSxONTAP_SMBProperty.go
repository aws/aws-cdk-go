package awsdatasync


// Specifies the Server Message Block (SMB) protocol configuration that AWS DataSync uses to access a storage virtual machine (SVM) on your Amazon FSx for NetApp ONTAP file system.
//
// For more information, see [Accessing FSx for ONTAP file systems](https://docs.aws.amazon.com/datasync/latest/userguide/create-ontap-location.html#create-ontap-location-access) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sMBProperty := &sMBProperty{
//   	mountOptions: &smbMountOptionsProperty{
//   		version: jsii.String("version"),
//   	},
//   	password: jsii.String("password"),
//   	user: jsii.String("user"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
//   }
//
type CfnLocationFSxONTAP_SMBProperty struct {
	// Specifies how DataSync can access a location using the SMB protocol.
	MountOptions interface{} `field:"required" json:"mountOptions" yaml:"mountOptions"`
	// Specifies the password of a user who has permission to access your SVM.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies a user name that can mount the location and access the files, folders, and metadata that you need in the SVM.
	//
	// If you provide a user in your Active Directory, note the following:
	//
	// - If you're using AWS Directory Service for Microsoft Active Directory , the user must be a member of the AWS Delegated FSx Administrators group.
	// - If you're using a self-managed Active Directory, the user must be a member of either the Domain Admins group or a custom group that you specified for file system administration when you created your file system.
	//
	// Make sure that the user has the permissions it needs to copy the data you want:
	//
	// - `SE_TCB_NAME` : Required to set object ownership and file metadata. With this privilege, you also can copy NTFS discretionary access lists (DACLs).
	// - `SE_SECURITY_NAME` : May be needed to copy NTFS system access control lists (SACLs). This operation specifically requires the Windows privilege, which is granted to members of the Domain Admins group. If you configure your task to copy SACLs, make sure that the user has the required privileges. For information about copying SACLs, see [Ownership and permissions-related options](https://docs.aws.amazon.com/datasync/latest/userguide/create-task.html#configure-ownership-and-permissions) .
	User *string `field:"required" json:"user" yaml:"user"`
	// Specifies the fully qualified domain name (FQDN) of the Microsoft Active Directory that your storage virtual machine (SVM) belongs to.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

