package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of `Ec2DeployAction`.
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
type Ec2DeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Default: 1.
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names,
	// if any of the action's variables were referenced - otherwise,
	// no namespace will be set.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	// Default: a new Role will be generated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The deploy specifications.
	DeploySpecifications Ec2DeploySpecifications `field:"required" json:"deploySpecifications" yaml:"deploySpecifications"`
	// The input artifact to deploy to EC2 instances.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The tag key of the instances that you created in Amazon EC2.
	InstanceTagKey *string `field:"required" json:"instanceTagKey" yaml:"instanceTagKey"`
	// The type of instances or SSM nodes created in Amazon EC2.
	//
	// You must have already created, tagged, and installed the SSM agent on all instances.
	InstanceType Ec2InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// The tag value of the instances that you created in Amazon EC2.
	// Default: - all instances with `instanceTagKey` will be matched.
	//
	InstanceTagValue *string `field:"optional" json:"instanceTagValue" yaml:"instanceTagValue"`
	// The number or percentage of instances that can deploy in parallel.
	// Default: - No configuration.
	//
	MaxBatch Ec2MaxInstances `field:"optional" json:"maxBatch" yaml:"maxBatch"`
	// Stop the task after the task fails on the specified number or percentage of instances.
	// Default: - No configuration.
	//
	MaxError Ec2MaxInstances `field:"optional" json:"maxError" yaml:"maxError"`
	// The list of target groups for deployment. You must have already created the target groups.
	//
	// Target groups provide a set of instances to process specific requests.
	// If the target group is specified, instances will be removed from the target group before deployment and added back to the target group after deployment.
	// Default: - No target groups.
	//
	TargetGroups *[]awselasticloadbalancingv2.ITargetGroup `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

