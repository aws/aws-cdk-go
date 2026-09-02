//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAudioSelector_ByLanguageParameters(name *string, languageCode *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if languageCode == nil {
		return fmt.Errorf("parameter languageCode is required, but nil was provided")
	}

	return nil
}

func validateAudioSelector_ByPidParameters(name *string, pids *[]*AudioPidConfig) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if pids == nil {
		return fmt.Errorf("parameter pids is required, but nil was provided")
	}
	for idx_d99a9d, v := range *pids {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter pids[%#v]", idx_d99a9d) }); err != nil {
			return err
		}
	}

	return nil
}

func validateAudioSelector_ByTrackParameters(name *string, tracks *[]*AudioTrackConfig) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if tracks == nil {
		return fmt.Errorf("parameter tracks is required, but nil was provided")
	}
	for idx_387668, v := range *tracks {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter tracks[%#v]", idx_387668) }); err != nil {
			return err
		}
	}

	return nil
}

func validateAudioSelector_DefaultParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateAudioSelector_HlsRenditionParameters(name *string, options *HlsRenditionSelectionOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewAudioSelectorParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

