package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Define a QueryString.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &QueryDefinitionProps{
//   	QueryDefinitionName: jsii.String("MyQuery"),
//   	QueryString: logs.NewQueryString(&QueryStringProps{
//   		Fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		Sort: jsii.String("@timestamp desc"),
//   		Limit: jsii.Number(20),
//   	}),
//   })
//
// Experimental.
type QueryString interface {
	// String representation of this QueryString.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QueryString
type jsiiProxy_QueryString struct {
	_ byte // padding
}

// Experimental.
func NewQueryString(props *QueryStringProps) QueryString {
	_init_.Initialize()

	if err := validateNewQueryStringParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueryString{}

	_jsii_.Create(
		"monocdk.aws_logs.QueryString",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewQueryString_Override(q QueryString, props *QueryStringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs.QueryString",
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

