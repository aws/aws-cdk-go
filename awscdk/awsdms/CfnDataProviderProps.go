package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataProviderProps := &CfnDataProviderProps{
//   	Engine: jsii.String("engine"),
//
//   	// the properties below are optional
//   	DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   	DataProviderName: jsii.String("dataProviderName"),
//   	Description: jsii.String("description"),
//   	ExactSettings: jsii.Boolean(false),
//   	Settings: &SettingsProperty{
//   		PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html
//
type CfnDataProviderProps struct {
	// The property describes a data engine for the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-engine
	//
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The property describes an identifier for the data provider.
	//
	// It is used for describing/deleting/modifying can be name/arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-dataprovideridentifier
	//
	DataProviderIdentifier *string `field:"optional" json:"dataProviderIdentifier" yaml:"dataProviderIdentifier"`
	// The property describes a name to identify the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-dataprovidername
	//
	DataProviderName *string `field:"optional" json:"dataProviderName" yaml:"dataProviderName"`
	// The optional description of the data provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The property describes the exact settings which can be modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-exactsettings
	//
	// Default: - false.
	//
	ExactSettings interface{} `field:"optional" json:"exactSettings" yaml:"exactSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html#cfn-dms-dataprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

