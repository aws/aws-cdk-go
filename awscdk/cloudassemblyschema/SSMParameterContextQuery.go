package cloudassemblyschema


// Query to SSM Parameter Context Provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRoleAdditionalOptions interface{}
//
//   sSMParameterContextQuery := &SSMParameterContextQuery{
//   	Account: jsii.String("account"),
//   	ParameterName: jsii.String("parameterName"),
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	AssumeRoleAdditionalOptions: map[string]interface{}{
//   		"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   	},
//   	LookupRoleArn: jsii.String("lookupRoleArn"),
//   	LookupRoleExternalId: jsii.String("lookupRoleExternalId"),
//   }
//
type SSMParameterContextQuery struct {
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
	// Parameter name to query.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
}

