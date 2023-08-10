package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Use an AWS Lambda function as an event destination.
//
// Example:
//   var fn function
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []actionPoint{
//   				appconfig.*actionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewLambdaDestination(fn),
//   		}),
//   	},
//   })
//
// Experimental.
type LambdaDestination interface {
	IEventDestination
	// The URI of the extension event destination.
	// Experimental.
	ExtensionUri() *string
	// The IAM policy document to invoke the event destination.
	// Experimental.
	PolicyDocument() awsiam.PolicyDocument
	// The type of the extension event destination.
	// Experimental.
	Type() SourceType
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	jsiiProxy_IEventDestination
}

func (j *jsiiProxy_LambdaDestination) ExtensionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDestination) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDestination) Type() SourceType {
	var returns SourceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaDestination(func_ awslambda.IFunction) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(func_); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.LambdaDestination",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaDestination_Override(l LambdaDestination, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.LambdaDestination",
		[]interface{}{func_},
		l,
	)
}

