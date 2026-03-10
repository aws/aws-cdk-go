package previewawsecrevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr"
)

// EventBridge event patterns for Repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repositoryRef IRepositoryRef
//
//   repositoryEvents := awscdkmixinspreview.Events.RepositoryEvents_FromRepository(repositoryRef)
//
// Experimental.
type RepositoryEvents interface {
	// EventBridge event pattern for Repository AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository ECR Image Action.
	// Experimental.
	EcrImageActionPattern(options *ECRImageAction_ECRImageActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository ECR Image Scan.
	// Experimental.
	EcrImageScanPattern(options *ECRImageScan_ECRImageScanProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository ECR Pull Through Cache Action.
	// Experimental.
	EcrPullThroughCacheActionPattern(options *ECRPullThroughCacheAction_ECRPullThroughCacheActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository ECR Referrer Action.
	// Experimental.
	EcrReferrerActionPattern(options *ECRReferrerAction_ECRReferrerActionProps) *awsevents.EventPattern
	// EventBridge event pattern for Repository ECR Replication Action.
	// Experimental.
	EcrReplicationActionPattern(options *ECRReplicationAction_ECRReplicationActionProps) *awsevents.EventPattern
}

// The jsii proxy struct for RepositoryEvents
type jsiiProxy_RepositoryEvents struct {
	_ byte // padding
}

// Create RepositoryEvents from a Repository reference.
// Experimental.
func RepositoryEvents_FromRepository(repositoryRef interfacesawsecr.IRepositoryRef) RepositoryEvents {
	_init_.Initialize()

	if err := validateRepositoryEvents_FromRepositoryParameters(repositoryRef); err != nil {
		panic(err)
	}
	var returns RepositoryEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents",
		"fromRepository",
		[]interface{}{repositoryRef},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) AwsAPICallViaCloudTrailPattern(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := r.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) EcrImageActionPattern(options *ECRImageAction_ECRImageActionProps) *awsevents.EventPattern {
	if err := r.validateEcrImageActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"ecrImageActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) EcrImageScanPattern(options *ECRImageScan_ECRImageScanProps) *awsevents.EventPattern {
	if err := r.validateEcrImageScanPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"ecrImageScanPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) EcrPullThroughCacheActionPattern(options *ECRPullThroughCacheAction_ECRPullThroughCacheActionProps) *awsevents.EventPattern {
	if err := r.validateEcrPullThroughCacheActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"ecrPullThroughCacheActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) EcrReferrerActionPattern(options *ECRReferrerAction_ECRReferrerActionProps) *awsevents.EventPattern {
	if err := r.validateEcrReferrerActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"ecrReferrerActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEvents) EcrReplicationActionPattern(options *ECRReplicationAction_ECRReplicationActionProps) *awsevents.EventPattern {
	if err := r.validateEcrReplicationActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"ecrReplicationActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

