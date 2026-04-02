package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2"
)

// Collection of grant methods for a IChannelRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var channelRef IChannelRef
//
//   channelGrants := mediapackagev2_alpha.ChannelGrants_FromChannel(channelRef)
//
// Experimental.
type ChannelGrants interface {
	// Experimental.
	PolicyResource() awsiam.IResourceWithPolicyV2
	// Experimental.
	Resource() interfacesawsmediapackagev2.IChannelRef
	// Grant the given identity custom permissions.
	// Experimental.
	Actions(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) awsiam.Grant
	// Grant permissions to ingest content into this channel.
	// Experimental.
	Ingest(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ChannelGrants
type jsiiProxy_ChannelGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ChannelGrants) PolicyResource() awsiam.IResourceWithPolicyV2 {
	var returns awsiam.IResourceWithPolicyV2
	_jsii_.Get(
		j,
		"policyResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChannelGrants) Resource() interfacesawsmediapackagev2.IChannelRef {
	var returns interfacesawsmediapackagev2.IChannelRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ChannelGrants.
// Experimental.
func ChannelGrants_FromChannel(resource interfacesawsmediapackagev2.IChannelRef) ChannelGrants {
	_init_.Initialize()

	if err := validateChannelGrants_FromChannelParameters(resource); err != nil {
		panic(err)
	}
	var returns ChannelGrants

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelGrants",
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

func (c *jsiiProxy_ChannelGrants) Ingest(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateIngestParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"ingest",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

