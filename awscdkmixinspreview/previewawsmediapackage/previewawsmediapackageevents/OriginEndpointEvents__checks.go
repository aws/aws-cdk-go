//go:build !no_runtime_type_checking

package previewawsmediapackageevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackage"
)

func (o *jsiiProxy_OriginEndpointEvents) validateMediaPackageHarvestJobNotificationPatternParameters(options *OriginEndpointEvents_MediaPackageHarvestJobNotification_MediaPackageHarvestJobNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateOriginEndpointEvents_FromOriginEndpointParameters(originEndpointRef interfacesawsmediapackage.IOriginEndpointRef) error {
	if originEndpointRef == nil {
		return fmt.Errorf("parameter originEndpointRef is required, but nil was provided")
	}

	return nil
}

