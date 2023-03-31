package awslogs


// Properties for defining a `CfnQueryDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueryDefinitionProps := &cfnQueryDefinitionProps{
//   	name: jsii.String("name"),
//   	queryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	logGroupNames: []*string{
//   		jsii.String("logGroupNames"),
//   	},
//   }
//
type CfnQueryDefinitionProps struct {
	// A name for the query definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query string to use for this query definition.
	//
	// For more information, see [CloudWatch Logs Insights Query Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html) .
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Use this parameter if you want the query to query only certain log groups.
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
}

