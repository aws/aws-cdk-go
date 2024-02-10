package awslogs


// Properties for defining a `CfnQueryDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueryDefinitionProps := &CfnQueryDefinitionProps{
//   	Name: jsii.String("name"),
//   	QueryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	LogGroupNames: []*string{
//   		jsii.String("logGroupNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html
//
type CfnQueryDefinitionProps struct {
	// A name for the query definition.
	//
	// > You can use the name to create a folder structure for your queries. To create a folder, use a forward slash (/) to prefix your desired query name with your desired folder name. For example, `/ *folder-name* / *query-name*` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query string to use for this query definition.
	//
	// For more information, see [CloudWatch Logs Insights Query Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Use this parameter if you want the query to query only certain log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-querydefinition.html#cfn-logs-querydefinition-loggroupnames
	//
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
}

