package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

type IVpnConnection interface {
	awscdk.IResource
	interfacesawsec2.IVPNConnectionRef
	// Return the given named metric for this VPNConnection.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The bytes received through the VPN tunnel.
	//
	// Sum over 5 minutes.
	MetricTunnelDataIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The bytes sent through the VPN tunnel.
	//
	// Sum over 5 minutes.
	MetricTunnelDataOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The state of the tunnel. 0 indicates DOWN and 1 indicates UP.
	//
	// Average over 5 minutes.
	MetricTunnelState(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ASN of the customer gateway.
	CustomerGatewayAsn() *float64
	// The id of the customer gateway.
	CustomerGatewayId() *string
	// The ip address of the customer gateway.
	CustomerGatewayIp() *string
	// The id of the VPN connection.
	VpnId() *string
}

// The jsii proxy for IVpnConnection
type jsiiProxy_IVpnConnection struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsec2IVPNConnectionRef
}

func (i *jsiiProxy_IVpnConnection) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpnConnection) MetricTunnelDataIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTunnelDataInParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTunnelDataIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpnConnection) MetricTunnelDataOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTunnelDataOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTunnelDataOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpnConnection) MetricTunnelState(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTunnelStateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTunnelState",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVpnConnection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IVpnConnection) CustomerGatewayAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerGatewayAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) CustomerGatewayIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) VpnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpnConnection) VpnConnectionRef() *interfacesawsec2.VPNConnectionReference {
	var returns *interfacesawsec2.VPNConnectionReference
	_jsii_.Get(
		j,
		"vpnConnectionRef",
		&returns,
	)
	return returns
}

