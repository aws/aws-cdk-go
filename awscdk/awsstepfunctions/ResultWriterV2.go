package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Configuration for writing Distributed Map state results to S3 The ResultWriter field cannot be empty.
//
// You must specify one of these sets of sub-fields.
//   writerConfig - to preview the formatted output, without saving the results to Amazon S3.
//   bucket and prefix - to save the results to Amazon S3 without additional formatting.
//   All three fields: writerConfig, bucket and prefix - to format the output and save it to Amazon S3.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a bucket
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   // create a WriterConfig
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	ResultWriterV2: sfn.NewResultWriterV2(&ResultWriterV2Props{
//   		Bucket: bucket,
//   		Prefix: jsii.String("my-prefix"),
//   		WriterConfig: map[string]interface{}{
//   			"outputType": sfn.OutputType_JSONL,
//   			"transformation": sfn.Transformation_NONE,
//   		},
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
type ResultWriterV2 interface {
	// S3 Bucket in which to save Map Run results.
	Bucket() awss3.IBucket
	// S3 prefix in which to save Map Run results.
	// Default: - No prefix.
	//
	Prefix() *string
	// Configuration to format the output of the Child Workflow executions.
	WriterConfig() WriterConfig
	// Compile policy statements to provide relevent permissions to the state machine.
	ProvidePolicyStatements() *[]awsiam.PolicyStatement
	// Render ResultWriter in ASL JSON format.
	Render(queryLanguage QueryLanguage) interface{}
}

// The jsii proxy struct for ResultWriterV2
type jsiiProxy_ResultWriterV2 struct {
	_ byte // padding
}

func (j *jsiiProxy_ResultWriterV2) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResultWriterV2) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResultWriterV2) WriterConfig() WriterConfig {
	var returns WriterConfig
	_jsii_.Get(
		j,
		"writerConfig",
		&returns,
	)
	return returns
}


func NewResultWriterV2(props *ResultWriterV2Props) ResultWriterV2 {
	_init_.Initialize()

	if err := validateNewResultWriterV2Parameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResultWriterV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ResultWriterV2",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewResultWriterV2_Override(r ResultWriterV2, props *ResultWriterV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ResultWriterV2",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_ResultWriterV2) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		r,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResultWriterV2) Render(queryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"render",
		[]interface{}{queryLanguage},
		&returns,
	)

	return returns
}

