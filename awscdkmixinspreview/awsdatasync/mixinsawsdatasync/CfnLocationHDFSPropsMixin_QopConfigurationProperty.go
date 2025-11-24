package mixinsawsdatasync


// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer privacy settings configured on the Hadoop Distributed File System (HDFS) cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   qopConfigurationProperty := &QopConfigurationProperty{
//   	DataTransferProtection: jsii.String("dataTransferProtection"),
//   	RpcProtection: jsii.String("rpcProtection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationhdfs-qopconfiguration.html
//
type CfnLocationHDFSPropsMixin_QopConfigurationProperty struct {
	// The data transfer protection setting configured on the HDFS cluster.
	//
	// This setting corresponds to your `dfs.data.transfer.protection` setting in the `hdfs-site.xml` file on your Hadoop cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationhdfs-qopconfiguration.html#cfn-datasync-locationhdfs-qopconfiguration-datatransferprotection
	//
	// Default: - "PRIVACY".
	//
	DataTransferProtection *string `field:"optional" json:"dataTransferProtection" yaml:"dataTransferProtection"`
	// The Remote Procedure Call (RPC) protection setting configured on the HDFS cluster.
	//
	// This setting corresponds to your `hadoop.rpc.protection` setting in your `core-site.xml` file on your Hadoop cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationhdfs-qopconfiguration.html#cfn-datasync-locationhdfs-qopconfiguration-rpcprotection
	//
	// Default: - "PRIVACY".
	//
	RpcProtection *string `field:"optional" json:"rpcProtection" yaml:"rpcProtection"`
}

