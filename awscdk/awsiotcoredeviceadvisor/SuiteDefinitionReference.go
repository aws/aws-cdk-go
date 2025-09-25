package awsiotcoredeviceadvisor


// A reference to a SuiteDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   suiteDefinitionReference := &SuiteDefinitionReference{
//   	SuiteDefinitionArn: jsii.String("suiteDefinitionArn"),
//   	SuiteDefinitionId: jsii.String("suiteDefinitionId"),
//   }
//
type SuiteDefinitionReference struct {
	// The ARN of the SuiteDefinition resource.
	SuiteDefinitionArn *string `field:"required" json:"suiteDefinitionArn" yaml:"suiteDefinitionArn"`
	// The SuiteDefinitionId of the SuiteDefinition resource.
	SuiteDefinitionId *string `field:"required" json:"suiteDefinitionId" yaml:"suiteDefinitionId"`
}

