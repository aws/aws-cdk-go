package mixinsawsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	AppConfigs: []interface{}{
//   		&AppConfigProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DataSources: []interface{}{
//   		&DataSourceProperty{
//   			DataSourceArn: jsii.String("dataSourceArn"),
//   			DataSourceDescription: jsii.String("dataSourceDescription"),
//   		},
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   	IamIdentityCenterOptions: &IamIdentityCenterOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		IamIdentityCenterInstanceArn: jsii.String("iamIdentityCenterInstanceArn"),
//   		IamRoleForIdentityCenterApplicationArn: jsii.String("iamRoleForIdentityCenterApplicationArn"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html
//
type CfnApplicationMixinProps struct {
	// List of application configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-appconfigs
	//
	AppConfigs interface{} `field:"optional" json:"appConfigs" yaml:"appConfigs"`
	// List of data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-datasources
	//
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// The endpoint URL of an OpenSearch application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Settings container for integrating IAM Identity Center with OpenSearch UI applications, which enables enabling secure user authentication and access control across multiple data sources.
	//
	// This setup supports single sign-on (SSO) through IAM Identity Center, allowing centralized user management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-iamidentitycenteroptions
	//
	IamIdentityCenterOptions interface{} `field:"optional" json:"iamIdentityCenterOptions" yaml:"iamIdentityCenterOptions"`
	// The name of an OpenSearch application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An arbitrary set of tags (key-value pairs) for this application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-application.html#cfn-opensearchservice-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

