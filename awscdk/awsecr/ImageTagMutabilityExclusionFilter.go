package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an image tag mutability exclusion filter for ECR repository.
//
// Example:
//   // Make all tags immutable except for those starting with 'dev-' or 'test-'
//   // Make all tags immutable except for those starting with 'dev-' or 'test-'
//   ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	ImageTagMutability: ecr.TagMutability_IMMUTABLE_WITH_EXCLUSION,
//   	ImageTagMutabilityExclusionFilters: []ImageTagMutabilityExclusionFilter{
//   		ecr.ImageTagMutabilityExclusionFilter_Wildcard(jsii.String("dev-*")),
//   		ecr.ImageTagMutabilityExclusionFilter_*Wildcard(jsii.String("test-*")),
//   	},
//   })
//
type ImageTagMutabilityExclusionFilter interface {
}

// The jsii proxy struct for ImageTagMutabilityExclusionFilter
type jsiiProxy_ImageTagMutabilityExclusionFilter struct {
	_ byte // padding
}

// Creates a wildcard filter for image tag mutability exclusion.
func ImageTagMutabilityExclusionFilter_Wildcard(pattern *string) ImageTagMutabilityExclusionFilter {
	_init_.Initialize()

	if err := validateImageTagMutabilityExclusionFilter_WildcardParameters(pattern); err != nil {
		panic(err)
	}
	var returns ImageTagMutabilityExclusionFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.ImageTagMutabilityExclusionFilter",
		"wildcard",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

