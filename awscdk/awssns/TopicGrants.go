package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssns"
)

// Collection of grant methods for a ITopicRef.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   // This would work the same way if topic was a CfnTopic (L1)
//   sns.TopicGrants_FromTopic(topic).Subscribe(iam.NewAnyPrincipal())
//
type TopicGrants interface {
	EncryptedResource() awsiam.IEncryptedResource
	PolicyResource() awsiam.IResourceWithPolicyV2
	Resource() interfacesawssns.ITopicRef
	// Grant topic publishing permissions to the given identity.
	Publish(grantee awsiam.IGrantable) awsiam.Grant
	// Grant topic subscribing permissions to the given identity.
	Subscribe(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for TopicGrants
type jsiiProxy_TopicGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_TopicGrants) EncryptedResource() awsiam.IEncryptedResource {
	var returns awsiam.IEncryptedResource
	_jsii_.Get(
		j,
		"encryptedResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicGrants) PolicyResource() awsiam.IResourceWithPolicyV2 {
	var returns awsiam.IResourceWithPolicyV2
	_jsii_.Get(
		j,
		"policyResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicGrants) Resource() interfacesawssns.ITopicRef {
	var returns interfacesawssns.ITopicRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for TopicGrants.
func TopicGrants_FromTopic(resource interfacesawssns.ITopicRef) TopicGrants {
	_init_.Initialize()

	if err := validateTopicGrants_FromTopicParameters(resource); err != nil {
		panic(err)
	}
	var returns TopicGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.TopicGrants",
		"fromTopic",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicGrants) Publish(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validatePublishParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"publish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicGrants) Subscribe(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validateSubscribeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"subscribe",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

