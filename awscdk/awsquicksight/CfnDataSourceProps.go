package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataSource`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html
//
type CfnDataSourceProps struct {
	// A display name for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the data source. To return a list of all data sources, use `ListDataSources` .
	//
	// Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the credentials from this existing data source. If the `AlternateDataSourceParameters` list is null, the `Credentials` originally used with this `DataSourceParameters` are automatically allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-alternatedatasourceparameters
	//
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The credentials Amazon Quick Sight that uses to connect to your underlying source.
	//
	// Currently, only credentials based on user name and password are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-credentials
	//
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// An ID for the data source.
	//
	// This ID is unique per AWS Region for each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceid
	//
	DataSourceId *string `field:"optional" json:"dataSourceId" yaml:"dataSourceId"`
	// The parameters that Amazon Quick Sight uses to connect to your underlying source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-datasourceparameters
	//
	DataSourceParameters interface{} `field:"optional" json:"dataSourceParameters" yaml:"dataSourceParameters"`
	// Error information from the last update or the creation of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-errorinfo
	//
	ErrorInfo interface{} `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-folderarns
	//
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// A list of resource permissions on the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon Quick Sight connects to your underlying source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-sslproperties
	//
	SslProperties interface{} `field:"optional" json:"sslProperties" yaml:"sslProperties"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Use this parameter only when you want Amazon Quick Sight to use a VPC connection when connecting to your underlying source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-datasource.html#cfn-quicksight-datasource-vpcconnectionproperties
	//
	VpcConnectionProperties interface{} `field:"optional" json:"vpcConnectionProperties" yaml:"vpcConnectionProperties"`
}

