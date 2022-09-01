package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of resource to create the flow log for.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type FlowLogResourceType interface {
	// The Id of the resource that the flow log should be attached to.
	ResourceId() *string
	SetResourceId(val *string)
	// The type of resource to attach a flow log to.
	ResourceType() *string
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


func NewFlowLogResourceType_Override(f FlowLogResourceType) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.FlowLogResourceType",
		nil, // no parameters
		f,
	)
}

func (j *jsiiProxy_FlowLogResourceType)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_FlowLogResourceType)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

// The Network Interface to attach the Flow Log to.
func FlowLogResourceType_FromNetworkInterfaceId(id *string) FlowLogResourceType {
	_init_.Initialize()

	if err := validateFlowLogResourceType_FromNetworkInterfaceIdParameters(id); err != nil {
		panic(err)
	}
	var returns FlowLogResourceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.FlowLogResourceType",
		"fromNetworkInterfaceId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// The subnet to attach the Flow Log to.
func FlowLogResourceType_FromSubnet(subnet ISubnet) FlowLogResourceType {
	_init_.Initialize()

	if err := validateFlowLogResourceType_FromSubnetParameters(subnet); err != nil {
		panic(err)
	}
	var returns FlowLogResourceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.FlowLogResourceType",
		"fromSubnet",
		[]interface{}{subnet},
		&returns,
	)

	return returns
}

// The VPC to attach the Flow Log to.
func FlowLogResourceType_FromVpc(vpc IVpc) FlowLogResourceType {
	_init_.Initialize()

	if err := validateFlowLogResourceType_FromVpcParameters(vpc); err != nil {
		panic(err)
	}
	var returns FlowLogResourceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.FlowLogResourceType",
		"fromVpc",
		[]interface{}{vpc},
		&returns,
	)

	return returns
}

