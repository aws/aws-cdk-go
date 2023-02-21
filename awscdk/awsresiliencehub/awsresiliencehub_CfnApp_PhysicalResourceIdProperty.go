package awsresiliencehub


// Defines a physical resource identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   physicalResourceIdProperty := &physicalResourceIdProperty{
//   	identifier: jsii.String("identifier"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	awsAccountId: jsii.String("awsAccountId"),
//   	awsRegion: jsii.String("awsRegion"),
//   }
//
type CfnApp_PhysicalResourceIdProperty struct {
	// The identifier of the physical resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Specifies the type of physical resource identifier.
	//
	// - **Arn** - The resource identifier is an Amazon Resource Name (ARN) .
	// - **Native** - The resource identifier is a Resilience Hub-native identifier.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The AWS account that owns the physical resource.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The AWS Region that the physical resource is located in.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
}

