package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRestApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var body interface{}
//   var policy interface{}
//
//   cfnRestApiProps := &cfnRestApiProps{
//   	apiKeySourceType: jsii.String("apiKeySourceType"),
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	body: body,
//   	bodyS3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		eTag: jsii.String("eTag"),
//   		key: jsii.String("key"),
//   		version: jsii.String("version"),
//   	},
//   	cloneFrom: jsii.String("cloneFrom"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	endpointConfiguration: &endpointConfigurationProperty{
//   		types: []*string{
//   			jsii.String("types"),
//   		},
//   		vpcEndpointIds: []*string{
//   			jsii.String("vpcEndpointIds"),
//   		},
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	minimumCompressionSize: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   	name: jsii.String("name"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	policy: policy,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRestApiProps struct {
	// The source of the API key for metering requests according to a usage plan. Valid values are:.
	//
	// - `HEADER` to read the API key from the `X-API-Key` header of a request.
	// - `AUTHORIZER` to read the API key from the `UsageIdentifierKey` from a Lambda authorizer.
	ApiKeySourceType *string `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media types that are supported by the `RestApi` resource.
	//
	// Use `~1` instead of `/` in the media types, for example `image~1png` or `application~1octet-stream` . By default, `RestApi` supports only UTF-8-encoded text payloads. Duplicates are not allowed. For more information, see [Enable Support for Binary Payloads in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-payload-encodings.html) in the *API Gateway Developer Guide* .
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// An OpenAPI specification that defines a set of RESTful APIs in JSON format.
	//
	// For YAML templates, you can also provide the specification in YAML format.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
	BodyS3Location interface{} `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// The ID of the `RestApi` resource that you want to clone.
	CloneFrom *string `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// A description of the `RestApi` resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
	//
	// By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the `Parameters` property.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the `RestApi` resource.
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API.
	//
	// When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// This property applies only when you use OpenAPI to define your REST API.
	//
	// The `Mode` determines how API Gateway handles resource updates.
	//
	// Valid values are `overwrite` or `merge` .
	//
	// For `overwrite` , the new API definition replaces the existing one. The existing API identifier remains unchanged.
	//
	// For `merge` , the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. Use `merge` to define top-level `RestApi` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
	//
	// If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is `overwrite` . Otherwise, the default value is `merge` .
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A name for the `RestApi` resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Custom header parameters for the request.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for the `RestApi` resource.
	//
	// To set the ARN for the policy, use the `!Join` intrinsic function with `""` as delimiter and values of `"execute-api:/"` and `"*"` .
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// An array of arbitrary tags (key-value pairs) to associate with the API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

