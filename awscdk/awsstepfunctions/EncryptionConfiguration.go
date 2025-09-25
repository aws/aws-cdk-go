package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for creating an EncryptionConfiguration for either state machines or activities.
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
type EncryptionConfiguration interface {
	// Encryption option for the state machine or activity.
	//
	// Can be either CUSTOMER_MANAGED_KMS_KEY or AWS_OWNED_KEY.
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for EncryptionConfiguration
type jsiiProxy_EncryptionConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_EncryptionConfiguration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewEncryptionConfiguration_Override(e EncryptionConfiguration, type_ *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.EncryptionConfiguration",
		[]interface{}{type_},
		e,
	)
}

func (j *jsiiProxy_EncryptionConfiguration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

