package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Collection of grant methods for a ILogGroupRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroupRef ILogGroupRef
//
//   logGroupGrants := awscdk.Aws_logs.LogGroupGrants_FromLogGroup(logGroupRef)
//
type LogGroupGrants interface {
	PolicyResource() awsiam.IResourceWithPolicyV2
	Resource() interfacesawslogs.ILogGroupRef
	// Give permissions to read and filter events from this log group.
	Read(grantee awsiam.IGrantable) awsiam.Grant
	// Give permissions to create and write to streams in this log group.
	Write(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for LogGroupGrants
type jsiiProxy_LogGroupGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_LogGroupGrants) PolicyResource() awsiam.IResourceWithPolicyV2 {
	var returns awsiam.IResourceWithPolicyV2
	_jsii_.Get(
		j,
		"policyResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroupGrants) Resource() interfacesawslogs.ILogGroupRef {
	var returns interfacesawslogs.ILogGroupRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for LogGroupGrants.
func LogGroupGrants_FromLogGroup(resource interfacesawslogs.ILogGroupRef) LogGroupGrants {
	_init_.Initialize()

	if err := validateLogGroupGrants_FromLogGroupParameters(resource); err != nil {
		panic(err)
	}
	var returns LogGroupGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogGroupGrants",
		"fromLogGroup",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroupGrants) Read(grantee awsiam.IGrantable) awsiam.Grant {
	if err := l.validateReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"read",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroupGrants) Write(grantee awsiam.IGrantable) awsiam.Grant {
	if err := l.validateWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"write",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

