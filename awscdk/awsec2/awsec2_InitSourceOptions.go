package awsec2


// Additional options for an InitSource.
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
type InitSourceOptions struct {
	// Restart the given services after this archive has been extracted.
	// Experimental.
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
}

