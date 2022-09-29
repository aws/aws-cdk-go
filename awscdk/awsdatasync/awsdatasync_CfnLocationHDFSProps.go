package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationHDFS`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationHDFSProps := &cfnLocationHDFSProps{
//   	agentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	authenticationType: jsii.String("authenticationType"),
//   	nameNodes: []interface{}{
//   		&nameNodeProperty{
//   			hostname: jsii.String("hostname"),
//   			port: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	blockSize: jsii.Number(123),
//   	kerberosKeytab: jsii.String("kerberosKeytab"),
//   	kerberosKrb5Conf: jsii.String("kerberosKrb5Conf"),
//   	kerberosPrincipal: jsii.String("kerberosPrincipal"),
//   	kmsKeyProviderUri: jsii.String("kmsKeyProviderUri"),
//   	qopConfiguration: &qopConfigurationProperty{
//   		dataTransferProtection: jsii.String("dataTransferProtection"),
//   		rpcProtection: jsii.String("rpcProtection"),
//   	},
//   	replicationFactor: jsii.Number(123),
//   	simpleUser: jsii.String("simpleUser"),
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationHDFSProps struct {
	// The Amazon Resource Names (ARNs) of the agents that are used to connect to the HDFS cluster.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// `AWS::DataSync::LocationHDFS.AuthenticationType`.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The NameNode that manages the HDFS namespace.
	//
	// The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode.
	NameNodes interface{} `field:"required" json:"nameNodes" yaml:"nameNodes"`
	// The size of data blocks to write into the HDFS cluster.
	//
	// The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
	BlockSize *float64 `field:"optional" json:"blockSize" yaml:"blockSize"`
	// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys.
	//
	// Provide the base64-encoded file text. If `KERBEROS` is specified for `AuthType` , this value is required.
	KerberosKeytab *string `field:"optional" json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// The `krb5.conf` file that contains the Kerberos configuration information. You can load the `krb5.conf` by providing a string of the file's contents or an Amazon S3 presigned URL of the file. If `KERBEROS` is specified for `AuthType` , this value is required.
	KerberosKrb5Conf *string `field:"optional" json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// The Kerberos principal with access to the files and folders on the HDFS cluster.
	//
	// > If `KERBEROS` is specified for `AuthenticationType` , this parameter is required.
	KerberosPrincipal *string `field:"optional" json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri *string `field:"optional" json:"kmsKeyProviderUri" yaml:"kmsKeyProviderUri"`
	// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster.
	//
	// If `QopConfiguration` isn't specified, `RpcProtection` and `DataTransferProtection` default to `PRIVACY` . If you set `RpcProtection` or `DataTransferProtection` , the other parameter assumes the same value.
	QopConfiguration interface{} `field:"optional" json:"qopConfiguration" yaml:"qopConfiguration"`
	// The number of DataNodes to replicate the data to when writing to the HDFS cluster.
	//
	// By default, data is replicated to three DataNodes.
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// The user name used to identify the client on the host operating system.
	//
	// > If `SIMPLE` is specified for `AuthenticationType` , this parameter is required.
	SimpleUser *string `field:"optional" json:"simpleUser" yaml:"simpleUser"`
	// A subdirectory in the HDFS cluster.
	//
	// This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to `/` .
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

