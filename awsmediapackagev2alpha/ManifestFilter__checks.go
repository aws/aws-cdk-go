//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateManifestFilter_AudioCodecParameters(value AudioCodec) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_AudioCodecListParameters(values *[]AudioCodec) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_BitrateParameters(key BitrateFilterKey, value awscdk.Bitrate) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_BitrateComboParameters(key BitrateFilterKey, expressions *[]BitrateExpression) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if expressions == nil {
		return fmt.Errorf("parameter expressions is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_BitrateRangeParameters(key BitrateFilterKey, start awscdk.Bitrate, end awscdk.Bitrate) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_CustomParameters(custom *string) error {
	if custom == nil {
		return fmt.Errorf("parameter custom is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_NumericParameters(key NumericFilterKey, value *float64) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_NumericComboParameters(key NumericFilterKey, expressions *[]NumericExpression) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if expressions == nil {
		return fmt.Errorf("parameter expressions is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_NumericListParameters(key NumericFilterKey, values *[]*float64) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_NumericRangeParameters(key NumericFilterKey, start *float64, end *float64) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_TextParameters(key TextFilterKey, value *string) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_TextListParameters(key TextFilterKey, values *[]*string) error {
	if key == "" {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_TrickplayTypeParameters(value TrickplayType) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_TrickplayTypeListParameters(values *[]TrickplayType) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_VideoCodecParameters(value VideoCodec) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_VideoCodecListParameters(values *[]VideoCodec) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_VideoDynamicRangeParameters(value VideoDynamicRange) error {
	if value == "" {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateManifestFilter_VideoDynamicRangeListParameters(values *[]VideoDynamicRange) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateNewManifestFilterParameters(filterString *string) error {
	if filterString == nil {
		return fmt.Errorf("parameter filterString is required, but nil was provided")
	}

	return nil
}

