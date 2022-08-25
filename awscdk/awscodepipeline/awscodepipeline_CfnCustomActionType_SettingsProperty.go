package awscodepipeline


// `Settings` is a property of the `AWS::CodePipeline::CustomActionType` resource that provides URLs that users can access to view information about the CodePipeline custom action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   settingsProperty := &settingsProperty{
//   	entityUrlTemplate: jsii.String("entityUrlTemplate"),
//   	executionUrlTemplate: jsii.String("executionUrlTemplate"),
//   	revisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   	thirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   }
//
type CfnCustomActionType_SettingsProperty struct {
	// The URL returned to the CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for a CodeDeploy deployment group.
	//
	// This link is provided as part of the action display in the pipeline.
	EntityUrlTemplate *string `field:"optional" json:"entityUrlTemplate" yaml:"entityUrlTemplate"`
	// The URL returned to the CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for CodeDeploy.
	//
	// This link is shown on the pipeline view page in the CodePipeline console and provides a link to the execution entity of the external action.
	ExecutionUrlTemplate *string `field:"optional" json:"executionUrlTemplate" yaml:"executionUrlTemplate"`
	// The URL returned to the CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	RevisionUrlTemplate *string `field:"optional" json:"revisionUrlTemplate" yaml:"revisionUrlTemplate"`
	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	ThirdPartyConfigurationUrl *string `field:"optional" json:"thirdPartyConfigurationUrl" yaml:"thirdPartyConfigurationUrl"`
}

