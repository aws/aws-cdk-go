package awsdatasync


// The NFS mount options that DataSync can use to mount your NFS share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountOptionsProperty := &mountOptionsProperty{
//   	version: jsii.String("version"),
//   }
//
type CfnLocationNFS_MountOptionsProperty struct {
	// The specific NFS version that you want DataSync to use to mount your NFS share.
	//
	// If the server refuses to use the version specified, the sync will fail. If you don't specify a version, DataSync defaults to `AUTOMATIC` . That is, DataSync automatically selects a version based on negotiation with the NFS server.
	//
	// You can specify the following NFS versions:
	//
	// - *[NFSv3](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc1813)* - stateless protocol version that allows for asynchronous writes on the server.
	// - *[NFSv4.0](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc3530)* - stateful, firewall-friendly protocol version that supports delegations and pseudo file systems.
	// - *[NFSv4.1](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc5661)* - stateful protocol version that supports sessions, directory delegations, and parallel data processing. Version 4.1 also includes all features available in version 4.0.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

