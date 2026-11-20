package awscdk


// Reference to supporting context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContextRef := &CfnContextRef{
//   	At: jsii.String("at"),
//
//   	// the properties below are optional
//   	Has: jsii.String("has"),
//   	Scope: jsii.String("scope"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnContextRef struct {
	// Location of the supporting context.
	At *string `field:"required" json:"at" yaml:"at"`
	// Summary of the referenced context.
	// Default: - no summary.
	//
	Has *string `field:"optional" json:"has" yaml:"has"`
	// Scope of the referenced context.
	// Default: - no scope.
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

