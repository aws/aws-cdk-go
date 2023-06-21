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
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
//   	Vpc: Vpc,
//   })
//
//   proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &DatabaseProxyProps{
//   	ProxyTarget: rds.ProxyTarget_FromCluster(cluster),
//   	Secrets: []iSecret{
//   		cluster.Secret,
//   	},
//   	Vpc: Vpc,
//   })
//
//   role := iam.NewRole(this, jsii.String("DBProxyRole"), &RoleProps{
//   	AssumedBy: iam.NewAccountPrincipal(this.Account),
//   })
//   proxy.GrantConnect(role, jsii.String("admin"))
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

	if err := validateProxyTarget_FromClusterParameters(cluster); err != nil {
		panic(err)
	}
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

	if err := validateProxyTarget_FromInstanceParameters(instance); err != nil {
		panic(err)
	}
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

