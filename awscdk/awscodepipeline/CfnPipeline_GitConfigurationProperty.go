package awscodepipeline


// A type of trigger configuration for Git-based source actions.
//
// > You can specify the Git configuration trigger type for all third-party Git-based source actions that are supported by the `CodeStarSourceConnection` action type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitConfigurationProperty := &GitConfigurationProperty{
//   	SourceActionName: jsii.String("sourceActionName"),
//
//   	// the properties below are optional
//   	Push: []interface{}{
//   		&GitPushFilterProperty{
//   			Tags: &GitTagFilterCriteriaProperty{
//   				Excludes: []*string{
//   					jsii.String("excludes"),
//   				},
//   				Includes: []*string{
//   					jsii.String("includes"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitconfiguration.html
//
type CfnPipeline_GitConfigurationProperty struct {
	// The name of the pipeline source action where the trigger configuration, such as Git tags, is specified.
	//
	// The trigger configuration will start the pipeline upon the specified change only.
	//
	// > You can only specify one trigger configuration per source action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitconfiguration.html#cfn-codepipeline-pipeline-gitconfiguration-sourceactionname
	//
	SourceActionName *string `field:"required" json:"sourceActionName" yaml:"sourceActionName"`
	// The field where the repository event that will start the pipeline, such as pushing Git tags, is specified with details.
	//
	// > Git tags is the only supported event type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitconfiguration.html#cfn-codepipeline-pipeline-gitconfiguration-push
	//
	Push interface{} `field:"optional" json:"push" yaml:"push"`
}

