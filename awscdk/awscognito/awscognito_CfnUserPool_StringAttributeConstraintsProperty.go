package awscognito


// The `StringAttributeConstraints` property type defines the string attribute constraints of an Amazon Cognito user pool.
//
// `StringAttributeConstraints` is a subproperty of the [SchemaAttribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringAttributeConstraintsProperty := &stringAttributeConstraintsProperty{
//   	maxLength: jsii.String("maxLength"),
//   	minLength: jsii.String("minLength"),
//   }
//
type CfnUserPool_StringAttributeConstraintsProperty struct {
	// The maximum length.
	MaxLength *string `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The minimum length.
	MinLength *string `field:"optional" json:"minLength" yaml:"minLength"`
}

