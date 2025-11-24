package mixinsawslambda


// Specific schema validation configuration settings that tell Lambda the message attributes you want to validate and filter using your schema registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaValidationConfigProperty := &SchemaValidationConfigProperty{
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemavalidationconfig.html
//
type CfnEventSourceMappingPropsMixin_SchemaValidationConfigProperty struct {
	// The attributes you want your schema registry to validate and filter for.
	//
	// If you selected `JSON` as the `EventRecordFormat` , Lambda also deserializes the selected message attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemavalidationconfig.html#cfn-lambda-eventsourcemapping-schemavalidationconfig-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

