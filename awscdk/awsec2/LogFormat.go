package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The following table describes all of the available fields for a flow log record.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   vpc.addFlowLog(jsii.String("FlowLog"), &FlowLogOptions{
//   	LogFormat: []logFormat{
//   		ec2.*logFormat_DST_PORT(),
//   		ec2.*logFormat_SRC_PORT(),
//   	},
//   })
//
//   // If you just want to add a field to the default field
//   vpc.addFlowLog(jsii.String("FlowLog"), &FlowLogOptions{
//   	LogFormat: []*logFormat{
//   		ec2.*logFormat_VERSION(),
//   		ec2.*logFormat_ALL_DEFAULT_FIELDS(),
//   	},
//   })
//
//   // If AWS CDK does not support the new fields
//   vpc.addFlowLog(jsii.String("FlowLog"), &FlowLogOptions{
//   	LogFormat: []*logFormat{
//   		ec2.*logFormat_SRC_PORT(),
//   		ec2.*logFormat_Custom(jsii.String("${new-field}")),
//   	},
//   })
//
type LogFormat interface {
	Value() *string
}

// The jsii proxy struct for LogFormat
type jsiiProxy_LogFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_LogFormat) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewLogFormat(value *string) LogFormat {
	_init_.Initialize()

	if err := validateNewLogFormatParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.LogFormat",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewLogFormat_Override(l LogFormat, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.LogFormat",
		[]interface{}{value},
		l,
	)
}

// A custom format string.
//
// Gives full control over the format string fragment.
func LogFormat_Custom(formatString *string) LogFormat {
	_init_.Initialize()

	if err := validateLogFormat_CustomParameters(formatString); err != nil {
		panic(err)
	}
	var returns LogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"custom",
		[]interface{}{formatString},
		&returns,
	)

	return returns
}

// A custom field name.
//
// If there is no ready-made constant for a new field yet, you can use this.
// The field name will automatically be wrapped in `${ ... }`.
func LogFormat_Field(field *string) LogFormat {
	_init_.Initialize()

	if err := validateLogFormat_FieldParameters(field); err != nil {
		panic(err)
	}
	var returns LogFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"field",
		[]interface{}{field},
		&returns,
	)

	return returns
}

func LogFormat_ACCOUNT_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"ACCOUNT_ID",
		&returns,
	)
	return returns
}

func LogFormat_ACTION() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"ACTION",
		&returns,
	)
	return returns
}

func LogFormat_ALL_DEFAULT_FIELDS() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"ALL_DEFAULT_FIELDS",
		&returns,
	)
	return returns
}

func LogFormat_AZ_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"AZ_ID",
		&returns,
	)
	return returns
}

func LogFormat_BYTES() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"BYTES",
		&returns,
	)
	return returns
}

func LogFormat_DST_ADDR() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"DST_ADDR",
		&returns,
	)
	return returns
}

func LogFormat_DST_PORT() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"DST_PORT",
		&returns,
	)
	return returns
}

func LogFormat_END_TIMESTAMP() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"END_TIMESTAMP",
		&returns,
	)
	return returns
}

func LogFormat_FLOW_DIRECTION() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"FLOW_DIRECTION",
		&returns,
	)
	return returns
}

func LogFormat_INSTANCE_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"INSTANCE_ID",
		&returns,
	)
	return returns
}

func LogFormat_INTERFACE_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"INTERFACE_ID",
		&returns,
	)
	return returns
}

func LogFormat_LOG_STATUS() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"LOG_STATUS",
		&returns,
	)
	return returns
}

func LogFormat_PACKETS() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"PACKETS",
		&returns,
	)
	return returns
}

func LogFormat_PKT_DST_ADDR() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"PKT_DST_ADDR",
		&returns,
	)
	return returns
}

func LogFormat_PKT_DST_AWS_SERVICE() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"PKT_DST_AWS_SERVICE",
		&returns,
	)
	return returns
}

func LogFormat_PKT_SRC_ADDR() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"PKT_SRC_ADDR",
		&returns,
	)
	return returns
}

func LogFormat_PKT_SRC_AWS_SERVICE() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"PKT_SRC_AWS_SERVICE",
		&returns,
	)
	return returns
}

func LogFormat_PROTOCOL() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"PROTOCOL",
		&returns,
	)
	return returns
}

func LogFormat_REGION() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"REGION",
		&returns,
	)
	return returns
}

func LogFormat_SRC_ADDR() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"SRC_ADDR",
		&returns,
	)
	return returns
}

func LogFormat_SRC_PORT() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"SRC_PORT",
		&returns,
	)
	return returns
}

func LogFormat_START_TIMESTAMP() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"START_TIMESTAMP",
		&returns,
	)
	return returns
}

func LogFormat_SUBLOCATION_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"SUBLOCATION_ID",
		&returns,
	)
	return returns
}

func LogFormat_SUBLOCATION_TYPE() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"SUBLOCATION_TYPE",
		&returns,
	)
	return returns
}

func LogFormat_SUBNET_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"SUBNET_ID",
		&returns,
	)
	return returns
}

func LogFormat_TCP_FLAGS() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"TCP_FLAGS",
		&returns,
	)
	return returns
}

func LogFormat_TRAFFIC_PATH() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"TRAFFIC_PATH",
		&returns,
	)
	return returns
}

func LogFormat_TRAFFIC_TYPE() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"TRAFFIC_TYPE",
		&returns,
	)
	return returns
}

func LogFormat_VERSION() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"VERSION",
		&returns,
	)
	return returns
}

func LogFormat_VPC_ID() LogFormat {
	_init_.Initialize()
	var returns LogFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.LogFormat",
		"VPC_ID",
		&returns,
	)
	return returns
}

