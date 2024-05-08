package awsqbusiness


// Configuration information required to create a custom plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPluginConfigurationProperty := &CustomPluginConfigurationProperty{
//   	ApiSchema: &APISchemaProperty{
//   		Payload: jsii.String("payload"),
//   		S3: &S3Property{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	ApiSchemaType: jsii.String("apiSchemaType"),
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html
//
type CfnPlugin_CustomPluginConfigurationProperty struct {
	// Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html#cfn-qbusiness-plugin-custompluginconfiguration-apischema
	//
	ApiSchema interface{} `field:"required" json:"apiSchema" yaml:"apiSchema"`
	// The type of OpenAPI schema to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html#cfn-qbusiness-plugin-custompluginconfiguration-apischematype
	//
	ApiSchemaType *string `field:"required" json:"apiSchemaType" yaml:"apiSchemaType"`
	// A description for your custom plugin configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-custompluginconfiguration.html#cfn-qbusiness-plugin-custompluginconfiguration-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
}

