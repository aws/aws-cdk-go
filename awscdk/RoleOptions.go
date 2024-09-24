package awscdk


// Options for specifying a role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRoleAdditionalOptions interface{}
//
//   roleOptions := &RoleOptions{
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//
//   	// the properties below are optional
//   	AssumeRoleAdditionalOptions: map[string]interface{}{
//   		"assumeRoleAdditionalOptionsKey": assumeRoleAdditionalOptions,
//   	},
//   	AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   }
//
type RoleOptions struct {
	// ARN of the role to assume.
	AssumeRoleArn *string `field:"required" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Additional options to pass to STS when assuming the role for cloudformation deployments.
	//
	// - `RoleArn` should not be used. Use the dedicated `assumeRoleArn` property instead.
	// - `ExternalId` should not be used. Use the dedicated `assumeRoleExternalId` instead.
	// - `TransitiveTagKeys` defaults to use all keys (if any) specified in `Tags`. E.g, all tags are transitive by default.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/STS.html#assumeRole-property
	//
	// Default: - No additional options.
	//
	AssumeRoleAdditionalOptions *map[string]interface{} `field:"optional" json:"assumeRoleAdditionalOptions" yaml:"assumeRoleAdditionalOptions"`
	// External ID to use when assuming the role.
	// Default: - No external ID.
	//
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
}

