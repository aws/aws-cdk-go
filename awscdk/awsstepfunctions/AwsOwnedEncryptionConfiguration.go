package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Define a new AwsOwnedEncryptionConfiguration.
//
// Example:
//   stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	StateMachineName: jsii.String("StateMachine"),
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass")))),
//   	StateMachineType: sfn.StateMachineType_STANDARD,
//   	EncryptionConfiguration: sfn.NewAwsOwnedEncryptionConfiguration(),
//   })
//
type AwsOwnedEncryptionConfiguration interface {
	EncryptionConfiguration
	// Encryption option for the state machine or activity.
	//
	// Can be either CUSTOMER_MANAGED_KMS_KEY or AWS_OWNED_KEY.
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for AwsOwnedEncryptionConfiguration
type jsiiProxy_AwsOwnedEncryptionConfiguration struct {
	jsiiProxy_EncryptionConfiguration
}

func (j *jsiiProxy_AwsOwnedEncryptionConfiguration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewAwsOwnedEncryptionConfiguration() AwsOwnedEncryptionConfiguration {
	_init_.Initialize()

	j := jsiiProxy_AwsOwnedEncryptionConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.AwsOwnedEncryptionConfiguration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAwsOwnedEncryptionConfiguration_Override(a AwsOwnedEncryptionConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.AwsOwnedEncryptionConfiguration",
		nil, // no parameters
		a,
	)
}

func (j *jsiiProxy_AwsOwnedEncryptionConfiguration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

