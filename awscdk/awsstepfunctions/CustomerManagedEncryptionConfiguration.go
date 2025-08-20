package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Define a new CustomerManagedEncryptionConfiguration.
//
// Example:
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   kmsKey := kms.NewKey(this, jsii.String("Key"))
//   stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachineWithCMKEncryptionConfiguration"), &StateMachineProps{
//   	StateMachineName: jsii.String("StateMachineWithCMKEncryptionConfiguration"),
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass")))),
//   	StateMachineType: sfn.StateMachineType_STANDARD,
//   	EncryptionConfiguration: sfn.NewCustomerManagedEncryptionConfiguration(kmsKey, cdk.Duration_Seconds(jsii.Number(60))),
//   })
//
type CustomerManagedEncryptionConfiguration interface {
	EncryptionConfiguration
	// Maximum duration that Step Functions will reuse customer managed data keys. When the period expires, Step Functions will call GenerateDataKey.
	//
	// Must be between 60 and 900 seconds.
	// Default: Duration.seconds(300)
	//
	KmsDataKeyReusePeriodSeconds() awscdk.Duration
	// The symmetric customer managed KMS key for server-side encryption of the state machine definition, and execution history or activity inputs.
	//
	// Step Functions will reuse the key for a maximum of `kmsDataKeyReusePeriodSeconds`.
	// Default: - data is transparently encrypted using an AWS owned key.
	//
	KmsKey() awskms.IKey
	// Encryption option for the state machine or activity.
	//
	// Can be either CUSTOMER_MANAGED_KMS_KEY or AWS_OWNED_KEY.
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for CustomerManagedEncryptionConfiguration
type jsiiProxy_CustomerManagedEncryptionConfiguration struct {
	jsiiProxy_EncryptionConfiguration
}

func (j *jsiiProxy_CustomerManagedEncryptionConfiguration) KmsDataKeyReusePeriodSeconds() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"kmsDataKeyReusePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerManagedEncryptionConfiguration) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerManagedEncryptionConfiguration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewCustomerManagedEncryptionConfiguration(kmsKey awskms.IKey, kmsDataKeyReusePeriodSeconds awscdk.Duration) CustomerManagedEncryptionConfiguration {
	_init_.Initialize()

	if err := validateNewCustomerManagedEncryptionConfigurationParameters(kmsKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerManagedEncryptionConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.CustomerManagedEncryptionConfiguration",
		[]interface{}{kmsKey, kmsDataKeyReusePeriodSeconds},
		&j,
	)

	return &j
}

func NewCustomerManagedEncryptionConfiguration_Override(c CustomerManagedEncryptionConfiguration, kmsKey awskms.IKey, kmsDataKeyReusePeriodSeconds awscdk.Duration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.CustomerManagedEncryptionConfiguration",
		[]interface{}{kmsKey, kmsDataKeyReusePeriodSeconds},
		c,
	)
}

func (j *jsiiProxy_CustomerManagedEncryptionConfiguration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

