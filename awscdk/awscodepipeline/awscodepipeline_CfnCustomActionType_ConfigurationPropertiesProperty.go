package awscodepipeline


// The configuration properties for the custom action.
//
// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationPropertiesProperty := &configurationPropertiesProperty{
//   	key: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	required: jsii.Boolean(false),
//   	secret: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	queryable: jsii.Boolean(false),
//   	type: jsii.String("type"),
//   }
//
type CfnCustomActionType_ConfigurationPropertiesProperty struct {
	// Whether the configuration property is a key.
	Key interface{} `field:"required" json:"key" yaml:"key"`
	// The name of the action configuration property.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the configuration property is a required value.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Whether the configuration property is secret.
	//
	// Secrets are hidden from all calls except for `GetJobDetails` , `GetThirdPartyJobDetails` , `PollForJobs` , and `PollForThirdPartyJobs` .
	//
	// When updating a pipeline, passing * * * * * without changing any other values of the action preserves the previous value of the secret.
	Secret interface{} `field:"required" json:"secret" yaml:"secret"`
	// The description of the action configuration property that is displayed to users.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates that the property is used with `PollForJobs` .
	//
	// When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.
	//
	// If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens.
	Queryable interface{} `field:"optional" json:"queryable" yaml:"queryable"`
	// The type of the configuration property.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

