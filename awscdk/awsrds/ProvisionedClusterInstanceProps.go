package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Options for creating a provisioned instance.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"),
//   Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   	Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   }),
//   Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//    // Optional - will default to 'admin' username and generated password
//   Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_*Provisioned(jsii.String("reader1"), &ProvisionedClusterInstanceProps{
//   			PromotionTier: jsii.Number(1),
//   		}),
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader2")),
//   	},
//   	VpcSubnets: map[string]subnetType{
//   		"subnetType": ec2.*subnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	Vpc: *Vpc,
//   }),
//
type ProvisionedClusterInstanceProps struct {
	// Whether to allow upgrade of major version for the DB instance.
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The identifier for the database instance.
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The DB parameter group to associate with the instance.
	//
	// This is only needed if you need to configure different parameter
	// groups for each individual instance, otherwise you should not
	// provide this and just use the cluster parameter group.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The cluster instance type.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Only used for migrating existing clusters from using `instanceProps` to `writer` and `readers`.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   // existing cluster
	//   var vpc vpc
	//
	//   var vpc vpc
	//
	//   cluster := rds.NewDatabaseCluster(stack, jsii.String("Database"), &DatabaseClusterProps{
	//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
	//   		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	//   	}),
	//   	Instances: jsii.Number(2),
	//   	InstanceProps: &InstanceProps{
	//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
	//   		VpcSubnets: &SubnetSelection{
	//   			SubnetType: ec2.SubnetType_PUBLIC,
	//   		},
	//   		Vpc: *Vpc,
	//   	},
	//   })
	//
	//   // migration
	//
	//   instanceProps := map[string]interface{}{
	//   	"instanceType": ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
	//   	"isFromLegacyInstanceProps": jsii.Boolean(true),
	//   }
	//   cluster := rds.NewDatabaseCluster(stack, jsii.String("Database"), &DatabaseClusterProps{
	//   	Engine: rds.DatabaseClusterEngine_*AuroraMysql(&AuroraMysqlClusterEngineProps{
	//   		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	//   	}),
	//   	VpcSubnets: &SubnetSelection{
	//   		SubnetType: ec2.SubnetType_PUBLIC,
	//   	},
	//   	Vpc: Vpc,
	//   	Writer: clusterInstance_Provisioned(jsii.String("Instance1"), map[string]interface{}{
	//   		(SpreadAssignment ...instanceProps
	//   				instanceProps),
	//   	}),
	//   	Readers: []iClusterInstance{
	//   		*clusterInstance_*Provisioned(jsii.String("Instance2"), map[string]interface{}{
	//   			(SpreadAssignment ...instanceProps
	//   					instanceProps),
	//   		}),
	//   	},
	//   })
	//
	IsFromLegacyInstanceProps *bool `field:"optional" json:"isFromLegacyInstanceProps" yaml:"isFromLegacyInstanceProps"`
	// The promotion tier of the cluster instance.
	//
	// Can be between 0-15
	//
	// For provisioned instances this just determines the failover priority.
	// If multiple instances have the same priority then one will be picked at random.
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
}

