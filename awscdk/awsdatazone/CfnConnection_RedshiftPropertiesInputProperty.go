package awsdatazone


// Redshift Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftPropertiesInputProperty := &RedshiftPropertiesInputProperty{
//   	Credentials: &RedshiftCredentialsProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		UsernamePassword: &UsernamePasswordProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   	DatabaseName: jsii.String("databaseName"),
//   	Host: jsii.String("host"),
//   	LineageSync: &RedshiftLineageSyncConfigurationInputProperty{
//   		Enabled: jsii.Boolean(false),
//   		Schedule: &LineageSyncScheduleProperty{
//   			Schedule: jsii.String("schedule"),
//   		},
//   	},
//   	Port: jsii.Number(123),
//   	Storage: &RedshiftStoragePropertiesProperty{
//   		ClusterName: jsii.String("clusterName"),
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html
//
type CfnConnection_RedshiftPropertiesInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Redshift Lineage Sync Configuration Input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-lineagesync
	//
	LineageSync interface{} `field:"optional" json:"lineageSync" yaml:"lineageSync"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-storage
	//
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
}

