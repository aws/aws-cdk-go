package awsec2


// Options for creating a SystemD configuration file.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//
//   	Init: ec2.CloudFormationInit_FromElements(ec2.InitService_SystemdConfigFile(jsii.String("simpleserver"), &SystemdConfigFileOptions{
//   		Command: jsii.String("/usr/bin/python3 -m http.server 8080"),
//   		Cwd: jsii.String("/var/www/html"),
//   	}), ec2.InitService_Enable(jsii.String("simpleserver"), &InitServiceOptions{
//   		ServiceManager: ec2.ServiceManager_SYSTEMD,
//   	}), ec2.InitFile_FromString(jsii.String("/var/www/html/index.html"), jsii.String("Hello! It's working!"))),
//   })
//
type SystemdConfigFileOptions struct {
	// The command to run to start this service.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Start the service after the networking part of the OS comes up.
	// Default: true.
	//
	AfterNetwork *bool `field:"optional" json:"afterNetwork" yaml:"afterNetwork"`
	// The working directory for the command.
	// Default: Root directory or home directory of specified user.
	//
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// A description of this service.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The group to execute the process under.
	// Default: root.
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Keep the process running all the time.
	//
	// Restarts the process when it exits for any reason other
	// than the machine shutting down.
	// Default: true.
	//
	KeepRunning *bool `field:"optional" json:"keepRunning" yaml:"keepRunning"`
	// The user to execute the process under.
	// Default: root.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
}

