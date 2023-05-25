package awsgreengrass


// Properties for defining a `CfnLoggerDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggerDefinitionVersionProps := &CfnLoggerDefinitionVersionProps{
//   	LoggerDefinitionId: jsii.String("loggerDefinitionId"),
//   	Loggers: []interface{}{
//   		&LoggerProperty{
//   			Component: jsii.String("component"),
//   			Id: jsii.String("id"),
//   			Level: jsii.String("level"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Space: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnLoggerDefinitionVersionProps struct {
	// The ID of the logger definition associated with this version.
	//
	// This value is a GUID.
	LoggerDefinitionId *string `field:"required" json:"loggerDefinitionId" yaml:"loggerDefinitionId"`
	// The loggers in this version.
	Loggers interface{} `field:"required" json:"loggers" yaml:"loggers"`
}

