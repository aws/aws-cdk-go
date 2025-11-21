package awssqs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssqs"
)

// Collection of grant methods for a IQueueRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queueRef IQueueRef
//
//   queueGrants := awscdk.Aws_sqs.QueueGrants_FromQueue(queueRef)
//
type QueueGrants interface {
	EncryptedResource() awsiam.IEncryptedResource
	PolicyResource() awsiam.IResourceWithPolicyV2
	Resource() interfacesawssqs.IQueueRef
	// Grant permissions to consume messages from a queue.
	//
	// This will grant the following permissions:
	//
	//   - sqs:ChangeMessageVisibility
	//   - sqs:DeleteMessage
	//   - sqs:ReceiveMessage
	//   - sqs:GetQueueAttributes
	//   - sqs:GetQueueUrl
	//
	// If encryption is used, permission to use the key to decrypt the contents of the queue will also be granted to the same principal.
	//
	// This will grant the following KMS permissions:
	//
	// - kms:Decrypt.
	ConsumeMessages(grantee awsiam.IGrantable) awsiam.Grant
	// Grants purge permissions.
	Purge(grantee awsiam.IGrantable) awsiam.Grant
	// Grant access to send messages to a queue to the given identity.
	//
	// This will grant the following permissions:
	//
	//  - sqs:SendMessage
	//  - sqs:GetQueueAttributes
	//  - sqs:GetQueueUrl
	//
	// If encryption is used, permission to use the key to encrypt/decrypt the contents of the queue will also be granted to the same principal.
	//
	// This will grant the following KMS permissions:
	//
	//  - kms:Decrypt
	//  - kms:Encrypt
	//  - kms:ReEncrypt*
	// - kms:GenerateDataKey*.
	SendMessages(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for QueueGrants
type jsiiProxy_QueueGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_QueueGrants) EncryptedResource() awsiam.IEncryptedResource {
	var returns awsiam.IEncryptedResource
	_jsii_.Get(
		j,
		"encryptedResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueGrants) PolicyResource() awsiam.IResourceWithPolicyV2 {
	var returns awsiam.IResourceWithPolicyV2
	_jsii_.Get(
		j,
		"policyResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueueGrants) Resource() interfacesawssqs.IQueueRef {
	var returns interfacesawssqs.IQueueRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for QueueGrants.
func QueueGrants_FromQueue(resource interfacesawssqs.IQueueRef) QueueGrants {
	_init_.Initialize()

	if err := validateQueueGrants_FromQueueParameters(resource); err != nil {
		panic(err)
	}
	var returns QueueGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.QueueGrants",
		"fromQueue",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueGrants) ConsumeMessages(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateConsumeMessagesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"consumeMessages",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueGrants) Purge(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validatePurgeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"purge",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueGrants) SendMessages(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateSendMessagesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"sendMessages",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

