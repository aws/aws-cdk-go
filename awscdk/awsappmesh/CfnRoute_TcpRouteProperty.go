package awsappmesh


// An object that represents a TCP route type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpRouteProperty := &TcpRouteProperty{
//   	Action: &TcpRouteActionProperty{
//   		WeightedTargets: []interface{}{
//   			&WeightedTargetProperty{
//   				VirtualNode: jsii.String("virtualNode"),
//   				Weight: jsii.Number(123),
//
//   				// the properties below are optional
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Match: &TcpRouteMatchProperty{
//   		Port: jsii.Number(123),
//   	},
//   	Timeout: &TcpTimeoutProperty{
//   		Idle: &DurationProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnRoute_TcpRouteProperty struct {
	// The action to take if a match is determined.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An object that represents the criteria for determining a request match.
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// An object that represents types of timeouts.
	Timeout interface{} `field:"optional" json:"timeout" yaml:"timeout"`
}

