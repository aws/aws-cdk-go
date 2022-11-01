//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CloudFrontTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	if _record == nil {
		return fmt.Errorf("parameter _record is required, but nil was provided")
	}

	return nil
}

func validateCloudFrontTarget_GetHostedZoneIdParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateNewCloudFrontTargetParameters(distribution awscloudfront.IDistribution) error {
	if distribution == nil {
		return fmt.Errorf("parameter distribution is required, but nil was provided")
	}

	return nil
}

