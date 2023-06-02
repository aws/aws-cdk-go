package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The interface representing the ReportGroup resource - either an existing one, imported using the `ReportGroup.fromReportGroupName` method, or a new one, created with the `ReportGroup` class.
type IReportGroup interface {
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

