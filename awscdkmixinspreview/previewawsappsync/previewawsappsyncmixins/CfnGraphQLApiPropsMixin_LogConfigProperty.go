package previewawsappsyncmixins


// The `LogConfig` property type specifies the logging configuration when writing GraphQL operations and tracing to Amazon CloudWatch for an AWS AppSync GraphQL API.
//
// `LogConfig` is a property of the [AWS::AppSync::GraphQLApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logConfigProperty := &LogConfigProperty{
//   	CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	ExcludeVerboseContent: jsii.Boolean(false),
//   	FieldLogLevel: jsii.String("fieldLogLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html
//
type CfnGraphQLApiPropsMixin_LogConfigProperty struct {
	// The service role that AWS AppSync will assume to publish to Amazon CloudWatch Logs in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html#cfn-appsync-graphqlapi-logconfig-cloudwatchlogsrolearn
	//
	CloudWatchLogsRoleArn *string `field:"optional" json:"cloudWatchLogsRoleArn" yaml:"cloudWatchLogsRoleArn"`
	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html#cfn-appsync-graphqlapi-logconfig-excludeverbosecontent
	//
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// The field logging level. Values can be NONE, ERROR, INFO, DEBUG, or ALL.
	//
	// - *NONE* : No field-level logs are captured.
	// - *ERROR* : Logs the following information *only* for the fields that are in the error category:
	//
	// - The error section in the server response.
	// - Field-level errors.
	// - The generated request/response functions that got resolved for error fields.
	// - *INFO* : Logs the following information *only* for the fields that are in the info and error categories:
	//
	// - Info-level messages.
	// - The user messages sent through `$util.log.info` and `console.log` .
	// - Field-level tracing and mapping logs are not shown.
	// - *DEBUG* : Logs the following information *only* for the fields that are in the debug, info, and error categories:
	//
	// - Debug-level messages.
	// - The user messages sent through `$util.log.info` , `$util.log.debug` , `console.log` , and `console.debug` .
	// - Field-level tracing and mapping logs are not shown.
	// - *ALL* : The following information is logged for all fields in the query:
	//
	// - Field-level tracing information.
	// - The generated request/response functions that were resolved for each field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-logconfig.html#cfn-appsync-graphqlapi-logconfig-fieldloglevel
	//
	FieldLogLevel *string `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
}

