package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
)

// An Aurora Cluster Instance.
type IAuroraClusterInstance interface {
	awscdk.IResource
	// The instance ARN.
	DbInstanceArn() *string
	// The instance endpoint address.
	DbInstanceEndpointAddress() *string
	// The instance resource ID.
	DbiResourceId() *string
	// The instance identifier.
	InstanceIdentifier() *string
	// The instance size if the instance is a provisioned type.
	InstanceSize() *string
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey() awskms.IKey
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention() PerformanceInsightRetention
	// Whether Performance Insights is enabled.
	PerformanceInsightsEnabled() *bool
	// The promotion tier the instance was created in.
	Tier() *float64
	// The instance type (provisioned vs serverless v2).
	Type() InstanceType
}

// The jsii proxy for IAuroraClusterInstance
type jsiiProxy_IAuroraClusterInstance struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAuroraClusterInstance) DbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) DbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) InstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) PerformanceInsightEncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"performanceInsightEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) PerformanceInsightRetention() PerformanceInsightRetention {
	var returns PerformanceInsightRetention
	_jsii_.Get(
		j,
		"performanceInsightRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) PerformanceInsightsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) Tier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuroraClusterInstance) Type() InstanceType {
	var returns InstanceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

