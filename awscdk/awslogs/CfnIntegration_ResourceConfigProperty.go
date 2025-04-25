package awslogs


// This structure contains configuration details about an integration between CloudWatch Logs and another entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceConfigProperty := &ResourceConfigProperty{
//   	OpenSearchResourceConfig: &OpenSearchResourceConfigProperty{
//   		DashboardViewerPrincipals: []*string{
//   			jsii.String("dashboardViewerPrincipals"),
//   		},
//   		DataSourceRoleArn: jsii.String("dataSourceRoleArn"),
//
//   		// the properties below are optional
//   		ApplicationArn: jsii.String("applicationArn"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		RetentionDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-resourceconfig.html
//
type CfnIntegration_ResourceConfigProperty struct {
	// This structure contains configuration details about an integration between CloudWatch Logs and OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-integration-resourceconfig.html#cfn-logs-integration-resourceconfig-opensearchresourceconfig
	//
	OpenSearchResourceConfig interface{} `field:"optional" json:"openSearchResourceConfig" yaml:"openSearchResourceConfig"`
}

