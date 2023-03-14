package awslightsail


// `State` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the status code and the state (for example, `running` ) of an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateProperty := &stateProperty{
//   	code: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnInstance_StateProperty struct {
	// The status code of the instance.
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// The state of the instance (for example, `running` or `pending` ).
	Name *string `field:"optional" json:"name" yaml:"name"`
}

