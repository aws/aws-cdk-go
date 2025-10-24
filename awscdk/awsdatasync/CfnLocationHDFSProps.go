package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationHDFS`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationHDFSProps := &CfnLocationHDFSProps{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	AuthenticationType: jsii.String("authenticationType"),
//   	NameNodes: []interface{}{
//   		&NameNodeProperty{
//   			Hostname: jsii.String("hostname"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	BlockSize: jsii.Number(123),
//   	KerberosKeytab: jsii.String("kerberosKeytab"),
//   	KerberosKrb5Conf: jsii.String("kerberosKrb5Conf"),
//   	KerberosPrincipal: jsii.String("kerberosPrincipal"),
//   	KmsKeyProviderUri: jsii.String("kmsKeyProviderUri"),
//   	QopConfiguration: &QopConfigurationProperty{
//   		DataTransferProtection: jsii.String("dataTransferProtection"),
//   		RpcProtection: jsii.String("rpcProtection"),
//   	},
//   	ReplicationFactor: jsii.Number(123),
//   	SimpleUser: jsii.String("simpleUser"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html
//
type CfnLocationHDFSProps struct {
	// The Amazon Resource Names (ARNs) of the DataSync agents that can connect to your HDFS cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-agentarns
	//
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The authentication mode used to determine identity of user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The NameNode that manages the HDFS namespace.
	//
	// The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-namenodes
	//
	NameNodes interface{} `field:"required" json:"nameNodes" yaml:"nameNodes"`
	// The size of data blocks to write into the HDFS cluster.
	//
	// The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-blocksize
	//
	BlockSize *float64 `field:"optional" json:"blockSize" yaml:"blockSize"`
	// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys.
	//
	// Provide the base64-encoded file text. If `KERBEROS` is specified for `AuthType` , this value is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-kerberoskeytab
	//
	KerberosKeytab *string `field:"optional" json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// The `krb5.conf` file that contains the Kerberos configuration information. You can load the `krb5.conf` by providing a string of the file's contents or an Amazon S3 presigned URL of the file. If `KERBEROS` is specified for `AuthType` , this value is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-kerberoskrb5conf
	//
	KerberosKrb5Conf *string `field:"optional" json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// The Kerberos principal with access to the files and folders on the HDFS cluster.
	//
	// > If `KERBEROS` is specified for `AuthenticationType` , this parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-kerberosprincipal
	//
	KerberosPrincipal *string `field:"optional" json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// The URI of the HDFS cluster's Key Management Server (KMS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-kmskeyprovideruri
	//
	KmsKeyProviderUri *string `field:"optional" json:"kmsKeyProviderUri" yaml:"kmsKeyProviderUri"`
	// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster.
	//
	// If `QopConfiguration` isn't specified, `RpcProtection` and `DataTransferProtection` default to `PRIVACY` . If you set `RpcProtection` or `DataTransferProtection` , the other parameter assumes the same value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-qopconfiguration
	//
	QopConfiguration interface{} `field:"optional" json:"qopConfiguration" yaml:"qopConfiguration"`
	// The number of DataNodes to replicate the data to when writing to the HDFS cluster.
	//
	// By default, data is replicated to three DataNodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-replicationfactor
	//
	// Default: - 3.
	//
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// The user name used to identify the client on the host operating system.
	//
	// > If `SIMPLE` is specified for `AuthenticationType` , this parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-simpleuser
	//
	SimpleUser *string `field:"optional" json:"simpleUser" yaml:"simpleUser"`
	// A subdirectory in the HDFS cluster.
	//
	// This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to `/` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationhdfs.html#cfn-datasync-locationhdfs-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

