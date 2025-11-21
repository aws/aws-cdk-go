package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A set of permissions to grant on a Table Stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//   var tableRef ITableRef
//
//   streamGrants := awscdk.Aws_dynamodb.NewStreamGrants(&StreamGrantsProps{
//   	Table: tableRef,
//   	TableStreamArn: jsii.String("tableStreamArn"),
//
//   	// the properties below are optional
//   	EncryptionKey: key,
//   })
//
type StreamGrants interface {
	// Adds an IAM policy statement associated with this table's stream to an IAM principal's policy.
	//
	// If `encryptionKey` is present, appropriate grants to the key needs to be added
	// separately using the `table.encryptionKey.grant*` methods.
	Actions(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits an IAM Principal to list streams attached to current dynamodb table.
	List(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all stream data read operations for this table's stream: DescribeStream, GetRecords, GetShardIterator, ListStreams.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	Read(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for StreamGrants
type jsiiProxy_StreamGrants struct {
	_ byte // padding
}

func NewStreamGrants(props *StreamGrantsProps) StreamGrants {
	_init_.Initialize()

	if err := validateNewStreamGrantsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StreamGrants{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.StreamGrants",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewStreamGrants_Override(s StreamGrants, props *StreamGrantsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.StreamGrants",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StreamGrants) Actions(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := s.validateActionsParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"actions",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamGrants) List(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateListParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"list",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

