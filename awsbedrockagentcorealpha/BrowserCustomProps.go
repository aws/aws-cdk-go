package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Browser resource.
//
// Example:
//   // Create a custom execution role
//   executionRole := iam.NewRole(this, jsii.String("BrowserExecutionRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
//   	ManagedPolicies: []IManagedPolicy{
//   		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonBedrockAgentCoreBrowserExecutionRolePolicy")),
//   	},
//   })
//
//   // Create browser with custom execution role
//   browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
//   	BrowserCustomName: jsii.String("my_browser"),
//   	Description: jsii.String("Browser with custom execution role"),
//   	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
//   	ExecutionRole: executionRole,
//   })
//
// Experimental.
type BrowserCustomProps struct {
	// The name of the browser Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Experimental.
	BrowserCustomName *string `field:"required" json:"browserCustomName" yaml:"browserCustomName"`
	// Optional description for the browser Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the browser to access AWS services.
	// Default: - A new role will be created.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Network configuration for browser.
	// Default: - PUBLIC network mode.
	//
	// Experimental.
	NetworkConfiguration BrowserNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Recording configuration for browser.
	// Default: - No recording configuration.
	//
	// Experimental.
	RecordingConfig *RecordingConfig `field:"optional" json:"recordingConfig" yaml:"recordingConfig"`
	// Tags (optional) A list of key:value pairs of tags to apply to this Browser resource.
	// Default: {} - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

