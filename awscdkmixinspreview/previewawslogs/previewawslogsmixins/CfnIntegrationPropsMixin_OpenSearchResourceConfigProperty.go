package previewawslogsmixins


// This structure contains configuration details about an integration between CloudWatch Logs and OpenSearch Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openSearchResourceConfigProperty := &OpenSearchResourceConfigProperty{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	DashboardViewerPrincipals: []*string{
//   		jsii.String("dashboardViewerPrincipals"),
//   	},
//   	DataSourceRoleArn: jsii.String("dataSourceRoleArn"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	RetentionDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-opensearchresourceconfig.html
//
type CfnIntegrationPropsMixin_OpenSearchResourceConfigProperty struct {
	// If you want to use an existing OpenSearch Service application for your integration with OpenSearch Service, specify it here.
	//
	// If you omit this, a new application will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-opensearchresourceconfig.html#cfn-logs-integration-opensearchresourceconfig-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// Specify the ARNs of IAM roles and IAM users who you want to grant permission to for viewing the dashboards.
	//
	// > In addition to specifying these users here, you must also grant them the *CloudWatchOpenSearchDashboardAccess* IAM policy. For more information, see [IAM policies for users](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/OpenSearch-Dashboards-UserRoles.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-opensearchresourceconfig.html#cfn-logs-integration-opensearchresourceconfig-dashboardviewerprincipals
	//
	DashboardViewerPrincipals *[]*string `field:"optional" json:"dashboardViewerPrincipals" yaml:"dashboardViewerPrincipals"`
	// Specify the ARN of an IAM role that CloudWatch Logs will use to create the integration.
	//
	// This role must have the permissions necessary to access the OpenSearch Service collection to be able to create the dashboards. For more information about the permissions needed, see [Permissions that the integration needs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/OpenSearch-Dashboards-CreateRole.html) in the CloudWatch Logs User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-opensearchresourceconfig.html#cfn-logs-integration-opensearchresourceconfig-datasourcerolearn
	//
	DataSourceRoleArn *string `field:"optional" json:"dataSourceRoleArn" yaml:"dataSourceRoleArn"`
	// To have the vended dashboard data encrypted with AWS  instead of the CloudWatch Logs default encryption method, specify the ARN of the AWS  key that you want to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-opensearchresourceconfig.html#cfn-logs-integration-opensearchresourceconfig-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Specify how many days that you want the data derived by OpenSearch Service to be retained in the index that the dashboard refers to.
	//
	// This also sets the maximum time period that you can choose when viewing data in the dashboard. Choosing a longer time frame will incur additional costs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-opensearchresourceconfig.html#cfn-logs-integration-opensearchresourceconfig-retentiondays
	//
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

