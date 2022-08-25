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
//   actionDeclarationProperty := &actionDeclarationProperty{
//   	actionTypeId: &actionTypeIdProperty{
//   		category: jsii.String("category"),
//   		owner: jsii.String("owner"),
//   		provider: jsii.String("provider"),
//   		version: jsii.String("version"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	configuration: configuration,
//   	inputArtifacts: []interface{}{
//   		&inputArtifactProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	namespace: jsii.String("namespace"),
//   	outputArtifacts: []interface{}{
//   		&outputArtifactProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	region: jsii.String("region"),
//   	roleArn: jsii.String("roleArn"),
//   	runOrder: jsii.Number(123),
//   }
//
type CfnPipeline_ActionDeclarationProperty struct {
	// Specifies the action type and the provider of the action.
	ActionTypeId interface{} `field:"required" json:"actionTypeId" yaml:"actionTypeId"`
	// The action declaration's name.
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
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The name or ID of the artifact consumed by the action, such as a test or build artifact.
	//
	// > For a CodeBuild action with multiple input artifacts, one of your input sources must be designated the PrimarySource. For more information, see the [CodeBuild action reference page](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeBuild.html) in the *AWS CodePipeline User Guide* .
	InputArtifacts interface{} `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// The variable namespace associated with the action.
	//
	// All variables produced as output by this action fall under this namespace.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name or ID of the result of the action declaration, such as a test or build artifact.
	OutputArtifacts interface{} `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// The action declaration's AWS Region, such as us-east-1.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the IAM service role that performs the declared action.
	//
	// This is assumed through the roleArn for the pipeline.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The order in which actions are run.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
}

