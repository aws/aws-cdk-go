package interfacesawsgreengrass


// A reference to a LoggerDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerDefinitionReference := &LoggerDefinitionReference{
//   	LoggerDefinitionArn: jsii.String("loggerDefinitionArn"),
//   	LoggerDefinitionId: jsii.String("loggerDefinitionId"),
//   }
//
type LoggerDefinitionReference struct {
	// The ARN of the LoggerDefinition resource.
	LoggerDefinitionArn *string `field:"required" json:"loggerDefinitionArn" yaml:"loggerDefinitionArn"`
	// The Id of the LoggerDefinition resource.
	LoggerDefinitionId *string `field:"required" json:"loggerDefinitionId" yaml:"loggerDefinitionId"`
}

