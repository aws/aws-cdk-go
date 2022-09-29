package awslogs


// Properties for a QueryDefinition.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &queryDefinitionProps{
//   	queryDefinitionName: jsii.String("MyQuery"),
//   	queryString: logs.NewQueryString(&queryStringProps{
//   		fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		sort: jsii.String("@timestamp desc"),
//   		limit: jsii.Number(20),
//   	}),
//   })
//
// Experimental.
type QueryDefinitionProps struct {
	// Name of the query definition.
	// Experimental.
	QueryDefinitionName *string `field:"required" json:"queryDefinitionName" yaml:"queryDefinitionName"`
	// The query string to use for this query definition.
	// Experimental.
	QueryString QueryString `field:"required" json:"queryString" yaml:"queryString"`
	// Specify certain log groups for the query definition.
	// Experimental.
	LogGroups *[]ILogGroup `field:"optional" json:"logGroups" yaml:"logGroups"`
}

