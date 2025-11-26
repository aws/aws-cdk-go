//go:build !no_runtime_type_checking

package previewawsecrevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr"
)

func (r *jsiiProxy_RepositoryEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *RepositoryEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateECRImageActionPatternParameters(options *RepositoryEvents_ECRImageAction_ECRImageActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateECRImageScanPatternParameters(options *RepositoryEvents_ECRImageScan_ECRImageScanProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateECRPullThroughCacheActionPatternParameters(options *RepositoryEvents_ECRPullThroughCacheAction_ECRPullThroughCacheActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateECRReferrerActionPatternParameters(options *RepositoryEvents_ECRReferrerAction_ECRReferrerActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RepositoryEvents) validateECRReplicationActionPatternParameters(options *RepositoryEvents_ECRReplicationAction_ECRReplicationActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateRepositoryEvents_FromRepositoryParameters(repositoryRef interfacesawsecr.IRepositoryRef) error {
	if repositoryRef == nil {
		return fmt.Errorf("parameter repositoryRef is required, but nil was provided")
	}

	return nil
}

