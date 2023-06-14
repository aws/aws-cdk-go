package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an ECR repository.
type IRepository interface {
	awscdk.IResource
	// Add a policy statement to the repository's resource policy.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grant the given principal identity permissions to perform the actions on this repository.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to pull images in this repository.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push images to this repository.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to push images in this repository.
	GrantPush(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to read images in this repository.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when the image scan is completed.
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	// Returns the URI of the repository for a certain digest. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	RepositoryUriForDigest(digest *string) *string
	// Returns the URI of the repository for a certain tag. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	RepositoryUriForTag(tag *string) *string
	// Returns the URI of the repository for a certain tag or digest, inferring based on the syntax of the tag.
	//
	// Can be used in `docker push/pull`.
	//
	//    ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	//    ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	RepositoryUriForTagOrDigest(tagOrDigest *string) *string
	// The ARN of the repository.
	RepositoryArn() *string
	// The name of the repository.
	RepositoryName() *string
	// The URI of this repository (represents the latest image):.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY
	RepositoryUri() *string
}

// The jsii proxy for IRepository
type jsiiProxy_IRepository struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRepository) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPull(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPullParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPullPushParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantPush(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnCloudTrailEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule {
	if err := i.validateOnCloudTrailImagePushedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailImagePushed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule {
	if err := i.validateOnImageScanCompletedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onImageScanCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) RepositoryUriForDigest(digest *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"repositoryUriForDigest",
		[]interface{}{digest},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) RepositoryUriForTag(tag *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"repositoryUriForTag",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRepository) RepositoryUriForTagOrDigest(tagOrDigest *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"repositoryUriForTagOrDigest",
		[]interface{}{tagOrDigest},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRepository) RepositoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepository) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

