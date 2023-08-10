package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Use an Amazon SNS topic as an event destination.
//
// Example:
//   var topic topic
//
//
//   appconfig.NewExtension(this, jsii.String("MyExtension"), &ExtensionProps{
//   	Actions: []action{
//   		appconfig.NewAction(&ActionProps{
//   			ActionPoints: []actionPoint{
//   				appconfig.*actionPoint_ON_DEPLOYMENT_START,
//   			},
//   			EventDestination: appconfig.NewSnsDestination(topic),
//   		}),
//   	},
//   })
//
// Experimental.
type SnsDestination interface {
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

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	jsiiProxy_IEventDestination
}

func (j *jsiiProxy_SnsDestination) ExtensionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsDestination) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsDestination) Type() SourceType {
	var returns SourceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	if err := validateNewSnsDestinationParameters(topic); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

