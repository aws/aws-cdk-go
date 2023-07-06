package awsverifiedpermissions


// Properties for defining a `CfnPolicyStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyStoreProps := &CfnPolicyStoreProps{
//   	ValidationSettings: &ValidationSettingsProperty{
//   		Mode: jsii.String("mode"),
//   	},
//
//   	// the properties below are optional
//   	Schema: &SchemaDefinitionProperty{
//   		CedarJson: jsii.String("cedarJson"),
//   	},
//   }
//
type CfnPolicyStoreProps struct {
	// Specifies the validation setting for this policy store.
	//
	// Currently, the only valid and required value is `Mode` .
	//
	// > We recommend that you turn on `STRICT` mode only after you define a schema. If a schema doesn't exist, then `STRICT` mode causes any policy to fail validation, and Verified Permissions rejects the policy. You can turn off validation by using the [UpdatePolicyStore](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyStore) . Then, when you have a schema defined, use [UpdatePolicyStore](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyStore) again to turn validation back on.
	ValidationSettings interface{} `field:"required" json:"validationSettings" yaml:"validationSettings"`
	// Creates or updates the policy schema in a policy store.
	//
	// Cedar can use the schema to validate any Cedar policies and policy templates submitted to the policy store. Any changes to the schema validate only policies and templates submitted after the schema change. Existing policies and templates are not re-evaluated against the changed schema. If you later update a policy, then it is evaluated against the new schema at that time.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

