//go:build !no_runtime_type_checking

package previewawskmsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

func (k *jsiiProxy_KeyEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *KeyEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KeyEvents) validateKMSCMKDeletionPatternParameters(options *KeyEvents_KMSCMKDeletion_KMSCMKDeletionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KeyEvents) validateKMSCMKRotationPatternParameters(options *KeyEvents_KMSCMKRotation_KMSCMKRotationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KeyEvents) validateKMSImportedKeyMaterialExpirationPatternParameters(options *KeyEvents_KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateKeyEvents_FromKeyParameters(keyRef interfacesawskms.IKeyRef) error {
	if keyRef == nil {
		return fmt.Errorf("parameter keyRef is required, but nil was provided")
	}

	return nil
}

