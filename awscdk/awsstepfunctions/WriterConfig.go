package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration to format the output.
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
type WriterConfig interface {
	// The format of the Output of the child workflow executions.
	OutputType() OutputType
	// The transformation to be applied to the Output of the Child Workflow executions.
	Transformation() Transformation
}

// The jsii proxy struct for WriterConfig
type jsiiProxy_WriterConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_WriterConfig) OutputType() OutputType {
	var returns OutputType
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WriterConfig) Transformation() Transformation {
	var returns Transformation
	_jsii_.Get(
		j,
		"transformation",
		&returns,
	)
	return returns
}


func NewWriterConfig(props *WriterConfigProps) WriterConfig {
	_init_.Initialize()

	if err := validateNewWriterConfigParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WriterConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.WriterConfig",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewWriterConfig_Override(w WriterConfig, props *WriterConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.WriterConfig",
		[]interface{}{props},
		w,
	)
}

