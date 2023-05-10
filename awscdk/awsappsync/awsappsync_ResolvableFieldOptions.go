package awsappsync


// Properties for configuring a resolvable field.
//
// Example:
//   var api graphqlApi
//   var filmNode objectType
//   var dummyRequest mappingTemplate
//   var dummyResponse mappingTemplate
//
//
//   string := appsync.GraphqlType_String()
//   int := appsync.GraphqlType_Int()
//   api.AddMutation(jsii.String("addFilm"), appsync.NewResolvableField(&ResolvableFieldOptions{
//   	ReturnType: filmNode.Attribute(),
//   	Args: map[string]graphqlType{
//   		"name": string,
//   		"film_number": int,
//   	},
//   	DataSource: api.AddNoneDataSource(jsii.String("none")),
//   	RequestMappingTemplate: dummyRequest,
//   	ResponseMappingTemplate: dummyResponse,
//   }))
//
// Experimental.
type ResolvableFieldOptions struct {
	// The return type for this field.
	// Experimental.
	ReturnType GraphqlType `field:"required" json:"returnType" yaml:"returnType"`
	// The arguments for this field.
	//
	// i.e. type Example (first: String second: String) {}
	// - where 'first' and 'second' are key values for args
	// and 'String' is the GraphqlType.
	// Experimental.
	Args *map[string]GraphqlType `field:"optional" json:"args" yaml:"args"`
	// the directives for this field.
	// Experimental.
	Directives *[]Directive `field:"optional" json:"directives" yaml:"directives"`
	// The data source creating linked to this resolvable field.
	// Experimental.
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
}

