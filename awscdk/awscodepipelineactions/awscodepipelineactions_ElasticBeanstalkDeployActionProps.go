package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of the {@link ElasticBeanstalkDeployAction Elastic Beanstalk deploy CodePipeline Action}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceOutput := codepipeline.NewArtifact()
//   targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   deployAction := codepipeline_actions.NewElasticBeanstalkDeployAction(&elasticBeanstalkDeployActionProps{
//   	actionName: jsii.String("ElasticBeanstalkDeploy"),
//   	input: sourceOutput,
//   	environmentName: jsii.String("envName"),
//   	applicationName: jsii.String("appName"),
//   })
//
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type ElasticBeanstalkDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The name of the AWS Elastic Beanstalk application to deploy.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The name of the AWS Elastic Beanstalk environment to deploy to.
	EnvironmentName *string `field:"required" json:"environmentName" yaml:"environmentName"`
	// The source to use as input for deployment.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
}

