package awslightsail


// `RelationalDatabaseParameter` is a property of the [AWS::Lightsail::Database](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html) resource. It describes parameters for the database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationalDatabaseParameterProperty := &relationalDatabaseParameterProperty{
//   	allowedValues: jsii.String("allowedValues"),
//   	applyMethod: jsii.String("applyMethod"),
//   	applyType: jsii.String("applyType"),
//   	dataType: jsii.String("dataType"),
//   	description: jsii.String("description"),
//   	isModifiable: jsii.Boolean(false),
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnDatabase_RelationalDatabaseParameterProperty struct {
	// The valid range of values for the parameter.
	AllowedValues *string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Indicates when parameter updates are applied.
	//
	// Can be `immediate` or `pending-reboot` .
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
	// Specifies the engine-specific parameter type.
	ApplyType *string `field:"optional" json:"applyType" yaml:"applyType"`
	// The valid data type of the parameter.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// A description of the parameter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A Boolean value indicating whether the parameter can be modified.
	IsModifiable interface{} `field:"optional" json:"isModifiable" yaml:"isModifiable"`
	// The name of the parameter.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value for the parameter.
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

