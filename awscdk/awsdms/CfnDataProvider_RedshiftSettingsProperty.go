package awsdms


// Provides information that defines an Amazon Redshift endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftSettingsProperty := &RedshiftSettingsProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	Port: jsii.Number(123),
//   	ServerName: jsii.String("serverName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-redshiftsettings.html
//
type CfnDataProvider_RedshiftSettingsProperty struct {
	// The name of the Amazon Redshift data warehouse (service) that you are working with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-redshiftsettings.html#cfn-dms-dataprovider-redshiftsettings-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port number for Amazon Redshift.
	//
	// The default value is 5439.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-redshiftsettings.html#cfn-dms-dataprovider-redshiftsettings-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The name of the Amazon Redshift cluster you are using.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-redshiftsettings.html#cfn-dms-dataprovider-redshiftsettings-servername
	//
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
}

