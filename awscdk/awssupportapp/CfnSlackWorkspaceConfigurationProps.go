package awssupportapp


// Properties for defining a `CfnSlackWorkspaceConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackWorkspaceConfigurationProps := &CfnSlackWorkspaceConfigurationProps{
//   	TeamId: jsii.String("teamId"),
//
//   	// the properties below are optional
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html
//
type CfnSlackWorkspaceConfigurationProps struct {
	// The team ID in Slack.
	//
	// This ID uniquely identifies a Slack workspace, such as `T012ABCDEFG` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html#cfn-supportapp-slackworkspaceconfiguration-teamid
	//
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// An identifier used to update an existing Slack workspace configuration in AWS CloudFormation , such as `100` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackworkspaceconfiguration.html#cfn-supportapp-slackworkspaceconfiguration-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

