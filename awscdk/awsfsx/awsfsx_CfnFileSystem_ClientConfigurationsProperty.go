package awsfsx


// Specifies who can mount an OpenZFS file system and the options available while mounting the file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientConfigurationsProperty := &clientConfigurationsProperty{
//   	clients: jsii.String("clients"),
//   	options: []*string{
//   		jsii.String("options"),
//   	},
//   }
//
type CfnFileSystem_ClientConfigurationsProperty struct {
	// A value that specifies who can mount the file system.
	//
	// You can provide a wildcard character ( `*` ), an IP address ( `0.0.0.0` ), or a CIDR address ( `192.0.2.0/24` ). By default, Amazon FSx uses the wildcard character when specifying the client.
	Clients *string `field:"optional" json:"clients" yaml:"clients"`
	// The options to use when mounting the file system.
	//
	// For a list of options that you can use with Network File System (NFS), see the [exports(5) - Linux man page](https://docs.aws.amazon.com/https://linux.die.net/man/5/exports) . When choosing your options, consider the following:
	//
	// - `crossmnt` is used by default. If you don't specify `crossmnt` when changing the client configuration, you won't be able to see or access snapshots in your file system's snapshot directory.
	// - `sync` is used by default. If you instead specify `async` , the system acknowledges writes before writing to disk. If the system crashes before the writes are finished, you lose the unwritten data.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
}

