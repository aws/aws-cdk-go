package awssupportapp


// Properties for defining a `CfnSlackWorkspaceConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackWorkspaceConfigurationProps := &cfnSlackWorkspaceConfigurationProps{
//   	teamId: jsii.String("teamId"),
//
//   	// the properties below are optional
//   	versionId: jsii.String("versionId"),
//   }
//
type CfnSlackWorkspaceConfigurationProps struct {
	// `AWS::SupportApp::SlackWorkspaceConfiguration.TeamId`.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// `AWS::SupportApp::SlackWorkspaceConfiguration.VersionId`.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

