package awsdatasync


// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer privacy settings configured on the Hadoop Distributed File System (HDFS) cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   qopConfigurationProperty := &qopConfigurationProperty{
//   	dataTransferProtection: jsii.String("dataTransferProtection"),
//   	rpcProtection: jsii.String("rpcProtection"),
//   }
//
type CfnLocationHDFS_QopConfigurationProperty struct {
	// The data transfer protection setting configured on the HDFS cluster.
	//
	// This setting corresponds to your `dfs.data.transfer.protection` setting in the `hdfs-site.xml` file on your Hadoop cluster.
	DataTransferProtection *string `field:"optional" json:"dataTransferProtection" yaml:"dataTransferProtection"`
	// The Remote Procedure Call (RPC) protection setting configured on the HDFS cluster.
	//
	// This setting corresponds to your `hadoop.rpc.protection` setting in your `core-site.xml` file on your Hadoop cluster.
	RpcProtection *string `field:"optional" json:"rpcProtection" yaml:"rpcProtection"`
}

