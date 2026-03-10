//go:build !no_runtime_type_checking

package previewawscodecommitevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCodeCommitPullRequestStateChange_CodeCommitPullRequestStateChangePatternParameters(options *CodeCommitPullRequestStateChange_CodeCommitPullRequestStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

