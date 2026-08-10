package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Glue Data Catalog, either the implicit account-wide catalog or one created as an `AWS::Glue::Catalog` resource.
// Experimental.
type ICatalog interface {
	interfacesawsglue.ICatalogRef
	awscdk.IResource
	// The ARN of the catalog.
	// Experimental.
	CatalogArn() *string
	// The id of the catalog (for the account-wide catalog, the AWS account id).
	// Experimental.
	CatalogId() *string
	// The customer-managed KMS key used to encrypt connection passwords, if one was configured.
	//
	// Undefined when password encryption uses an AWS-managed key or is not
	// configured. Grant access to it via `KeyGrants`, e.g.
	// `if (catalog.connectionPasswordKey) { KeyGrants.fromKey(catalog.connectionPasswordKey).encrypt(grantee); }`.
	// Experimental.
	ConnectionPasswordKey() interfacesawskms.IKeyRef
	// The customer-managed KMS key used for the catalog's encryption at rest, if one was configured.
	//
	// Undefined when encryption is disabled or an AWS-managed key is used. Grant
	// access to it via `KeyGrants`, e.g.
	// `if (catalog.encryptionKey) { KeyGrants.fromKey(catalog.encryptionKey).encrypt(grantee); }`.
	// Experimental.
	EncryptionKey() interfacesawskms.IKeyRef
}

// The jsii proxy for ICatalog
type jsiiProxy_ICatalog struct {
	internal.Type__interfacesawsglueICatalogRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICatalog) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ICatalog) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICatalog) CatalogArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) ConnectionPasswordKey() interfacesawskms.IKeyRef {
	var returns interfacesawskms.IKeyRef
	_jsii_.Get(
		j,
		"connectionPasswordKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) EncryptionKey() interfacesawskms.IKeyRef {
	var returns interfacesawskms.IKeyRef
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) CatalogRef() *interfacesawsglue.CatalogReference {
	var returns *interfacesawsglue.CatalogReference
	_jsii_.Get(
		j,
		"catalogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICatalog) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

