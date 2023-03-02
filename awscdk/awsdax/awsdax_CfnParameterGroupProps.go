package awsdax


// Properties for defining a `CfnParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterNameValues interface{}
//
//   cfnParameterGroupProps := &cfnParameterGroupProps{
//   	description: jsii.String("description"),
//   	parameterGroupName: jsii.String("parameterGroupName"),
//   	parameterNameValues: parameterNameValues,
//   }
//
type CfnParameterGroupProps struct {
	// A description of the parameter group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter group.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// An array of name-value pairs for the parameters in the group.
	//
	// Each element in the array represents a single parameter.
	//
	// > `record-ttl-millis` and `query-ttl-millis` are the only supported parameter names. For more details, see [Configuring TTL Settings](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.cluster-management.html#DAX.cluster-management.custom-settings.ttl) .
	ParameterNameValues interface{} `field:"optional" json:"parameterNameValues" yaml:"parameterNameValues"`
}

