package awscertificatemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscertificatemanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a certificate in AWS Certificate Manager.
type ICertificate interface {
	interfacesawscertificatemanager.ICertificateRef
	awscdk.IResource
	// Return the DaysToExpiry metric for this AWS Certificate Manager Certificate. By default, this is the minimum value over 1 day.
	//
	// This metric is no longer emitted once the certificate has effectively
	// expired, so alarms configured on this metric should probably treat missing
	// data as "breaching".
	MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The certificate's ARN.
	CertificateArn() *string
}

// The jsii proxy for ICertificate
type jsiiProxy_ICertificate struct {
	internal.Type__interfacesawscertificatemanagerICertificateRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICertificate) MetricDaysToExpiry(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDaysToExpiryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDaysToExpiry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ICertificate) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificate) CertificateRef() *interfacesawscertificatemanager.CertificateReference {
	var returns *interfacesawscertificatemanager.CertificateReference
	_jsii_.Get(
		j,
		"certificateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificate) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

