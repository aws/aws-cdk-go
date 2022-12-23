// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Identifies a version of code that AWS App Runner refers to within a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   sourceCodeVersion := &sourceCodeVersion{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourcecodeversion.html
//
// Experimental.
type SourceCodeVersion struct {
	// The type of version identifier.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A source code version.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

