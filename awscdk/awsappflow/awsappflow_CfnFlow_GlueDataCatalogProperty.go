package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueDataCatalogProperty := &GlueDataCatalogProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	RoleArn: jsii.String("roleArn"),
//   	TablePrefix: jsii.String("tablePrefix"),
//   }
//
type CfnFlow_GlueDataCatalogProperty struct {
	// `CfnFlow.GlueDataCatalogProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnFlow.GlueDataCatalogProperty.RoleArn`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnFlow.GlueDataCatalogProperty.TablePrefix`.
	TablePrefix *string `field:"required" json:"tablePrefix" yaml:"tablePrefix"`
}

