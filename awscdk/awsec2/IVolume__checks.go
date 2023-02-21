//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IVolume) validateGrantAttachVolumeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVolume) validateGrantAttachVolumeByResourceTagParameters(grantee awsiam.IGrantable, constructs *[]constructs.Construct) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if constructs == nil {
		return fmt.Errorf("parameter constructs is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVolume) validateGrantDetachVolumeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVolume) validateGrantDetachVolumeByResourceTagParameters(grantee awsiam.IGrantable, constructs *[]constructs.Construct) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if constructs == nil {
		return fmt.Errorf("parameter constructs is required, but nil was provided")
	}

	return nil
}

