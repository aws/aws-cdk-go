//go:build no_runtime_type_checking

package previewawskmsevents

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *KeyEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (k *jsiiProxy_KeyEvents) validateKMSCMKDeletionPatternParameters(options *KeyEvents_KMSCMKDeletion_KMSCMKDeletionProps) error {
	return nil
}

func (k *jsiiProxy_KeyEvents) validateKMSCMKRotationPatternParameters(options *KeyEvents_KMSCMKRotation_KMSCMKRotationProps) error {
	return nil
}

func (k *jsiiProxy_KeyEvents) validateKMSImportedKeyMaterialExpirationPatternParameters(options *KeyEvents_KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps) error {
	return nil
}

func validateKeyEvents_FromKeyParameters(keyRef interfacesawskms.IKeyRef) error {
	return nil
}

