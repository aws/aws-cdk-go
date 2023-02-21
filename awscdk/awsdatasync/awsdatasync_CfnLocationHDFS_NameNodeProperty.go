package awsdatasync


// The NameNode of the Hadoop Distributed File System (HDFS).
//
// The NameNode manages the file system's namespace and performs operations such as opening, closing, and renaming files and directories. The NameNode also contains the information to map blocks of data to the DataNodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nameNodeProperty := &nameNodeProperty{
//   	hostname: jsii.String("hostname"),
//   	port: jsii.Number(123),
//   }
//
type CfnLocationHDFS_NameNodeProperty struct {
	// The hostname of the NameNode in the HDFS cluster.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the NameNode. An agent that's installed on-premises uses this hostname to communicate with the NameNode in the network.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The port that the NameNode uses to listen to client requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

