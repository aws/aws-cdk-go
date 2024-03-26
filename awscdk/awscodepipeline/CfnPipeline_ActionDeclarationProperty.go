package awscodepipeline


// Represents information about an action declaration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   actionDeclarationProperty := &ActionDeclarationProperty{
//   	ActionTypeId: &ActionTypeIdProperty{
//   		Category: jsii.String("category"),
//   		Owner: jsii.String("owner"),
//   		Provider: jsii.String("provider"),
//   		Version: jsii.String("version"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Configuration: configuration,
//   	InputArtifacts: []interface{}{
//   		&InputArtifactProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	OutputArtifacts: []interface{}{
//   		&OutputArtifactProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Region: jsii.String("region"),
//   	RoleArn: jsii.String("roleArn"),
//   	RunOrder: jsii.Number(123),
//   	TimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html
//
type CfnPipeline_ActionDeclarationProperty struct {
	// Specifies the action type and the provider of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-actiontypeid
	//
	ActionTypeId interface{} `field:"required" json:"actionTypeId" yaml:"actionTypeId"`
	// The action declaration's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action's configuration.
	//
	// These are key-value pairs that specify input values for an action. For more information, see [Action Structure Requirements in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) . For the list of configuration properties for the AWS CloudFormation action type in CodePipeline, see [Configuration Properties Reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-action-reference.html) in the *AWS CloudFormation User Guide* . For template snippets with examples, see [Using Parameter Override Functions with CodePipeline Pipelines](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-parameter-override-functions.html) in the *AWS CloudFormation User Guide* .
	//
	// The values can be represented in either JSON or YAML format. For example, the JSON configuration item format is as follows:
	//
	// *JSON:*
	//
	// `"Configuration" : { Key : Value },`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The name or ID of the artifact consumed by the action, such as a test or build artifact.
	//
	// While the field is not a required parameter, most actions have an action configuration that requires a specified quantity of input artifacts. To refer to the action configuration specification by action provider, see the [Action structure reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) in the *AWS CodePipeline User Guide* .
	//
	// > For a CodeBuild action with multiple input artifacts, one of your input sources must be designated the PrimarySource. For more information, see the [CodeBuild action reference page](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeBuild.html) in the *AWS CodePipeline User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-inputartifacts
	//
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The variable namespace associated with the action.
	//
	// All variables produced as output by this action fall under this namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name or ID of the result of the action declaration, such as a test or build artifact.
	//
	// While the field is not a required parameter, most actions have an action configuration that requires a specified quantity of output artifacts. To refer to the action configuration specification by action provider, see the [Action structure reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) in the *AWS CodePipeline User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-outputartifacts
	//
	OutputArtifacts interface{} `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// The action declaration's AWS Region, such as us-east-1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM service role that performs the declared action.
	//
	// This is assumed through the roleArn for the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The order in which actions are run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-runorder
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// A timeout duration in minutes that can be applied against the ActionType’s default timeout value specified in [Quotas for AWS CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/limits.html) . This attribute is available only to the manual approval ActionType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiondeclaration.html#cfn-codepipeline-pipeline-actiondeclaration-timeoutinminutes
	//
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

