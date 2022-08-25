package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// The destination type for the flow log.
//
// Example:
//   var vpc vpc
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyCustomLogGroup"))
//
//   role := iam.NewRole(this, jsii.String("MyCustomRole"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("vpc-flow-logs.amazonaws.com")),
//   })
//
//   ec2.NewFlowLog(this, jsii.String("FlowLog"), &flowLogProps{
//   	resourceType: ec2.flowLogResourceType.fromVpc(vpc),
//   	destination: ec2.flowLogDestination.toCloudWatchLogs(logGroup, role),
//   })
//
// Experimental.
type FlowLogDestination interface {
	// Generates a flow log destination configuration.
	// Experimental.
	Bind(scope awscdk.Construct, flowLog FlowLog) *FlowLogDestinationConfig
}

// The jsii proxy struct for FlowLogDestination
type jsiiProxy_FlowLogDestination struct {
	_ byte // padding
}

// Experimental.
func NewFlowLogDestination_Override(f FlowLogDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.FlowLogDestination",
		nil, // no parameters
		f,
	)
}

// Use CloudWatch logs as the destination.
// Experimental.
func FlowLogDestination_ToCloudWatchLogs(logGroup awslogs.ILogGroup, iamRole awsiam.IRole) FlowLogDestination {
	_init_.Initialize()

	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.FlowLogDestination",
		"toCloudWatchLogs",
		[]interface{}{logGroup, iamRole},
		&returns,
	)

	return returns
}

// Use S3 as the destination.
// Experimental.
func FlowLogDestination_ToS3(bucket awss3.IBucket, keyPrefix *string) FlowLogDestination {
	_init_.Initialize()

	var returns FlowLogDestination

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.FlowLogDestination",
		"toS3",
		[]interface{}{bucket, keyPrefix},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FlowLogDestination) Bind(scope awscdk.Construct, flowLog FlowLog) *FlowLogDestinationConfig {
	var returns *FlowLogDestinationConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, flowLog},
		&returns,
	)

	return returns
}

