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
//   sMBProperty := &SMBProperty{
//   	MountOptions: &SmbMountOptionsProperty{
//   		Version: jsii.String("version"),
//   	},
//   	User: jsii.String("user"),
//
//   	// the properties below are optional
//   	CmkSecretConfig: &CmkSecretConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CustomSecretConfig: &CustomSecretConfigProperty{
//   		SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	Domain: jsii.String("domain"),
//   	ManagedSecretConfig: &ManagedSecretConfigProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	Password: jsii.String("password"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html
//
type CfnLocationFSxONTAP_SMBProperty struct {
	// Specifies how DataSync can access a location using the SMB protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-mountoptions
	//
	MountOptions interface{} `field:"required" json:"mountOptions" yaml:"mountOptions"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-user
	//
	User *string `field:"required" json:"user" yaml:"user"`
	// Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-cmksecretconfig
	//
	CmkSecretConfig interface{} `field:"optional" json:"cmkSecretConfig" yaml:"cmkSecretConfig"`
	// Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-customsecretconfig
	//
	CustomSecretConfig interface{} `field:"optional" json:"customSecretConfig" yaml:"customSecretConfig"`
	// Specifies the name of the Windows domain that your storage virtual machine (SVM) belongs to.
	//
	// If you have multiple domains in your environment, configuring this setting makes sure that DataSync connects to the right SVM.
	//
	// If you have multiple Active Directory domains in your environment, configuring this parameter makes sure that DataSync connects to the right SVM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location.
	//
	// DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-managedsecretconfig
	//
	ManagedSecretConfig interface{} `field:"optional" json:"managedSecretConfig" yaml:"managedSecretConfig"`
	// Specifies the password of a user who has permission to access your SVM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxontap-smb.html#cfn-datasync-locationfsxontap-smb-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
}

