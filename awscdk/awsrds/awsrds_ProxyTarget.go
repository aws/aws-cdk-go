package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
type ProxyTarget interface {
	// Bind this target to the specified database proxy.
	Bind(proxy DatabaseProxy) *ProxyTargetConfig
}

// The jsii proxy struct for ProxyTarget
type jsiiProxy_ProxyTarget struct {
	_ byte // padding
}

// From cluster.
func ProxyTarget_FromCluster(cluster IDatabaseCluster) ProxyTarget {
	_init_.Initialize()

	var returns ProxyTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ProxyTarget",
		"fromCluster",
		[]interface{}{cluster},
		&returns,
	)

	return returns
}

// From instance.
func ProxyTarget_FromInstance(instance IDatabaseInstance) ProxyTarget {
	_init_.Initialize()

	var returns ProxyTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ProxyTarget",
		"fromInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProxyTarget) Bind(proxy DatabaseProxy) *ProxyTargetConfig {
	var returns *ProxyTargetConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{proxy},
		&returns,
	)

	return returns
}

