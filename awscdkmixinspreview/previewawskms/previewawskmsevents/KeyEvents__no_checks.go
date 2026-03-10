//go:build no_runtime_type_checking

package previewawskmsevents

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (k *jsiiProxy_KeyEvents) validateKmsCMKDeletionPatternParameters(options *KMSCMKDeletion_KMSCMKDeletionProps) error {
	return nil
}

func (k *jsiiProxy_KeyEvents) validateKmsCMKRotationPatternParameters(options *KMSCMKRotation_KMSCMKRotationProps) error {
	return nil
}

func (k *jsiiProxy_KeyEvents) validateKmsImportedKeyMaterialExpirationPatternParameters(options *KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps) error {
	return nil
}

func validateKeyEvents_FromKeyParameters(keyRef interfacesawskms.IKeyRef) error {
	return nil
}

