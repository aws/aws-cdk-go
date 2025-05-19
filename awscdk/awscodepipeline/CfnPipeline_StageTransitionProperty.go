package awscodepipeline


// The name of the pipeline in which you want to disable the flow of artifacts from one stage to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageTransitionProperty := &StageTransitionProperty{
//   	Reason: jsii.String("reason"),
//   	StageName: jsii.String("stageName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagetransition.html
//
type CfnPipeline_StageTransitionProperty struct {
	// The reason given to the user that a stage is disabled, such as waiting for manual approval or manual tests.
	//
	// This message is displayed in the pipeline console UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagetransition.html#cfn-codepipeline-pipeline-stagetransition-reason
	//
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// The name of the stage where you want to disable the inbound or outbound transition of artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagetransition.html#cfn-codepipeline-pipeline-stagetransition-stagename
	//
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

