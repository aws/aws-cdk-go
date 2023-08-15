package awslogs


// Properties for a QueryString.
//
// Example:
//   logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &QueryDefinitionProps{
//   	QueryDefinitionName: jsii.String("MyQuery"),
//   	QueryString: logs.NewQueryString(&QueryStringProps{
//   		Fields: []*string{
//   			jsii.String("@timestamp"),
//   			jsii.String("@message"),
//   		},
//   		ParseStatements: []*string{
//   			jsii.String("@message \"[*] *\" as loggingType, loggingMessage"),
//   			jsii.String("@message \"<*>: *\" as differentLoggingType, differentLoggingMessage"),
//   		},
//   		FilterStatements: []*string{
//   			jsii.String("loggingType = \"ERROR\""),
//   			jsii.String("loggingMessage = \"A very strange error occurred!\""),
//   		},
//   		Sort: jsii.String("@timestamp desc"),
//   		Limit: jsii.Number(20),
//   	}),
//   })
//
type QueryStringProps struct {
	// Specifies which fields to display in the query results.
	// Default: - no display in QueryString.
	//
	Display *string `field:"optional" json:"display" yaml:"display"`
	// Retrieves the specified fields from log events for display.
	// Default: - no fields in QueryString.
	//
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// A single statement for filtering the results of a query based on a boolean expression.
	// Default: - no filter in QueryString.
	//
	// Deprecated: Use `filterStatements` instead.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// An array of one or more statements for filtering the results of a query based on a boolean expression.
	//
	// Each provided statement generates a separate filter line in the query string.
	//
	// Note: If provided, this property overrides any value provided for the `filter` property.
	// Default: - no filter in QueryString.
	//
	FilterStatements *[]*string `field:"optional" json:"filterStatements" yaml:"filterStatements"`
	// Specifies the number of log events returned by the query.
	// Default: - no limit in QueryString.
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// A single statement for parsing data from a log field and creating ephemeral fields that can be processed further in the query.
	// Default: - no parse in QueryString.
	//
	// Deprecated: Use `parseStatements` instead.
	Parse *string `field:"optional" json:"parse" yaml:"parse"`
	// An array of one or more statements for parsing data from a log field and creating ephemeral fields that can be processed further in the query.
	//
	// Each provided statement generates a separate
	// parse line in the query string.
	//
	// Note: If provided, this property overrides any value provided for the `parse` property.
	// Default: - no parse in QueryString.
	//
	ParseStatements *[]*string `field:"optional" json:"parseStatements" yaml:"parseStatements"`
	// Sorts the retrieved log events.
	// Default: - no sort in QueryString.
	//
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// Uses log field values to calculate aggregate statistics.
	// Default: - no stats in QueryString.
	//
	Stats *string `field:"optional" json:"stats" yaml:"stats"`
}

