package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// An EBS Volume in AWS EC2.
type IVolume interface {
	awscdk.IResource
	// Grants permission to attach this Volume to an instance.
	//
	// CAUTION: Granting an instance permission to attach to itself using this method will lead to
	// an unresolvable circular reference between the instance role and the instance.
	// Use `IVolume.grantAttachVolumeToSelf` to grant an instance permission to attach this
	// volume to itself.
	GrantAttachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant
	// Grants permission to attach the Volume by a ResourceTag condition.
	//
	// If you are looking to
	// grant an Instance, AutoScalingGroup, EC2-Fleet, SpotFleet, ECS host, etc the ability to attach
	// this volume to **itself** then this is the method you want to use.
	//
	// This is implemented by adding a Tag with key `VolumeGrantAttach-<suffix>` to the given
	// constructs and this Volume, and then conditioning the Grant such that the grantee is only
	// given the ability to AttachVolume if both the Volume and the destination Instance have that
	// tag applied to them.
	GrantAttachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant
	// Grants permission to detach this Volume from an instance CAUTION: Granting an instance permission to detach from itself using this method will lead to an unresolvable circular reference between the instance role and the instance.
	//
	// Use `IVolume.grantDetachVolumeFromSelf` to grant an instance permission to detach this
	// volume from itself.
	GrantDetachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant
	// Grants permission to detach the Volume by a ResourceTag condition.
	//
	// This is implemented via the same mechanism as `IVolume.grantAttachVolumeByResourceTag`,
	// and is subject to the same conditions.
	GrantDetachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant
	// The availability zone that the EBS Volume is contained within (ex: us-west-2a).
	AvailabilityZone() *string
	// The customer-managed encryption key that is used to encrypt the Volume.
	EncryptionKey() awskms.IKey
	// The EBS Volume's ID.
	VolumeId() *string
}

// The jsii proxy for IVolume
type jsiiProxy_IVolume struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVolume) GrantAttachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant {
	if err := i.validateGrantAttachVolumeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantAttachVolume",
		[]interface{}{grantee, instances},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVolume) GrantAttachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant {
	if err := i.validateGrantAttachVolumeByResourceTagParameters(grantee, constructs); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantAttachVolumeByResourceTag",
		[]interface{}{grantee, constructs, tagKeySuffix},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVolume) GrantDetachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant {
	if err := i.validateGrantDetachVolumeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDetachVolume",
		[]interface{}{grantee, instances},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVolume) GrantDetachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant {
	if err := i.validateGrantDetachVolumeByResourceTagParameters(grantee, constructs); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDetachVolumeByResourceTag",
		[]interface{}{grantee, constructs, tagKeySuffix},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVolume) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolume) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

