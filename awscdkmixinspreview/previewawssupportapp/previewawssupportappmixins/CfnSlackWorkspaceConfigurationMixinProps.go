package previewawssupportappmixins


// Properties for CfnSlackWorkspaceConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSlackWorkspaceConfigurationMixinProps := &CfnSlackWorkspaceConfigurationMixinProps{
//   	TeamId: jsii.String("teamId"),
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html
//
type CfnSlackWorkspaceConfigurationMixinProps struct {
	// The team ID in Slack.
	//
	// This ID uniquely identifies a Slack workspace, such as `T012ABCDEFG` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html#cfn-supportapp-slackworkspaceconfiguration-teamid
	//
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// An identifier used to update an existing Slack workspace configuration in AWS CloudFormation , such as `100` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html#cfn-supportapp-slackworkspaceconfiguration-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

