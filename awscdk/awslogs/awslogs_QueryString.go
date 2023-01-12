package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Define a QueryString.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &queryDefinitionProps{
//   	queryDefinitionName: jsii.String("MyQuery"),
//   	queryString: logs.NewQueryString(&queryStringProps{
//   		fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		sort: jsii.String("@timestamp desc"),
//   		limit: jsii.Number(20),
//   	}),
//   })
//
type QueryString interface {
	// String representation of this QueryString.
	ToString() *string
}

// The jsii proxy struct for QueryString
type jsiiProxy_QueryString struct {
	_ byte // padding
}

func NewQueryString(props *QueryStringProps) QueryString {
	_init_.Initialize()

	if err := validateNewQueryStringParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueryString{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.QueryString",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewQueryString_Override(q QueryString, props *QueryStringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.QueryString",
		[]interface{}{props},
		q,
	)
}

func (q *jsiiProxy_QueryString) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

