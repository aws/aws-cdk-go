package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkspacesPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspacesPoolProps := &CfnWorkspacesPoolProps{
//   	BundleId: jsii.String("bundleId"),
//   	Capacity: &CapacityProperty{
//   		DesiredUserSessions: jsii.Number(123),
//   	},
//   	DirectoryId: jsii.String("directoryId"),
//   	PoolName: jsii.String("poolName"),
//
//   	// the properties below are optional
//   	ApplicationSettings: &ApplicationSettingsProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		SettingsGroup: jsii.String("settingsGroup"),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutSettings: &TimeoutSettingsProperty{
//   		DisconnectTimeoutInSeconds: jsii.Number(123),
//   		IdleDisconnectTimeoutInSeconds: jsii.Number(123),
//   		MaxUserDurationInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html
//
type CfnWorkspacesPoolProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-bundleid
	//
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-capacity
	//
	Capacity interface{} `field:"required" json:"capacity" yaml:"capacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-directoryid
	//
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-poolname
	//
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-applicationsettings
	//
	ApplicationSettings interface{} `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-timeoutsettings
	//
	TimeoutSettings interface{} `field:"optional" json:"timeoutSettings" yaml:"timeoutSettings"`
}

