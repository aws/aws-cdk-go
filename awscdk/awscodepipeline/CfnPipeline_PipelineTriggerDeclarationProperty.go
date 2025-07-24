package awscodepipeline


// Represents information about the specified trigger configuration, such as the filter criteria and the source stage for the action that contains the trigger.
//
// > This is only supported for the `CodeStarSourceConnection` action type. > When a trigger configuration is specified, default change detection for repository and branch commits is disabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineTriggerDeclarationProperty := &PipelineTriggerDeclarationProperty{
//   	ProviderType: jsii.String("providerType"),
//
//   	// the properties below are optional
//   	GitConfiguration: &GitConfigurationProperty{
//   		SourceActionName: jsii.String("sourceActionName"),
//
//   		// the properties below are optional
//   		PullRequest: []interface{}{
//   			&GitPullRequestFilterProperty{
//   				Branches: &GitBranchFilterCriteriaProperty{
//   					Excludes: []*string{
//   						jsii.String("excludes"),
//   					},
//   					Includes: []*string{
//   						jsii.String("includes"),
//   					},
//   				},
//   				Events: []*string{
//   					jsii.String("events"),
//   				},
//   				FilePaths: &GitFilePathFilterCriteriaProperty{
//   					Excludes: []*string{
//   						jsii.String("excludes"),
//   					},
//   					Includes: []*string{
//   						jsii.String("includes"),
//   					},
//   				},
//   			},
//   		},
//   		Push: []interface{}{
//   			&GitPushFilterProperty{
//   				Branches: &GitBranchFilterCriteriaProperty{
//   					Excludes: []*string{
//   						jsii.String("excludes"),
//   					},
//   					Includes: []*string{
//   						jsii.String("includes"),
//   					},
//   				},
//   				FilePaths: &GitFilePathFilterCriteriaProperty{
//   					Excludes: []*string{
//   						jsii.String("excludes"),
//   					},
//   					Includes: []*string{
//   						jsii.String("includes"),
//   					},
//   				},
//   				Tags: &GitTagFilterCriteriaProperty{
//   					Excludes: []*string{
//   						jsii.String("excludes"),
//   					},
//   					Includes: []*string{
//   						jsii.String("includes"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.html
//
type CfnPipeline_PipelineTriggerDeclarationProperty struct {
	// The source provider for the event, such as connections configured for a repository with Git tags, for the specified trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.html#cfn-codepipeline-pipeline-pipelinetriggerdeclaration-providertype
	//
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// Provides the filter criteria and the source stage for the repository event that starts the pipeline, such as Git tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.html#cfn-codepipeline-pipeline-pipelinetriggerdeclaration-gitconfiguration
	//
	GitConfiguration interface{} `field:"optional" json:"gitConfiguration" yaml:"gitConfiguration"`
}

