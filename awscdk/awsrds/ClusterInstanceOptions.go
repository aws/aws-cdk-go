package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Common options for creating a cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var caCertificate caCertificate
//   var key key
//   var parameterGroup parameterGroup
//
//   clusterInstanceOptions := &ClusterInstanceOptions{
//   	AllowMajorVersionUpgrade: jsii.Boolean(false),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	CaCertificate: caCertificate,
//   	EnablePerformanceInsights: jsii.Boolean(false),
//   	InstanceIdentifier: jsii.String("instanceIdentifier"),
//   	IsFromLegacyInstanceProps: jsii.Boolean(false),
//   	ParameterGroup: parameterGroup,
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	PerformanceInsightEncryptionKey: key,
//   	PerformanceInsightRetention: awscdk.Aws_rds.PerformanceInsightRetention_DEFAULT,
//   	PubliclyAccessible: jsii.Boolean(false),
//   }
//
type ClusterInstanceOptions struct {
	// Whether to allow upgrade of major version for the DB instance.
	// Default: - false.
	//
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	// Default: - true.
	//
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The identifier of the CA certificate for this DB cluster's instances.
	//
	// Specifying or updating this property triggers a reboot.
	//
	// For RDS DB engines:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html
	//
	// Default: - RDS will choose a certificate authority.
	//
	CaCertificate CaCertificate `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// Whether to enable Performance Insights for the DB instance.
	// Default: - false, unless ``performanceInsightRetention`` or ``performanceInsightEncryptionKey`` is set.
	//
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The identifier for the database instance.
	// Default: - CloudFormation generated identifier.
	//
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// Only used for migrating existing clusters from using `instanceProps` to `writer` and `readers`.
	//
	// Example:
	//   // existing cluster
	//   var vpc vpc
	//
	//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
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
	//
	//   myCluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
	//   	Engine: rds.DatabaseClusterEngine_*AuroraMysql(&AuroraMysqlClusterEngineProps{
	//   		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
	//   	}),
	//   	VpcSubnets: &SubnetSelection{
	//   		SubnetType: ec2.SubnetType_PUBLIC,
	//   	},
	//   	Vpc: Vpc,
	//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("Instance1"), &ProvisionedClusterInstanceProps{
	//   		InstanceType: instanceProps.instanceType,
	//   		IsFromLegacyInstanceProps: instanceProps.isFromLegacyInstanceProps,
	//   	}),
	//   	Readers: []iClusterInstance{
	//   		rds.ClusterInstance_*Provisioned(jsii.String("Instance2"), &ProvisionedClusterInstanceProps{
	//   			InstanceType: instanceProps.instanceType,
	//   			IsFromLegacyInstanceProps: instanceProps.isFromLegacyInstanceProps,
	//   		}),
	//   	},
	//   })
	//
	// Default: false.
	//
	IsFromLegacyInstanceProps *bool `field:"optional" json:"isFromLegacyInstanceProps" yaml:"isFromLegacyInstanceProps"`
	// The DB parameter group to associate with the instance.
	//
	// This is only needed if you need to configure different parameter
	// groups for each individual instance, otherwise you should not
	// provide this and just use the cluster parameter group.
	// Default: the cluster parameter group is used.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	// Default: - None.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Default: - default master key.
	//
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Default: 7.
	//
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// Indicates whether the DB instance is an internet-facing instance.
	//
	// If not specified,
	// the cluster's vpcSubnets will be used to determine if the instance is internet-facing
	// or not.
	// Default: - `true` if the cluster's `vpcSubnets` is `subnetType: SubnetType.PUBLIC`, `false` otherwise
	//
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
}

