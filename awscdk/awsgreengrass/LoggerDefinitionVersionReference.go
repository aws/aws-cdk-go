package awsgreengrass


// A reference to a LoggerDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerDefinitionVersionReference := &LoggerDefinitionVersionReference{
//   	LoggerDefinitionVersionId: jsii.String("loggerDefinitionVersionId"),
//   }
//
type LoggerDefinitionVersionReference struct {
	// The Id of the LoggerDefinitionVersion resource.
	LoggerDefinitionVersionId *string `field:"required" json:"loggerDefinitionVersionId" yaml:"loggerDefinitionVersionId"`
}

