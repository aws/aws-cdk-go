package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Create an RDS Aurora Cluster Instance.
//
// You can create either provisioned or
// serverless v2 instances.
//
// Example:
//   var vpc vpc
//
//   myCluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R6G, ec2.InstanceSize_XLARGE4),
//   	}),
//   	ServerlessV2MinCapacity: jsii.Number(6.5),
//   	ServerlessV2MaxCapacity: jsii.Number(64),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader1"), &ServerlessV2ClusterInstanceProps{
//   			ScaleWithWriter: jsii.Boolean(true),
//   		}),
//   		rds.ClusterInstance_*ServerlessV2(jsii.String("reader2")),
//   	},
//   	Vpc: Vpc,
//   })
//
type ClusterInstance interface {
	IClusterInstance
	// Add the ClusterInstance to the cluster.
	Bind(scope constructs.Construct, cluster IDatabaseCluster, props *ClusterInstanceBindOptions) IAuroraClusterInstance
}

// The jsii proxy struct for ClusterInstance
type jsiiProxy_ClusterInstance struct {
	jsiiProxy_IClusterInstance
}

// Add a provisioned instance to the cluster.
//
// Example:
//   rds.ClusterInstance_Provisioned(jsii.String("ClusterInstance"), &ProvisionedClusterInstanceProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R6G, ec2.InstanceSize_XLARGE4),
//   })
//
func ClusterInstance_Provisioned(id *string, props *ProvisionedClusterInstanceProps) IClusterInstance {
	_init_.Initialize()

	if err := validateClusterInstance_ProvisionedParameters(id, props); err != nil {
		panic(err)
	}
	var returns IClusterInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ClusterInstance",
		"provisioned",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Add a serverless v2 instance to the cluster.
//
// Example:
//   rds.ClusterInstance_ServerlessV2(jsii.String("ClusterInstance"), &ServerlessV2ClusterInstanceProps{
//   	ScaleWithWriter: jsii.Boolean(true),
//   })
//
func ClusterInstance_ServerlessV2(id *string, props *ServerlessV2ClusterInstanceProps) IClusterInstance {
	_init_.Initialize()

	if err := validateClusterInstance_ServerlessV2Parameters(id, props); err != nil {
		panic(err)
	}
	var returns IClusterInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ClusterInstance",
		"serverlessV2",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterInstance) Bind(scope constructs.Construct, cluster IDatabaseCluster, props *ClusterInstanceBindOptions) IAuroraClusterInstance {
	if err := c.validateBindParameters(scope, cluster, props); err != nil {
		panic(err)
	}
	var returns IAuroraClusterInstance

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, cluster, props},
		&returns,
	)

	return returns
}

