package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The format of the source data.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket IBucket
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"))
//
//   dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ImportSource: &ImportSourceSpecification{
//   		CompressionType: dynamodb.InputCompressionType_GZIP,
//   		InputFormat: dynamodb.InputFormat_DynamoDBJson(),
//   		Bucket: *Bucket,
//   		KeyPrefix: jsii.String("prefix"),
//   	},
//   })
//
type InputFormat interface {
}

// The jsii proxy struct for InputFormat
type jsiiProxy_InputFormat struct {
	_ byte // padding
}

func NewInputFormat_Override(i InputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.InputFormat",
		nil, // no parameters
		i,
	)
}

// CSV format.
func InputFormat_Csv(options *CsvOptions) InputFormat {
	_init_.Initialize()

	if err := validateInputFormat_CsvParameters(options); err != nil {
		panic(err)
	}
	var returns InputFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.InputFormat",
		"csv",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// DynamoDB JSON format.
func InputFormat_DynamoDBJson() InputFormat {
	_init_.Initialize()

	var returns InputFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.InputFormat",
		"dynamoDBJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon Ion format.
func InputFormat_Ion() InputFormat {
	_init_.Initialize()

	var returns InputFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.InputFormat",
		"ion",
		nil, // no parameters
		&returns,
	)

	return returns
}

