package awscodepipeline


// Represents information about an action type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionTypeIdProperty := &ActionTypeIdProperty{
//   	Category: jsii.String("category"),
//   	Owner: jsii.String("owner"),
//   	Provider: jsii.String("provider"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html
//
type CfnPipeline_ActionTypeIdProperty struct {
	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action.
	//
	// Valid categories are limited to one of the values below.
	//
	// - `Source`
	// - `Build`
	// - `Test`
	// - `Deploy`
	// - `Invoke`
	// - `Approval`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-category
	//
	Category *string `field:"required" json:"category" yaml:"category"`
	// The creator of the action being called.
	//
	// There are three valid values for the `Owner` field in the action category section within your pipeline structure: `AWS` , `ThirdParty` , and `Custom` . For more information, see [Valid Action Types and Providers in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#actions-valid-providers) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-owner
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The provider of the service being called by the action.
	//
	// Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of CodeDeploy, which would be specified as `CodeDeploy` . For more information, see [Valid Action Types and Providers in CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#actions-valid-providers) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-provider
	//
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// A string that describes the action version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-actiontypeid.html#cfn-codepipeline-pipeline-actiontypeid-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
}

