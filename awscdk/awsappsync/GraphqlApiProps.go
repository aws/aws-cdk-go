package awsappsync


// Properties for an AppSync GraphQL API.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   api := appsync.NewGraphqlApi(stack, jsii.String("EventBridgeApi"), &GraphqlApiProps{
//   	Name: jsii.String("EventBridgeApi"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.eventbridge.graphql"))),
//   })
//
//   bus := events.NewEventBus(stack, jsii.String("DestinationEventBus"), map[string]interface{}{
//   })
//
//   dataSource := api.AddEventBridgeDataSource(jsii.String("NoneDS"), bus)
//
//   dataSource.CreateResolver(jsii.String("EventResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("emitEvent"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type GraphqlApiProps struct {
	// the name of the GraphQL API.
	Name *string `field:"required" json:"name" yaml:"name"`
	// GraphQL schema definition. Specify how you want to define your schema.
	//
	// Schema.fromFile(filePath: string) allows schema definition through schema.graphql file
	Schema ISchema `field:"required" json:"schema" yaml:"schema"`
	// Optional authorization configuration.
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The domain name configuration for the GraphQL API.
	//
	// The Route 53 hosted zone and CName DNS record must be configured in addition to this setting to
	// enable custom domain URL.
	DomainName *DomainOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Logging configuration for this api.
	LogConfig *LogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	XrayEnabled *bool `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

