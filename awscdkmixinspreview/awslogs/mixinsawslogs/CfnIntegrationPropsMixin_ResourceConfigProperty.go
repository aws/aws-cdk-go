package mixinsawslogs


// This structure contains configuration details about an integration between CloudWatch Logs and another entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceConfigProperty := &ResourceConfigProperty{
//   	OpenSearchResourceConfig: &OpenSearchResourceConfigProperty{
//   		ApplicationArn: jsii.String("applicationArn"),
//   		DashboardViewerPrincipals: []*string{
//   			jsii.String("dashboardViewerPrincipals"),
//   		},
//   		DataSourceRoleArn: jsii.String("dataSourceRoleArn"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		RetentionDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-resourceconfig.html
//
type CfnIntegrationPropsMixin_ResourceConfigProperty struct {
	// This structure contains configuration details about an integration between CloudWatch Logs and OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-resourceconfig.html#cfn-logs-integration-resourceconfig-opensearchresourceconfig
	//
	OpenSearchResourceConfig interface{} `field:"optional" json:"openSearchResourceConfig" yaml:"openSearchResourceConfig"`
}

