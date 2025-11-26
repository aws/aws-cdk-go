package previewawsworkspacesevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Workspace aws.workspaces@WorkSpacesAccess event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workSpacesAccessProps := &WorkSpacesAccessProps{
//   	ActionType: []*string{
//   		jsii.String("actionType"),
//   	},
//   	ClientIpAddress: []*string{
//   		jsii.String("clientIpAddress"),
//   	},
//   	ClientPlatform: []*string{
//   		jsii.String("clientPlatform"),
//   	},
//   	DirectoryId: []*string{
//   		jsii.String("directoryId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	LoginTime: []*string{
//   		jsii.String("loginTime"),
//   	},
//   	WorkspaceId: []*string{
//   		jsii.String("workspaceId"),
//   	},
//   	WorkspacesClientProductName: []*string{
//   		jsii.String("workspacesClientProductName"),
//   	},
//   }
//
// Experimental.
type WorkspaceEvents_WorkSpacesAccess_WorkSpacesAccessProps struct {
	// actionType property.
	//
	// Specify an array of string values to match this event if the actual value of actionType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionType *[]*string `field:"optional" json:"actionType" yaml:"actionType"`
	// clientIpAddress property.
	//
	// Specify an array of string values to match this event if the actual value of clientIpAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientIpAddress *[]*string `field:"optional" json:"clientIpAddress" yaml:"clientIpAddress"`
	// clientPlatform property.
	//
	// Specify an array of string values to match this event if the actual value of clientPlatform is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientPlatform *[]*string `field:"optional" json:"clientPlatform" yaml:"clientPlatform"`
	// directoryId property.
	//
	// Specify an array of string values to match this event if the actual value of directoryId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DirectoryId *[]*string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// loginTime property.
	//
	// Specify an array of string values to match this event if the actual value of loginTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LoginTime *[]*string `field:"optional" json:"loginTime" yaml:"loginTime"`
	// workspaceId property.
	//
	// Specify an array of string values to match this event if the actual value of workspaceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Workspace reference.
	//
	// Experimental.
	WorkspaceId *[]*string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
	// workspacesClientProductName property.
	//
	// Specify an array of string values to match this event if the actual value of workspacesClientProductName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WorkspacesClientProductName *[]*string `field:"optional" json:"workspacesClientProductName" yaml:"workspacesClientProductName"`
}

