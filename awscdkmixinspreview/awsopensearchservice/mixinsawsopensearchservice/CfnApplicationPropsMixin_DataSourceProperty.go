package mixinsawsopensearchservice


// Data sources that are associated with an OpenSearch application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceProperty := &DataSourceProperty{
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	DataSourceDescription: jsii.String("dataSourceDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-datasource.html
//
type CfnApplicationPropsMixin_DataSourceProperty struct {
	// Amazon Resource Name (ARN) format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-datasource.html#cfn-opensearchservice-application-datasource-datasourcearn
	//
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// Detailed description of a data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-application-datasource.html#cfn-opensearchservice-application-datasource-datasourcedescription
	//
	DataSourceDescription *string `field:"optional" json:"dataSourceDescription" yaml:"dataSourceDescription"`
}

