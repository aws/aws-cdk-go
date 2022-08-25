package awsapigateway


// Properties for defining a `CfnModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnModelProps := &cfnModelProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	schema: schema,
//   }
//
type CfnModelProps struct {
	// The ID of a REST API with which to associate this model.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The content type for the model.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A description that identifies this model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema to use to transform data to one or more output formats.
	//
	// Specify null ( `{}` ) if you don't want to specify a schema.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

