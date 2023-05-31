package awsredshiftserverless


// A array of parameters to set for more control over a serverless database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configParameterProperty := &ConfigParameterProperty{
//   	ParameterKey: jsii.String("parameterKey"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
type CfnWorkgroup_ConfigParameterProperty struct {
	// The key of the parameter.
	//
	// The options are `datestyle` , `enable_user_activity_logging` , `query_group` , `search_path` , and `max_query_execution_time` .
	ParameterKey *string `field:"optional" json:"parameterKey" yaml:"parameterKey"`
	// The value of the parameter to set.
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

