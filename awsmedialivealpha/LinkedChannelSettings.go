package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Linked channel settings for primary/follower channel configurations.
//
// Use the static factory methods to create.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var channel Channel
//
//   linkedChannelSettings := medialive_alpha.LinkedChannelSettings_Follower(channel)
//
// Experimental.
type LinkedChannelSettings interface {
}

// The jsii proxy struct for LinkedChannelSettings
type jsiiProxy_LinkedChannelSettings struct {
	_ byte // padding
}

// Experimental.
func NewLinkedChannelSettings_Override(l LinkedChannelSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.LinkedChannelSettings",
		nil, // no parameters
		l,
	)
}

// Configure this channel as a follower of a primary channel.
// Experimental.
func LinkedChannelSettings_Follower(primaryChannel IChannel) LinkedChannelSettings {
	_init_.Initialize()

	if err := validateLinkedChannelSettings_FollowerParameters(primaryChannel); err != nil {
		panic(err)
	}
	var returns LinkedChannelSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.LinkedChannelSettings",
		"follower",
		[]interface{}{primaryChannel},
		&returns,
	)

	return returns
}

// Configure this channel as a primary in a linked channel pair.
// Experimental.
func LinkedChannelSettings_Primary() LinkedChannelSettings {
	_init_.Initialize()

	var returns LinkedChannelSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.LinkedChannelSettings",
		"primary",
		nil, // no parameters
		&returns,
	)

	return returns
}

