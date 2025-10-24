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
//   	RunningMode: jsii.String("runningMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
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
	// The identifier of the bundle used by the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-bundleid
	//
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// Describes the user capacity for the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-capacity
	//
	Capacity interface{} `field:"required" json:"capacity" yaml:"capacity"`
	// The identifier of the directory used by the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-directoryid
	//
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// The name of the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-poolname
	//
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// The persistent application settings for users of the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-applicationsettings
	//
	ApplicationSettings interface{} `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// The description of the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The running mode of the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-runningmode
	//
	RunningMode *string `field:"optional" json:"runningMode" yaml:"runningMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-tags
	//
	// Deprecated: this property has been deprecated.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The amount of time that a pool session remains active after users disconnect.
	//
	// If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspacespool.html#cfn-workspaces-workspacespool-timeoutsettings
	//
	TimeoutSettings interface{} `field:"optional" json:"timeoutSettings" yaml:"timeoutSettings"`
}

