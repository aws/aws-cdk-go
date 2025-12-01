package awspcs


// Additional options related to the Slurm scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slurmConfigurationProperty := &SlurmConfigurationProperty{
//   	Accounting: &AccountingProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		DefaultPurgeTimeInDays: jsii.Number(123),
//   	},
//   	AuthKey: &AuthKeyProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		SecretVersion: jsii.String("secretVersion"),
//   	},
//   	JwtAuth: &JwtAuthProperty{
//   		JwtKey: &JwtKeyProperty{
//   			SecretArn: jsii.String("secretArn"),
//   			SecretVersion: jsii.String("secretVersion"),
//   		},
//   	},
//   	ScaleDownIdleTimeInSeconds: jsii.Number(123),
//   	SlurmCustomSettings: []interface{}{
//   		&SlurmCustomSettingProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	SlurmRest: &SlurmRestProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html
//
type CfnCluster_SlurmConfigurationProperty struct {
	// The accounting configuration includes configurable settings for Slurm accounting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html#cfn-pcs-cluster-slurmconfiguration-accounting
	//
	Accounting interface{} `field:"optional" json:"accounting" yaml:"accounting"`
	// The shared Slurm key for authentication, also known as the *cluster secret* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html#cfn-pcs-cluster-slurmconfiguration-authkey
	//
	AuthKey interface{} `field:"optional" json:"authKey" yaml:"authKey"`
	// The JWT authentication configuration for Slurm REST API access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html#cfn-pcs-cluster-slurmconfiguration-jwtauth
	//
	JwtAuth interface{} `field:"optional" json:"jwtAuth" yaml:"jwtAuth"`
	// The time (in seconds) before an idle node is scaled down.
	//
	// Default: `600`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html#cfn-pcs-cluster-slurmconfiguration-scaledownidletimeinseconds
	//
	ScaleDownIdleTimeInSeconds *float64 `field:"optional" json:"scaleDownIdleTimeInSeconds" yaml:"scaleDownIdleTimeInSeconds"`
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html#cfn-pcs-cluster-slurmconfiguration-slurmcustomsettings
	//
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
	// The Slurm REST API configuration for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html#cfn-pcs-cluster-slurmconfiguration-slurmrest
	//
	SlurmRest interface{} `field:"optional" json:"slurmRest" yaml:"slurmRest"`
}

