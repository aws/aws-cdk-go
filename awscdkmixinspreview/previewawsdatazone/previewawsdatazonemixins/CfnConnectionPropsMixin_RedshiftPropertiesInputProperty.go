package previewawsdatazonemixins


// The Amazon Redshift properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnConnectionPropsMixin_RedshiftPropertiesInputProperty struct {
	// The Amaon Redshift credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// The Amazon Redshift database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The Amazon Redshift host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The lineage sync of the Amazon Redshift.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-lineagesync
	//
	LineageSync interface{} `field:"optional" json:"lineageSync" yaml:"lineageSync"`
	// The Amaon Redshift port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The Amazon Redshift storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftpropertiesinput.html#cfn-datazone-connection-redshiftpropertiesinput-storage
	//
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
}

