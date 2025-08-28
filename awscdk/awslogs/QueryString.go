package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   		ParseStatements: []*string{
//   			jsii.String("@message \"[*] *\" as loggingType, loggingMessage"),
//   			jsii.String("@message \"<*>: *\" as differentLoggingType, differentLoggingMessage"),
//   		},
//   		FilterStatements: []*string{
//   			jsii.String("loggingType = \"ERROR\""),
//   			jsii.String("loggingMessage = \"A very strange error occurred!\""),
//   		},
//   		StatsStatements: []*string{
//   			jsii.String("count(loggingMessage) as loggingErrors"),
//   			jsii.String("count(differentLoggingMessage) as differentLoggingErrors"),
//   		},
//   		Sort: jsii.String("@timestamp desc"),
//   		Limit: jsii.Number(20),
//   	}),
//   })
//
type QueryString interface {
	// If the props for the query string have both stats and statsStatements.
	HasStatsAndStatsStatements() *bool
	// Length of statsStatements.
	StatsStatementsLength() *float64
	// String representation of this QueryString.
	ToString() *string
}

// The jsii proxy struct for QueryString
type jsiiProxy_QueryString struct {
	_ byte // padding
}

func (j *jsiiProxy_QueryString) HasStatsAndStatsStatements() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasStatsAndStatsStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QueryString) StatsStatementsLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statsStatementsLength",
		&returns,
	)
	return returns
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

