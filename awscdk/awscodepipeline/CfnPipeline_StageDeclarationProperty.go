package awscodepipeline


// Represents information about a stage and its definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   stageDeclarationProperty := &StageDeclarationProperty{
//   	Actions: []interface{}{
//   		&ActionDeclarationProperty{
//   			ActionTypeId: &ActionTypeIdProperty{
//   				Category: jsii.String("category"),
//   				Owner: jsii.String("owner"),
//   				Provider: jsii.String("provider"),
//   				Version: jsii.String("version"),
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Commands: []*string{
//   				jsii.String("commands"),
//   			},
//   			Configuration: configuration,
//   			EnvironmentVariables: []interface{}{
//   				&EnvironmentVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			InputArtifacts: []interface{}{
//   				&InputArtifactProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Namespace: jsii.String("namespace"),
//   			OutputArtifacts: []interface{}{
//   				&OutputArtifactProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Files: []*string{
//   						jsii.String("files"),
//   					},
//   				},
//   			},
//   			OutputVariables: []*string{
//   				jsii.String("outputVariables"),
//   			},
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			RunOrder: jsii.Number(123),
//   			TimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	BeforeEntry: &BeforeEntryConditionsProperty{
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				Result: jsii.String("result"),
//   				Rules: []interface{}{
//   					&RuleDeclarationProperty{
//   						Commands: []*string{
//   							jsii.String("commands"),
//   						},
//   						Configuration: configuration,
//   						InputArtifacts: []interface{}{
//   							&InputArtifactProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   						Region: jsii.String("region"),
//   						RoleArn: jsii.String("roleArn"),
//   						RuleTypeId: &RuleTypeIdProperty{
//   							Category: jsii.String("category"),
//   							Owner: jsii.String("owner"),
//   							Provider: jsii.String("provider"),
//   							Version: jsii.String("version"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Blockers: []interface{}{
//   		&BlockerDeclarationProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	OnFailure: &FailureConditionsProperty{
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				Result: jsii.String("result"),
//   				Rules: []interface{}{
//   					&RuleDeclarationProperty{
//   						Commands: []*string{
//   							jsii.String("commands"),
//   						},
//   						Configuration: configuration,
//   						InputArtifacts: []interface{}{
//   							&InputArtifactProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   						Region: jsii.String("region"),
//   						RoleArn: jsii.String("roleArn"),
//   						RuleTypeId: &RuleTypeIdProperty{
//   							Category: jsii.String("category"),
//   							Owner: jsii.String("owner"),
//   							Provider: jsii.String("provider"),
//   							Version: jsii.String("version"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Result: jsii.String("result"),
//   		RetryConfiguration: &RetryConfigurationProperty{
//   			RetryMode: jsii.String("retryMode"),
//   		},
//   	},
//   	OnSuccess: &SuccessConditionsProperty{
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				Result: jsii.String("result"),
//   				Rules: []interface{}{
//   					&RuleDeclarationProperty{
//   						Commands: []*string{
//   							jsii.String("commands"),
//   						},
//   						Configuration: configuration,
//   						InputArtifacts: []interface{}{
//   							&InputArtifactProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   						Name: jsii.String("name"),
//   						Region: jsii.String("region"),
//   						RoleArn: jsii.String("roleArn"),
//   						RuleTypeId: &RuleTypeIdProperty{
//   							Category: jsii.String("category"),
//   							Owner: jsii.String("owner"),
//   							Provider: jsii.String("provider"),
//   							Version: jsii.String("version"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html
//
type CfnPipeline_StageDeclarationProperty struct {
	// The actions included in a stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The name of the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The method to use when a stage allows entry.
	//
	// For example, configuring this field for conditions will allow entry to the stage when the conditions are met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-beforeentry
	//
	BeforeEntry interface{} `field:"optional" json:"beforeEntry" yaml:"beforeEntry"`
	// Reserved for future use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-blockers
	//
	Blockers interface{} `field:"optional" json:"blockers" yaml:"blockers"`
	// The method to use when a stage has not completed successfully.
	//
	// For example, configuring this field for rollback will roll back a failed stage automatically to the last successful pipeline execution in the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-onfailure
	//
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The method to use when a stage has succeeded.
	//
	// For example, configuring this field for conditions will allow the stage to succeed when the conditions are met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stagedeclaration.html#cfn-codepipeline-pipeline-stagedeclaration-onsuccess
	//
	OnSuccess interface{} `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

