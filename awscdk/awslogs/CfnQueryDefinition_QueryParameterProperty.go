package awslogs


// This structure defines a query parameter for a saved CloudWatch Logs Insights query definition.
//
// Query parameters are supported only for Logs Insights QL queries. They are placeholder variables that you can reference in a query string using the {{parameterName}} syntax. Each parameter can include a default value and a description.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryParameterProperty := &QueryParameterProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultValue: jsii.String("defaultValue"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-querydefinition-queryparameter.html
//
type CfnQueryDefinition_QueryParameterProperty struct {
	// The name of the query parameter.
	//
	// A query parameter name must start with a letter or underscore, and contain only letters, digits, and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-querydefinition-queryparameter.html#cfn-logs-querydefinition-queryparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The default value to use for this query parameter if no value is supplied at execution time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-querydefinition-queryparameter.html#cfn-logs-querydefinition-queryparameter-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A description of the query parameter that explains its purpose or expected values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-querydefinition-queryparameter.html#cfn-logs-querydefinition-queryparameter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

