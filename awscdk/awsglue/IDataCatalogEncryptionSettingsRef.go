package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataCatalogEncryptionSettings.
// Experimental.
type IDataCatalogEncryptionSettingsRef interface {
	constructs.IConstruct
	// A reference to a DataCatalogEncryptionSettings resource.
	// Experimental.
	DataCatalogEncryptionSettingsRef() *DataCatalogEncryptionSettingsReference
}

// The jsii proxy for IDataCatalogEncryptionSettingsRef
type jsiiProxy_IDataCatalogEncryptionSettingsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataCatalogEncryptionSettingsRef) DataCatalogEncryptionSettingsRef() *DataCatalogEncryptionSettingsReference {
	var returns *DataCatalogEncryptionSettingsReference
	_jsii_.Get(
		j,
		"dataCatalogEncryptionSettingsRef",
		&returns,
	)
	return returns
}

