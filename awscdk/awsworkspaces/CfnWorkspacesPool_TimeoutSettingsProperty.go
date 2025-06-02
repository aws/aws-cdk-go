package awsworkspaces


// Describes the timeout settings for the pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutSettingsProperty := &TimeoutSettingsProperty{
//   	DisconnectTimeoutInSeconds: jsii.Number(123),
//   	IdleDisconnectTimeoutInSeconds: jsii.Number(123),
//   	MaxUserDurationInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-timeoutsettings.html
//
type CfnWorkspacesPool_TimeoutSettingsProperty struct {
	// Specifies the amount of time, in seconds, that a streaming session remains active after users disconnect.
	//
	// If users try to reconnect to the streaming session after a disconnection or network interruption within the time set, they are connected to their previous session. Otherwise, they are connected to a new session with a new streaming instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-timeoutsettings.html#cfn-workspaces-workspacespool-timeoutsettings-disconnecttimeoutinseconds
	//
	DisconnectTimeoutInSeconds *float64 `field:"optional" json:"disconnectTimeoutInSeconds" yaml:"disconnectTimeoutInSeconds"`
	// The amount of time in seconds a connection will stay active while idle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-timeoutsettings.html#cfn-workspaces-workspacespool-timeoutsettings-idledisconnecttimeoutinseconds
	//
	IdleDisconnectTimeoutInSeconds *float64 `field:"optional" json:"idleDisconnectTimeoutInSeconds" yaml:"idleDisconnectTimeoutInSeconds"`
	// Specifies the maximum amount of time, in seconds, that a streaming session can remain active.
	//
	// If users are still connected to a streaming instance five minutes before this limit is reached, they are prompted to save any open documents before being disconnected. After this time elapses, the instance is terminated and replaced by a new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspacespool-timeoutsettings.html#cfn-workspaces-workspacespool-timeoutsettings-maxuserdurationinseconds
	//
	MaxUserDurationInSeconds *float64 `field:"optional" json:"maxUserDurationInSeconds" yaml:"maxUserDurationInSeconds"`
}

