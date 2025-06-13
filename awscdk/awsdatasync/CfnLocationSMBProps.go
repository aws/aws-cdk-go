package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationSMB`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationSMBProps := &CfnLocationSMBProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//
//   	// the properties below are optional
//   	AuthenticationType: jsii.String("authenticationType"),
//   	DnsIpAddresses: []*string{
//   		jsii.String("dnsIpAddresses"),
//   	},
//   	Domain: jsii.String("domain"),
//   	KerberosKeytab: jsii.String("kerberosKeytab"),
//   	KerberosKrb5Conf: jsii.String("kerberosKrb5Conf"),
//   	KerberosPrincipal: jsii.String("kerberosPrincipal"),
//   	MountOptions: &MountOptionsProperty{
//   		Version: jsii.String("version"),
//   	},
//   	Password: jsii.String("password"),
//   	ServerHostname: jsii.String("serverHostname"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	User: jsii.String("user"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html
//
type CfnLocationSMBProps struct {
	// Specifies the DataSync agent (or agents) that can connect to your SMB file server.
	//
	// You specify an agent by using its Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-agentarns
	//
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The authentication mode used to determine identity of user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to.
	//
	// This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-dnsipaddresses
	//
	DnsIpAddresses *[]*string `field:"optional" json:"dnsIpAddresses" yaml:"dnsIpAddresses"`
	// Specifies the Windows domain name that your SMB file server belongs to.
	//
	// This parameter applies only if `AuthenticationType` is set to `NTLM` .
	//
	// If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right file server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The Base64 string representation of the Keytab file.
	//
	// Specifies your Kerberos key table (keytab) file, which includes mappings between your service principal name (SPN) and encryption keys. To avoid task execution errors, make sure that the SPN in the keytab file matches exactly what you specify for KerberosPrincipal and in your krb5.conf file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-kerberoskeytab
	//
	KerberosKeytab *string `field:"optional" json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket. Specifies a Kerberos configuration file (krb5.conf) that defines your Kerberos realm configuration. To avoid task execution errors, make sure that the service principal name (SPN) in the krb5.conf file matches exactly what you specify for KerberosPrincipal and in your keytab file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-kerberoskrb5conf
	//
	KerberosKrb5Conf *string `field:"optional" json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server.
	//
	// SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-kerberosprincipal
	//
	KerberosPrincipal *string `field:"optional" json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// Specifies the version of the SMB protocol that DataSync uses to access your SMB file server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-mountoptions
	//
	MountOptions interface{} `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// Specifies the password of the user who can mount your SMB file server and has permission to access the files and folders involved in your transfer.
	//
	// This parameter applies only if `AuthenticationType` is set to `NTLM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies the domain name or IP address of the SMB file server that your DataSync agent connects to.
	//
	// Remember the following when configuring this parameter:
	//
	// - You can't specify an IP version 6 (IPv6) address.
	// - If you're using Kerberos authentication, you must specify a domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-serverhostname
	//
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the name of the share exported by your SMB file server where DataSync will read or write data.
	//
	// You can include a subdirectory in the share path (for example, `/path/to/subdirectory` ). Make sure that other SMB clients in your network can also mount this path.
	//
	// To copy all data in the subdirectory, DataSync must be able to mount the SMB share and access all of its data. For more information, see [Providing DataSync access to SMB file servers](https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the user that can mount and access the files, folders, and file metadata in your SMB file server.
	//
	// This parameter applies only if `AuthenticationType` is set to `NTLM` .
	//
	// For information about choosing a user with the right level of access for your transfer, see [Providing DataSync access to SMB file servers](https://docs.aws.amazon.com/datasync/latest/userguide/create-smb-location.html#configuring-smb-permissions) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationsmb.html#cfn-datasync-locationsmb-user
	//
	User *string `field:"optional" json:"user" yaml:"user"`
}

