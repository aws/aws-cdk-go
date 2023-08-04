package awslogs


// Properties for a QueryDefinition.
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
type QueryDefinitionProps struct {
	// Name of the query definition.
	QueryDefinitionName *string `field:"required" json:"queryDefinitionName" yaml:"queryDefinitionName"`
	// The query string to use for this query definition.
	QueryString QueryString `field:"required" json:"queryString" yaml:"queryString"`
	// Specify certain log groups for the query definition.
	// Default: - no specified log groups.
	//
	LogGroups *[]ILogGroup `field:"optional" json:"logGroups" yaml:"logGroups"`
}

