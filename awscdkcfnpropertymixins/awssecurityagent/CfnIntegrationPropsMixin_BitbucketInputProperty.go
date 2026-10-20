package awssecurityagent


// The input required to create a Bitbucket integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   bitbucketInputProperty := &BitbucketInputProperty{
//   	Code: jsii.String("code"),
//   	InstallationId: jsii.String("installationId"),
//   	State: jsii.String("state"),
//   	Workspace: jsii.String("workspace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-bitbucketinput.html
//
type CfnIntegrationPropsMixin_BitbucketInputProperty struct {
	// The OAuth 2.0 authorization code returned from the consent redirect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-bitbucketinput.html#cfn-securityagent-integration-bitbucketinput-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The Atlassian installation identifier, available from the Atlassian administration console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-bitbucketinput.html#cfn-securityagent-integration-bitbucketinput-installationid
	//
	InstallationId *string `field:"optional" json:"installationId" yaml:"installationId"`
	// The CSRF state token echoed back from the OAuth redirect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-bitbucketinput.html#cfn-securityagent-integration-bitbucketinput-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The Bitbucket workspace slug that identifies the workspace to integrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-integration-bitbucketinput.html#cfn-securityagent-integration-bitbucketinput-workspace
	//
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

