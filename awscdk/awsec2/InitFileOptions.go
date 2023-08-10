package awsec2


// Options for InitFile.
//
// Example:
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.CloudFormationInit_FromElements(ec2.InitFile_FromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &InitFileOptions{
//   	ServiceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitSource_FromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &InitSourceOptions{
//   	ServiceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitService_Enable(jsii.String("nginx"), &InitServiceOptions{
//   	ServiceRestartHandle: handle,
//   }))
//
type InitFileOptions struct {
	// True if the inlined content (from a string or file) should be treated as base64 encoded.
	//
	// Only applicable for inlined string and file content.
	Base64Encoded *bool `field:"optional" json:"base64Encoded" yaml:"base64Encoded"`
	// The name of the owning group for this file.
	//
	// Not supported for Windows systems.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// A six-digit octal value representing the mode for this file.
	//
	// Use the first three digits for symlinks and the last three digits for
	// setting permissions. To create a symlink, specify 120xxx, where xxx
	// defines the permissions of the target file. To specify permissions for a
	// file, use the last three digits, such as 000644.
	//
	// Not supported for Windows systems.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of the owning user for this file.
	//
	// Not supported for Windows systems.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Restart the given service after this file has been written.
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
}

