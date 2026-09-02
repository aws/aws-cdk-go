package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Encryption-at-rest configuration for a Glue Data Catalog.
//
// The Data Catalog encryption at rest and the connection password encryption
// are independent: enabling one does not require the other, and each may use a
// different KMS key.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var key Key
//   var role IRole
//
//   glue.Catalog_EncryptAccount(this, &CatalogEncryptionOptions{
//   	EncryptionAtRest: glue.DataCatalogEncryptionAtRest_KmsWithServiceRole(role, key),
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/webapi/API_EncryptionAtRest.html
//
// Experimental.
type DataCatalogEncryptionAtRest interface {
	// The customer-managed KMS key used for encryption at rest, if any.
	// Experimental.
	KmsKey() interfacesawskms.IKeyRef
	// The encryption mode.
	// Experimental.
	Mode() CatalogEncryptionMode
	// The service role that AWS Glue assumes to access the KMS key, if any.
	// Experimental.
	ServiceRole() awsiam.IRole
}

// The jsii proxy struct for DataCatalogEncryptionAtRest
type jsiiProxy_DataCatalogEncryptionAtRest struct {
	_ byte // padding
}

func (j *jsiiProxy_DataCatalogEncryptionAtRest) KmsKey() interfacesawskms.IKeyRef {
	var returns interfacesawskms.IKeyRef
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogEncryptionAtRest) Mode() CatalogEncryptionMode {
	var returns CatalogEncryptionMode
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogEncryptionAtRest) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Disable encryption at rest for the Data Catalog.
// Experimental.
func DataCatalogEncryptionAtRest_Disabled() DataCatalogEncryptionAtRest {
	_init_.Initialize()

	var returns DataCatalogEncryptionAtRest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.DataCatalogEncryptionAtRest",
		"disabled",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Encrypt the Data Catalog at rest with an AWS KMS key.
// Experimental.
func DataCatalogEncryptionAtRest_Kms(key interfacesawskms.IKeyRef) DataCatalogEncryptionAtRest {
	_init_.Initialize()

	var returns DataCatalogEncryptionAtRest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.DataCatalogEncryptionAtRest",
		"kms",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Encrypt the Data Catalog at rest with an AWS KMS key, accessed through a service role that AWS Glue assumes on your behalf.
//
// When a customer-managed `key` is provided, the `role` is automatically
// granted `kms:Encrypt`/`kms:Decrypt`/`kms:GenerateDataKey*` on it.
// Experimental.
func DataCatalogEncryptionAtRest_KmsWithServiceRole(role awsiam.IRole, key interfacesawskms.IKeyRef) DataCatalogEncryptionAtRest {
	_init_.Initialize()

	if err := validateDataCatalogEncryptionAtRest_KmsWithServiceRoleParameters(role); err != nil {
		panic(err)
	}
	var returns DataCatalogEncryptionAtRest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.DataCatalogEncryptionAtRest",
		"kmsWithServiceRole",
		[]interface{}{role, key},
		&returns,
	)

	return returns
}

