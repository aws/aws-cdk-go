package awslogs


// Properties for CfnQueryDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnQueryDefinitionMixinProps := &CfnQueryDefinitionMixinProps{
//   	LogGroupNames: []interface{}{
//   		jsii.String("logGroupNames"),
//   	},
//   	Name: jsii.String("name"),
//   	Parameters: []interface{}{
//   		&QueryParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	QueryLanguage: jsii.String("queryLanguage"),
//   	QueryString: jsii.String("queryString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html
//
type CfnQueryDefinitionMixinProps struct {
	// Use this parameter if you want the query to query only certain log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-loggroupnames
	//
	LogGroupNames *[]interface{} `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
	// A name for the query definition.
	//
	// > You can use the name to create a folder structure for your queries. To create a folder, use a forward slash (/) to prefix your desired query name with your desired folder name. For example, `*folder-name* / *query-name*` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Use this parameter to include specific query parameters as part of your query definition.
	//
	// Query parameters are supported only for Logs Insights QL queries. Query parameters allow you to use placeholder variables in your query string that are substituted with values at execution time. Use the {{parameterName}} syntax in your query string to reference a parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The query language used for this query.
	//
	// For more information about the query languages that CloudWatch Logs supports, see [Supported query languages](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_Languages.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-querylanguage
	//
	// Default: - "CWLI".
	//
	QueryLanguage *string `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// The query string to use for this query definition.
	//
	// For more information, see [CloudWatch Logs Insights Query Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

