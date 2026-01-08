package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodebuild"
	"github.com/aws/constructs-go/constructs/v10"
)

// The interface representing the ReportGroup resource - either an existing one, imported using the `ReportGroup.fromReportGroupName` method, or a new one, created with the `ReportGroup` class.
type IReportGroup interface {
	interfacesawscodebuild.IReportGroupRef
	awscdk.IResource
	// Grants the given entity permissions to write (that is, upload reports to) this report group.
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	// The ARN of the ReportGroup.
	ReportGroupArn() *string
	// The name of the ReportGroup.
	ReportGroupName() *string
}

// The jsii proxy for IReportGroup
type jsiiProxy_IReportGroup struct {
	internal.Type__interfacesawscodebuildIReportGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IReportGroup) GrantWrite(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IReportGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IReportGroup) ReportGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroup) ReportGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroup) ReportGroupRef() *interfacesawscodebuild.ReportGroupReference {
	var returns *interfacesawscodebuild.ReportGroupReference
	_jsii_.Get(
		j,
		"reportGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

