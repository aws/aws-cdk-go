package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Proxy target: Instance or Cluster.
//
// A target group is a collection of databases that the proxy can connect to.
// Currently, you can specify only one RDS DB instance or Aurora DB cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine_AURORA(),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   	},
//   })
//
//   proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &databaseProxyProps{
//   	proxyTarget: rds.proxyTarget.fromCluster(cluster),
//   	secrets: []iSecret{
//   		cluster.secret,
//   	},
//   	vpc: vpc,
//   })
//
//   role := iam.NewRole(this, jsii.String("DBProxyRole"), &roleProps{
//   	assumedBy: iam.NewAccountPrincipal(this.account),
//   })
//   proxy.grantConnect(role, jsii.String("admin"))
//
// Experimental.
type ProxyTarget interface {
	// Bind this target to the specified database proxy.
	// Experimental.
	Bind(proxy DatabaseProxy) *ProxyTargetConfig
}

// The jsii proxy struct for ProxyTarget
type jsiiProxy_ProxyTarget struct {
	_ byte // padding
}

// From cluster.
// Experimental.
func ProxyTarget_FromCluster(cluster IDatabaseCluster) ProxyTarget {
	_init_.Initialize()

	if err := validateProxyTarget_FromClusterParameters(cluster); err != nil {
		panic(err)
	}
	var returns ProxyTarget

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.ProxyTarget",
		"fromCluster",
		[]interface{}{cluster},
		&returns,
	)

	return returns
}

// From instance.
// Experimental.
func ProxyTarget_FromInstance(instance IDatabaseInstance) ProxyTarget {
	_init_.Initialize()

	if err := validateProxyTarget_FromInstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns ProxyTarget

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.ProxyTarget",
		"fromInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyTarget) Bind(proxy DatabaseProxy) *ProxyTargetConfig {
	if err := p.validateBindParameters(proxy); err != nil {
		panic(err)
	}
	var returns *ProxyTargetConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{proxy},
		&returns,
	)

	return returns
}

