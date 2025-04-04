package cloudassemblyschema


// Query input for lookup up Cloudformation resources using CC API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRoleAdditionalOptions interface{}
//   var propertyMatch interface{}
//
//   ccApiContextQuery := &CcApiContextQuery{
//   	Account: jsii.String("account"),
//   	PropertiesToReturn: []*string{
//   		jsii.String("propertiesToReturn"),
//   	},
//   	Region: jsii.String("region"),
//   	TypeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	AssumeRoleAdditionalOptions: map[string]interface{}{
//   		"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   	},
//   	ExactIdentifier: jsii.String("exactIdentifier"),
//   	LookupRoleArn: jsii.String("lookupRoleArn"),
//   	LookupRoleExternalId: jsii.String("lookupRoleExternalId"),
//   	PropertyMatch: map[string]interface{}{
//   		"propertyMatchKey": propertyMatch,
//   	},
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
	PropertiesToReturn *[]*string `field:"required" json:"propertiesToReturn" yaml:"propertiesToReturn"`
	// The Cloudformation resource type.
	//
	// See https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/supported-resources.html
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// exactIdentifier of the resource.
	//
	// Specifying exactIdentifier will return at most one result.
	// Either exactIdentifier or propertyMatch should be specified.
	// Default: - None.
	//
	ExactIdentifier *string `field:"optional" json:"exactIdentifier" yaml:"exactIdentifier"`
	// This indicates the property to search for.
	//
	// If both exactIdentifier and propertyMatch are specified, then exactIdentifier is used.
	// Specifying propertyMatch will return 0 or more results.
	// Either exactIdentifier or propertyMatch should be specified.
	// Default: - None.
	//
	PropertyMatch *map[string]interface{} `field:"optional" json:"propertyMatch" yaml:"propertyMatch"`
}

