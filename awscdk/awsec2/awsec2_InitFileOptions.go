package awsec2


// Options for InitFile.
//
// Example:
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &initFileOptions{
//   	serviceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initSource.fromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &initSourceOptions{
//   	serviceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initService.enable(jsii.String("nginx"), &initServiceOptions{
//   	serviceRestartHandle: handle,
//   }))
//
// Experimental.
type InitFileOptions struct {
	// True if the inlined content (from a string or file) should be treated as base64 encoded.
	//
	// Only applicable for inlined string and file content.
	// Experimental.
	Base64Encoded *bool `field:"optional" json:"base64Encoded" yaml:"base64Encoded"`
	// The name of the owning group for this file.
	//
	// Not supported for Windows systems.
	// Experimental.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// A six-digit octal value representing the mode for this file.
	//
	// Use the first three digits for symlinks and the last three digits for
	// setting permissions. To create a symlink, specify 120xxx, where xxx
	// defines the permissions of the target file. To specify permissions for a
	// file, use the last three digits, such as 000644.
	//
	// Not supported for Windows systems.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of the owning user for this file.
	//
	// Not supported for Windows systems.
	// Experimental.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Restart the given service after this file has been written.
	// Experimental.
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
}

