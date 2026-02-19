package cloudassemblyschema


// Query input for lookup up CloudFormation resources using CC API.
//
// The example below is required to successfully compile CDK (otherwise,
// the CDK build will generate a synthetic example for the below, but it
// doesn't have enough type information about the literal string union
// to generate a validly compiling example).
//
// Example:
//   import "github.com/cdklabs/cloud-assembly-schema-go/awscdkcloudassemblyschema"
//
//
//   x := &CcApiContextQuery{
//   	TypeName: jsii.String("AWS::Some::Type"),
//   	ExpectedMatchCount: jsii.String("exactly-one"),
//   	ResourceModel: map[string]interface{}{
//   		"SomeArn": jsii.String("arn:aws:...."),
//   	},
//   	PropertiesToReturn: []*string{
//   		jsii.String("SomeProp"),
//   	},
//   	Account: jsii.String("11111111111"),
//   	Region: jsii.String("us-east-1"),
//   }
//
type CcApiContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Additional options to pass to STS when assuming the lookup role.
	//
	// - `RoleArn` should not be used. Use the dedicated `lookupRoleArn` property instead.
	// - `ExternalId` should not be used. Use the dedicated `lookupRoleExternalId` instead.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/STS.html#assumeRole-property
	//
	// Default: - No additional options.
	//
	AssumeRoleAdditionalOptions *map[string]interface{} `field:"optional" json:"assumeRoleAdditionalOptions" yaml:"assumeRoleAdditionalOptions"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// The ExternalId that needs to be supplied while assuming this role.
	// Default: - No ExternalId will be supplied.
	//
	LookupRoleExternalId *string `field:"optional" json:"lookupRoleExternalId" yaml:"lookupRoleExternalId"`
	// This is a set of properties returned from CC API that we want to return from ContextQuery.
	//
	// If any properties listed here are absent from the target resource, an error will be thrown.
	//
	// The returned object will always include the key `Identifier` with the CC-API returned
	// field `Identifier`.
	//
	// ## Notes on property completeness
	//
	// CloudControl API's `ListResources` may return fewer properties than
	// `GetResource` would, depending on the resource implementation.
	//
	// The returned properties here are *currently* selected from the response
	// object that CloudControl API returns to the CDK CLI.
	//
	// However, if we find there is need to do so, we may decide to change this
	// behavior in the future: we might change it to perform an additional
	// `GetResource` call for resources matched by `propertyMatch`.
	PropertiesToReturn *[]*string `field:"required" json:"propertiesToReturn" yaml:"propertiesToReturn"`
	// The CloudFormation resource type.
	//
	// See https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/supported-resources.html
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The value to return if the resource was not found and `ignoreErrorOnMissingContext` is true.
	//
	// If supplied, `dummyValue` should be an array of objects.
	//
	// `dummyValue` does not have to have elements, and it may have objects with
	// different properties than the properties in `propertiesToReturn`, but it
	// will be easiest for downstream code if the `dummyValue` conforms to
	// the expected response shape.
	// Default: - No dummy value available.
	//
	DummyValue interface{} `field:"optional" json:"dummyValue" yaml:"dummyValue"`
	// Identifier of the resource to look up using `GetResource`.
	//
	// Specifying exactIdentifier will return exactly one result, or throw an error
	// unless `ignoreErrorOnMissingContext` is set.
	// Default: - Either exactIdentifier or propertyMatch should be specified.
	//
	ExactIdentifier *string `field:"optional" json:"exactIdentifier" yaml:"exactIdentifier"`
	// Expected count of results if `propertyMatch` is specified.
	//
	// If the expected result count does not match the actual count,
	// by default an error is produced and the result is not committed to cached
	// context, and the user can correct the situation and try again without
	// having to manually clear out the context key using `cdk context --remove`
	//
	// If the value of * `ignoreErrorOnMissingContext` is `true`, the value of
	// `expectedMatchCount` is `at-least-one | exactly-one` and the number
	// of found resources is 0, `dummyValue` is returned and committed to context
	// instead.
	// Default: 'any'.
	//
	ExpectedMatchCount *string `field:"optional" json:"expectedMatchCount" yaml:"expectedMatchCount"`
	// Ignore an error and return the `dummyValue` instead if the resource was not found.
	//
	// - In case of an `exactIdentifier` lookup, return the `dummyValue` if the resource with
	//   that identifier was not found.
	// - In case of a `propertyMatch` lookup, return the `dummyValue` if `expectedMatchCount`
	//   is `at-least-one | exactly-one` and the number of resources found was 0.
	//
	// if `ignoreErrorOnMissingContext` is set, `dummyValue` should be set and be an array.
	// Default: false.
	//
	IgnoreErrorOnMissingContext *bool `field:"optional" json:"ignoreErrorOnMissingContext" yaml:"ignoreErrorOnMissingContext"`
	// Returns any resources matching these properties, using `ListResources`.
	//
	// By default, specifying propertyMatch will successfully return 0 or more
	// results. To throw an error if the number of results is unexpected (and
	// prevent the query results from being committed to context), specify
	// `expectedMatchCount`.
	//
	// ## Notes on property completeness
	//
	// CloudControl API's `ListResources` may return fewer properties than
	// `GetResource` would, depending on the resource implementation.
	//
	// The resources that `propertyMatch` matches against will *only ever* be the
	// properties returned by the `ListResources` call.
	// Default: - Either exactIdentifier or propertyMatch should be specified.
	//
	PropertyMatch *map[string]interface{} `field:"optional" json:"propertyMatch" yaml:"propertyMatch"`
	// The resource model to use to select the resources, using `ListResources`..
	//
	// This is needed for sub-resources where the parent Arn is required.
	//
	// See https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-list.html#resource-operations-list-containers
	// Default: - no resource Model is provided.
	//
	ResourceModel *map[string]interface{} `field:"optional" json:"resourceModel" yaml:"resourceModel"`
}

