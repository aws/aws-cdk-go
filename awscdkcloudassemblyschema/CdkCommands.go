package awscdkcloudassemblyschema


// Options for specific cdk commands that are run as part of the integration test workflow.
type CdkCommands struct {
	// Options to for the cdk deploy command.
	// Default: - default deploy options.
	//
	Deploy *DeployCommand `field:"optional" json:"deploy" yaml:"deploy"`
	// Options to for the cdk destroy command.
	// Default: - default destroy options.
	//
	Destroy *DestroyCommand `field:"optional" json:"destroy" yaml:"destroy"`
}

