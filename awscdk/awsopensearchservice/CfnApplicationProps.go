package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AppConfigs: []interface{}{
//   		&AppConfigProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DataSources: []interface{}{
//   		&DataSourceProperty{
//   			DataSourceArn: jsii.String("dataSourceArn"),
//
//   			// the properties below are optional
//   			DataSourceDescription: jsii.String("dataSourceDescription"),
//   		},
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   	IamIdentityCenterOptions: &IamIdentityCenterOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		IamIdentityCenterInstanceArn: jsii.String("iamIdentityCenterInstanceArn"),
//   		IamRoleForIdentityCenterApplicationArn: jsii.String("iamRoleForIdentityCenterApplicationArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html
//
type CfnApplicationProps struct {
	// Name of an OpenSearch Application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of application configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-appconfigs
	//
	AppConfigs interface{} `field:"optional" json:"appConfigs" yaml:"appConfigs"`
	// List of data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-datasources
	//
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Endpoint URL of an OpenSearch Application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Container for IAM Identity Center Options settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-iamidentitycenteroptions
	//
	IamIdentityCenterOptions interface{} `field:"optional" json:"iamIdentityCenterOptions" yaml:"iamIdentityCenterOptions"`
	// An arbitrary set of tags (key-value pairs) for this application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

