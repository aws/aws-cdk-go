package awsappsync


// Properties for an AppSync GraphQL API.
//
// Example:
//   sourceApi := appsync.NewGraphqlApi(this, jsii.String("FirstSourceAPI"), &GraphqlApiProps{
//   	Name: jsii.String("FirstSourceAPI"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.merged-api-1.graphql"))),
//   })
//
//   importedMergedApi := appsync.GraphqlApi_FromGraphqlApiAttributes(this, jsii.String("ImportedMergedApi"), &GraphqlApiAttributes{
//   	GraphqlApiId: jsii.String("MyApiId"),
//   	GraphqlApiArn: jsii.String("MyApiArn"),
//   })
//
//   importedExecutionRole := iam.Role_FromRoleArn(this, jsii.String("ExecutionRole"), jsii.String("arn:aws:iam::ACCOUNT:role/MyExistingRole"))
//   appsync.NewSourceApiAssociation(this, jsii.String("SourceApiAssociation2"), &SourceApiAssociationProps{
//   	SourceApi: sourceApi,
//   	MergedApi: importedMergedApi,
//   	MergeType: appsync.MergeType_MANUAL_MERGE,
//   	MergedApiExecutionRole: importedExecutionRole,
//   })
//
type GraphqlApiProps struct {
	// the name of the GraphQL API.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional authorization configuration.
	// Default: - API Key authorization.
	//
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Definition (schema file or source APIs) for this GraphQL Api.
	Definition Definition `field:"optional" json:"definition" yaml:"definition"`
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
	// GraphQL schema definition. Specify how you want to define your schema.
	//
	// SchemaFile.fromAsset(filePath: string) allows schema definition through schema.graphql file
	// Default: - schema will be generated code-first (i.e. addType, addObjectType, etc.)
	//
	// Deprecated: use Definition.schema instead
	Schema ISchema `field:"optional" json:"schema" yaml:"schema"`
	// A value indicating whether the API is accessible from anywhere (GLOBAL) or can only be access from a VPC (PRIVATE).
	// Default: - GLOBAL.
	//
	Visibility Visibility `field:"optional" json:"visibility" yaml:"visibility"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	// Default: - false.
	//
	XrayEnabled *bool `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

