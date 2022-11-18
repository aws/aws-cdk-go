package awsec2


// Options for an InitService.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type InitServiceOptions struct {
	// Enable or disable this service.
	//
	// Set to true to ensure that the service will be started automatically upon boot.
	//
	// Set to false to ensure that the service will not be started automatically upon boot.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Make sure this service is running or not running after cfn-init finishes.
	//
	// Set to true to ensure that the service is running after cfn-init finishes.
	//
	// Set to false to ensure that the service is not running after cfn-init finishes.
	EnsureRunning *bool `field:"optional" json:"ensureRunning" yaml:"ensureRunning"`
	// Restart service when the actions registered into the restartHandle have been performed.
	//
	// Register actions into the restartHandle by passing it to `InitFile`, `InitCommand`,
	// `InitPackage` and `InitSource` objects.
	ServiceRestartHandle InitServiceRestartHandle `field:"optional" json:"serviceRestartHandle" yaml:"serviceRestartHandle"`
}

