package awsec2


// Options for InitCommand.
//
// Example:
//   handle := ec2.NewInitServiceRestartHandle()
//   ec2.CloudFormationInit_FromElements(ec2.InitCommand_ShellCommand(jsii.String("/usr/bin/custom-nginx-install.sh"), &InitCommandOptions{
//   	ServiceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitService_Enable(jsii.String("nginx"), &InitServiceOptions{
//   	ServiceRestartHandle: handle,
//   }))
//
type InitCommandOptions struct {
	// The working directory.
	// Default: - Use default working directory.
	//
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// Sets environment variables for the command.
	//
	// This property overwrites, rather than appends, the existing environment.
	// Default: - Use current environment.
	//
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Continue running if this command fails.
	// Default: false.
	//
	IgnoreErrors *bool `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Identifier key for this command.
	//
	// Commands are executed in lexicographical order of their key names.
	// Default: - Automatically generated based on index.
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Restart the given service(s) after this command has run.
	// Default: - Do not restart any service.
	//
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
	// Command to determine whether this command should be run.
	//
	// If the test passes (exits with error code of 0), the command is run.
	// Default: - Always run the command.
	//
	TestCmd *string `field:"optional" json:"testCmd" yaml:"testCmd"`
	// The duration to wait after a command has finished in case the command causes a reboot.
	//
	// Set this value to `InitCommandWaitDuration.none()` if you do not want to wait for every command;
	// `InitCommandWaitDuration.forever()` directs cfn-init to exit and resume only after the reboot is complete.
	//
	// For Windows systems only.
	// Default: - 60 seconds.
	//
	WaitAfterCompletion InitCommandWaitDuration `field:"optional" json:"waitAfterCompletion" yaml:"waitAfterCompletion"`
}

