package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Construction properties for a CdkStage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var stage iStage
//   var stageHost iStageHost
//   var topic topic
//
//   cdkStageProps := &cdkStageProps{
//   	cloudAssemblyArtifact: artifact,
//   	host: stageHost,
//   	pipelineStage: stage,
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkStageProps struct {
	// The CodePipeline Artifact with the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `field:"required" json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// Features the Stage needs from its environment.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Host IStageHost `field:"required" json:"host" yaml:"host"`
	// The underlying Pipeline Stage associated with thisCdkStage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PipelineStage awscodepipeline.IStage `field:"required" json:"pipelineStage" yaml:"pipelineStage"`
	// Name of the stage that should be created.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// Run a security check before every application prepare/deploy actions.
	//
	// Note: Stage level security check can be overriden per application as follows:
	//    `stage.addApplication(app, { confirmBroadeningPermissions: false })`
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ConfirmBroadeningPermissions *bool `field:"optional" json:"confirmBroadeningPermissions" yaml:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when any security check registers changes within a application.
	//
	// Note: The Stage Notification Topic can be overriden per application as follows:
	//    `stage.addApplication(app, { securityNotificationTopic: newTopic })`
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityNotificationTopic awssns.ITopic `field:"optional" json:"securityNotificationTopic" yaml:"securityNotificationTopic"`
}

