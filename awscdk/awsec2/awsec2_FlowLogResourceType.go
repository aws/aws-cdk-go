package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of resource to create the flow log for.
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
type FlowLogResourceType interface {
	// The Id of the resource that the flow log should be attached to.
	// Experimental.
	ResourceId() *string
	// Experimental.
	SetResourceId(val *string)
	// The type of resource to attach a flow log to.
	// Experimental.
	ResourceType() *string
	// Experimental.
	SetResourceType(val *string)
}

// The jsii proxy struct for FlowLogResourceType
type jsiiProxy_FlowLogResourceType struct {
	_ byte // padding
}

func (j *jsiiProxy_FlowLogResourceType) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FlowLogResourceType) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewFlowLogResourceType_Override(f FlowLogResourceType) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.FlowLogResourceType",
		nil, // no parameters
		f,
	)
}

func (j *jsiiProxy_FlowLogResourceType) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_FlowLogResourceType) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

// The Network Interface to attach the Flow Log to.
// Experimental.
func FlowLogResourceType_FromNetworkInterfaceId(id *string) FlowLogResourceType {
	_init_.Initialize()

	var returns FlowLogResourceType

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.FlowLogResourceType",
		"fromNetworkInterfaceId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// The subnet to attach the Flow Log to.
// Experimental.
func FlowLogResourceType_FromSubnet(subnet ISubnet) FlowLogResourceType {
	_init_.Initialize()

	var returns FlowLogResourceType

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.FlowLogResourceType",
		"fromSubnet",
		[]interface{}{subnet},
		&returns,
	)

	return returns
}

// The VPC to attach the Flow Log to.
// Experimental.
func FlowLogResourceType_FromVpc(vpc IVpc) FlowLogResourceType {
	_init_.Initialize()

	var returns FlowLogResourceType

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.FlowLogResourceType",
		"fromVpc",
		[]interface{}{vpc},
		&returns,
	)

	return returns
}

