package awscodepipelineactions


// Properties of `Ec2DeploySpecifications.inline()`.
//
// Example:
//   sourceOutput := codepipeline.NewArtifact()
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   })
//   deployAction := codepipeline_actions.NewEc2DeployAction(&Ec2DeployActionProps{
//   	ActionName: jsii.String("Ec2Deploy"),
//   	Input: sourceOutput,
//   	InstanceType: codepipeline_actions.Ec2InstanceType_EC2,
//   	InstanceTagKey: jsii.String("Name"),
//   	InstanceTagValue: jsii.String("MyInstance"),
//   	DeploySpecifications: codepipeline_actions.Ec2DeploySpecifications_Inline(&Ec2DeploySpecificationsInlineProps{
//   		TargetDirectory: jsii.String("/home/ec2-user/deploy"),
//   		PreScript: jsii.String("scripts/pre-deploy.sh"),
//   		PostScript: jsii.String("scripts/post-deploy.sh"),
//   	}),
//   })
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type Ec2DeploySpecificationsInlineProps struct {
	// Path to the executable script file that runs AFTER the Deploy phase.
	//
	// It should start from the root directory of your uploaded source artifact.
	// Use an absolute path like `uploadDir/postScript.sh`.
	PostScript *string `field:"required" json:"postScript" yaml:"postScript"`
	// The location of the target directory you want to deploy to.
	//
	// Use an absolute path like `/home/ec2-user/deploy`.
	TargetDirectory *string `field:"required" json:"targetDirectory" yaml:"targetDirectory"`
	// Path to the executable script file that runs BEFORE the Deploy phase.
	//
	// It should start from the root directory of your uploaded source artifact.
	// Use an absolute path like `uploadDir/preScript.sh`.
	// Default: - No script.
	//
	PreScript *string `field:"optional" json:"preScript" yaml:"preScript"`
}

