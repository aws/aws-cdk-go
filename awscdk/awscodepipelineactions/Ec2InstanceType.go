package awscodepipelineactions


// The type of instances or SSM nodes created in Amazon EC2.
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
//   	Actions: []IAction{
//   		deployAction,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-EC2Deploy.html#action-reference-EC2Deploy-parameters
//
type Ec2InstanceType string

const (
	// Amazon EC2 instances.
	Ec2InstanceType_EC2 Ec2InstanceType = "EC2"
	// AWS System Manager (SSM) managed nodes.
	Ec2InstanceType_SSM_MANAGED_NODE Ec2InstanceType = "SSM_MANAGED_NODE"
)

