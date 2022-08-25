package awslogs


// Properties for a QueryString.
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
type QueryStringProps struct {
	// Specifies which fields to display in the query results.
	Display *string `field:"optional" json:"display" yaml:"display"`
	// Retrieves the specified fields from log events for display.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Filters the results of a query that's based on one or more conditions.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Specifies the number of log events returned by the query.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Extracts data from a log field and creates one or more ephemeral fields that you can process further in the query.
	Parse *string `field:"optional" json:"parse" yaml:"parse"`
	// Sorts the retrieved log events.
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Uses log field values to calculate aggregate statistics.
	Stats *string `field:"optional" json:"stats" yaml:"stats"`
}

