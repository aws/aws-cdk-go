package awscertificatemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Represents a certificate in AWS Certificate Manager.
type ICertificate interface {
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

func (j *jsiiProxy_ICertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

