package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// The destination type for the flow log.
//
// Example:
//   var vpc vpc
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyCustomLogGroup"))
//
//   role := iam.NewRole(this, jsii.String("MyCustomRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("vpc-flow-logs.amazonaws.com")),
//   })
//
//   ec2.NewFlowLog(this, jsii.String("FlowLog"), &FlowLogProps{
//   	ResourceType: ec2.FlowLogResourceType_FromVpc(vpc),
//   	Destination: ec2.FlowLogDestination_ToCloudWatchLogs(logGroup, role),
//   })
//
type FlowLogDestination interface {
	// Generates a flow log destination configuration.
	Bind(scope constructs.Construct, flowLog FlowLog) *FlowLogDestinationConfig
}

// The jsii proxy struct for FlowLogDestination
type jsiiProxy_FlowLogDestination struct {
	_ byte // padding
}

func NewFlowLogDestination_Override(f FlowLogDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.FlowLogDestination",
		nil, // no parameters
		f,
	)
}

// Use CloudWatch logs as the destination.
func FlowLogDestination_ToCloudWatchLogs(logGroup awslogs.ILogGroup, iamRole awsiam.IRole) FlowLogDestination {
	_init_.Initialize()

	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.FlowLogDestination",
		"toCloudWatchLogs",
		[]interface{}{logGroup, iamRole},
		&returns,
	)

	return returns
}

// Use Kinesis Data Firehose as the destination.
func FlowLogDestination_ToKinesisDataFirehoseDestination(deliveryStreamArn *string) FlowLogDestination {
	_init_.Initialize()

	if err := validateFlowLogDestination_ToKinesisDataFirehoseDestinationParameters(deliveryStreamArn); err != nil {
		panic(err)
	}
	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.FlowLogDestination",
		"toKinesisDataFirehoseDestination",
		[]interface{}{deliveryStreamArn},
		&returns,
	)

	return returns
}

// Use S3 as the destination.
func FlowLogDestination_ToS3(bucket awss3.IBucket, keyPrefix *string, options *S3DestinationOptions) FlowLogDestination {
	_init_.Initialize()

	if err := validateFlowLogDestination_ToS3Parameters(options); err != nil {
		panic(err)
	}
	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.FlowLogDestination",
		"toS3",
		[]interface{}{bucket, keyPrefix, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestination) Bind(scope constructs.Construct, flowLog FlowLog) *FlowLogDestinationConfig {
	if err := f.validateBindParameters(scope, flowLog); err != nil {
		panic(err)
	}
	var returns *FlowLogDestinationConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, flowLog},
		&returns,
	)

	return returns
}

