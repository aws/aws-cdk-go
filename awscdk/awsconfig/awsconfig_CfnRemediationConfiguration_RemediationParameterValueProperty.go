package awsconfig


// The value is either a dynamic (resource) value or a static value.
//
// You must select either a dynamic value or a static value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remediationParameterValueProperty := &remediationParameterValueProperty{
//   	resourceValue: &resourceValueProperty{
//   		value: jsii.String("value"),
//   	},
//   	staticValue: &staticValueProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
type CfnRemediationConfiguration_RemediationParameterValueProperty struct {
	// The value is dynamic and changes at run-time.
	ResourceValue interface{} `field:"optional" json:"resourceValue" yaml:"resourceValue"`
	// The value is static and does not change at run-time.
	StaticValue interface{} `field:"optional" json:"staticValue" yaml:"staticValue"`
}

