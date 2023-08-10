package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Use an Amazon SQS queue as an event destination.
//
// Example:
//   var queue queue
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []actionPoint{
//   				appconfig.*actionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewSqsDestination(queue),
//   		}),
//   	},
//   })
//
// Experimental.
type SqsDestination interface {
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

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	jsiiProxy_IEventDestination
}

func (j *jsiiProxy_SqsDestination) ExtensionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsDestination) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsDestination) Type() SourceType {
	var returns SourceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	if err := validateNewSqsDestinationParameters(queue); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

