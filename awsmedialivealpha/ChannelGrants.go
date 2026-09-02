package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Collection of grant methods for a IChannelRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var channelRef IChannelRef
//
//   channelGrants := medialive_alpha.ChannelGrants_FromChannel(channelRef)
//
// Experimental.
type ChannelGrants interface {
	// Experimental.
	Resource() interfacesawsmedialive.IChannelRef
	// Grant the given identity custom permissions.
	// Experimental.
	Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant
	// Grant permissions to start this channel.
	// Experimental.
	Start(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to stop this channel.
	// Experimental.
	Stop(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to update the schedule of this channel.
	// Experimental.
	UpdateSchedule(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ChannelGrants
type jsiiProxy_ChannelGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ChannelGrants) Resource() interfacesawsmedialive.IChannelRef {
	var returns interfacesawsmedialive.IChannelRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ChannelGrants.
// Experimental.
func ChannelGrants_FromChannel(resource interfacesawsmedialive.IChannelRef) ChannelGrants {
	_init_.Initialize()

	if err := validateChannelGrants_FromChannelParameters(resource); err != nil {
		panic(err)
	}
	var returns ChannelGrants

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ChannelGrants",
		"fromChannel",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChannelGrants) Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant {
	if err := c.validateActionsParameters(grantee, actions, options); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"actions",
		[]interface{}{grantee, actions, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChannelGrants) Start(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateStartParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"start",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChannelGrants) Stop(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateStopParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"stop",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChannelGrants) UpdateSchedule(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateUpdateScheduleParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"updateSchedule",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

