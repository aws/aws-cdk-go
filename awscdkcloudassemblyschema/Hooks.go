package awscdkcloudassemblyschema


// Commands to run at predefined points during the integration test workflow.
type Hooks struct {
	// Commands to run prior after deploying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PostDeploy *[]*string `field:"optional" json:"postDeploy" yaml:"postDeploy"`
	// Commands to run after destroying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PostDestroy *[]*string `field:"optional" json:"postDestroy" yaml:"postDestroy"`
	// Commands to run prior to deploying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PreDeploy *[]*string `field:"optional" json:"preDeploy" yaml:"preDeploy"`
	// Commands to run prior to destroying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PreDestroy *[]*string `field:"optional" json:"preDestroy" yaml:"preDestroy"`
}

