//go:build !no_runtime_type_checking

package previewawsecsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs"
)

func (c *jsiiProxy_ClusterEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *ClusterEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ClusterEvents) validateECSContainerInstanceStateChangePatternParameters(options *ClusterEvents_ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ClusterEvents) validateECSServiceActionPatternParameters(options *ClusterEvents_ECSServiceAction_ECSServiceActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ClusterEvents) validateECSTaskStateChangePatternParameters(options *ClusterEvents_ECSTaskStateChange_ECSTaskStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateClusterEvents_FromClusterParameters(clusterRef interfacesawsecs.IClusterRef) error {
	if clusterRef == nil {
		return fmt.Errorf("parameter clusterRef is required, but nil was provided")
	}

	return nil
}

