package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Browser resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var browserNetworkConfiguration BrowserNetworkConfiguration
//   var role Role
//
//   browserCustomProps := &BrowserCustomProps{
//   	BrowserCustomName: jsii.String("browserCustomName"),
//   	BrowserSigning: bedrock_agentcore_alpha.BrowserSigning_ENABLED,
//   	Description: jsii.String("description"),
//   	ExecutionRole: role,
//   	NetworkConfiguration: browserNetworkConfiguration,
//   	RecordingConfig: &RecordingConfig{
//   		Enabled: jsii.Boolean(false),
//   		S3Location: &Location{
//   			BucketName: jsii.String("bucketName"),
//   			ObjectKey: jsii.String("objectKey"),
//
//   			// the properties below are optional
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type BrowserCustomProps struct {
	// The name of the browser Valid characters are a-z, A-Z, 0-9, _ (underscore) The name must start with a letter and can be up to 48 characters long Pattern: [a-zA-Z][a-zA-Z0-9_]{0,47}.
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	BrowserCustomName *string `field:"optional" json:"browserCustomName" yaml:"browserCustomName"`
	// Specifies whether browser signing is enabled.
	//
	// When enabled, the browser will cryptographically sign
	// HTTP requests to identify itself as an AI agent to bot control vendors.
	// Default: - BrowserSigning.DISABLED
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	BrowserSigning BrowserSigning `field:"optional" json:"browserSigning" yaml:"browserSigning"`
	// Optional description for the browser Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the browser to access AWS services.
	// Default: - A new role will be created.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// Network configuration for browser.
	// Default: - PUBLIC network mode.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	NetworkConfiguration BrowserNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Recording configuration for browser.
	// Default: - No recording configuration.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	RecordingConfig *RecordingConfig `field:"optional" json:"recordingConfig" yaml:"recordingConfig"`
	// Tags (optional) A list of key:value pairs of tags to apply to this Browser resource.
	// Default: {} - no tags.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

