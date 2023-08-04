package awsappsync


// Properties for an AppSync GraphQL API.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//
//   api := appsync.NewGraphqlApi(this, jsii.String("EventBridgeApi"), &GraphqlApiProps{
//   	Name: jsii.String("EventBridgeApi"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.eventbridge.graphql"))),
//   })
//
//   bus := events.NewEventBus(this, jsii.String("DestinationEventBus"), &EventBusProps{
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
	// SchemaFile.fromAsset(filePath: string) allows schema definition through schema.graphql file
	// Default: - schema will be generated code-first (i.e. addType, addObjectType, etc.)
	//
	Schema ISchema `field:"required" json:"schema" yaml:"schema"`
	// Optional authorization configuration.
	// Default: - API Key authorization.
	//
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The domain name configuration for the GraphQL API.
	//
	// The Route 53 hosted zone and CName DNS record must be configured in addition to this setting to
	// enable custom domain URL.
	// Default: - no domain name.
	//
	DomainName *DomainOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Logging configuration for this api.
	// Default: - None.
	//
	LogConfig *LogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// A value indicating whether the API is accessible from anywhere (GLOBAL) or can only be access from a VPC (PRIVATE).
	// Default: - GLOBAL.
	//
	Visibility Visibility `field:"optional" json:"visibility" yaml:"visibility"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	// Default: - false.
	//
	XrayEnabled *bool `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

