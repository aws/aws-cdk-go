package awslogs


// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationProps := &CfnIntegrationProps{
//   	IntegrationName: jsii.String("integrationName"),
//   	IntegrationType: jsii.String("integrationType"),
//   	ResourceConfig: &ResourceConfigProperty{
//   		OpenSearchResourceConfig: &OpenSearchResourceConfigProperty{
//   			DashboardViewerPrincipals: []*string{
//   				jsii.String("dashboardViewerPrincipals"),
//   			},
//   			DataSourceRoleArn: jsii.String("dataSourceRoleArn"),
//
//   			// the properties below are optional
//   			ApplicationArn: jsii.String("applicationArn"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			RetentionDays: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-integration.html
//
type CfnIntegrationProps struct {
	// The name of this integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-integration.html#cfn-logs-integration-integrationname
	//
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The type of integration.
	//
	// Integrations with OpenSearch Service have the type `OPENSEARCH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-integration.html#cfn-logs-integration-integrationtype
	//
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// This structure contains configuration details about an integration between CloudWatch Logs and another entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-integration.html#cfn-logs-integration-resourceconfig
	//
	ResourceConfig interface{} `field:"required" json:"resourceConfig" yaml:"resourceConfig"`
}

