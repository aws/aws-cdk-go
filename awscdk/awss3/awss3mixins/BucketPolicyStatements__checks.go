//go:build !no_runtime_type_checking

package awss3mixins

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BucketPolicyStatements) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketPolicyStatements) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBucketPolicyStatements_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewBucketPolicyStatementsParameters(statements *[]awsiam.PolicyStatement) error {
	if statements == nil {
		return fmt.Errorf("parameter statements is required, but nil was provided")
	}

	return nil
}

