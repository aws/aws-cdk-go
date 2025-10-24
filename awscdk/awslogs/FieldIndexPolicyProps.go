package awslogs


// Properties for creating field index policies.
//
// Example:
//   fieldIndexPolicy := logs.NewFieldIndexPolicy(&FieldIndexPolicyProps{
//   	Fields: []*string{
//   		jsii.String("Operation"),
//   		jsii.String("RequestId"),
//   	},
//   })
//
//   logs.NewLogGroup(this, jsii.String("LogGroup"), &LogGroupProps{
//   	LogGroupName: jsii.String("cdkIntegLogGroup"),
//   	FieldIndexPolicies: []FieldIndexPolicy{
//   		fieldIndexPolicy,
//   	},
//   })
//
type FieldIndexPolicyProps struct {
	// List of fields to index in log events.
	// Default: no fields.
	//
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
}

