package awsappmesh


// An object that represents a TCP route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpRouteProperty := &tcpRouteProperty{
//   	action: &tcpRouteActionProperty{
//   		weightedTargets: []interface{}{
//   			&weightedTargetProperty{
//   				virtualNode: jsii.String("virtualNode"),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	timeout: &tcpTimeoutProperty{
//   		idle: &durationProperty{
//   			unit: jsii.String("unit"),
//   			value: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnRoute_TcpRouteProperty struct {
	// The action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents types of timeouts.
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

