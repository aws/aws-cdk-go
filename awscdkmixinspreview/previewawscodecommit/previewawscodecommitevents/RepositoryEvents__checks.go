//go:build !no_runtime_type_checking

package previewawscodecommitevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodecommit"
)

func (r *jsiiProxy_RepositoryEvents) validateCodeCommitCommentOnCommitPatternParameters(options *RepositoryEvents_CodeCommitCommentOnCommit_CodeCommitCommentOnCommitProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateCodeCommitCommentOnPullRequestPatternParameters(options *RepositoryEvents_CodeCommitCommentOnPullRequest_CodeCommitCommentOnPullRequestProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateCodeCommitRepositoryStateChangePatternParameters(options *RepositoryEvents_CodeCommitRepositoryStateChange_CodeCommitRepositoryStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateRepositoryEvents_FromRepositoryParameters(repositoryRef interfacesawscodecommit.IRepositoryRef) error {
	if repositoryRef == nil {
		return fmt.Errorf("parameter repositoryRef is required, but nil was provided")
	}

	return nil
}

