package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Accessor for scoped pseudo parameters.
//
// These pseudo parameters are anchored to a stack somewhere in the construct
// tree, and their values will be exported automatically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   scopedAws := cdk.NewScopedAws(this)
//
type ScopedAws interface {
	AccountId() *string
	NotificationArns() *[]*string
	Partition() *string
	Region() *string
	StackId() *string
	StackName() *string
	UrlSuffix() *string
}

// The jsii proxy struct for ScopedAws
type jsiiProxy_ScopedAws struct {
	_ byte // padding
}

func (j *jsiiProxy_ScopedAws) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScopedAws) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


func NewScopedAws(scope constructs.Construct) ScopedAws {
	_init_.Initialize()

	if err := validateNewScopedAwsParameters(scope); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScopedAws{}

	_jsii_.Create(
		"aws-cdk-lib.ScopedAws",
		[]interface{}{scope},
		&j,
	)

	return &j
}

func NewScopedAws_Override(s ScopedAws, scope constructs.Construct) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.ScopedAws",
		[]interface{}{scope},
		s,
	)
}

