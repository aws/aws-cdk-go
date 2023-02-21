package awsappmesh


// An object that represents types of timeouts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpTimeoutProperty := &httpTimeoutProperty{
//   	idle: &durationProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//   	perRequest: &durationProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//   }
//
type CfnRoute_HttpTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `field:"optional" json:"idle" yaml:"idle"`
	// An object that represents a per request timeout.
	//
	// The default value is 15 seconds. If you set a higher timeout, then make sure that the higher value is set for each App Mesh resource in a conversation. For example, if a virtual node backend uses a virtual router provider to route to another virtual node, then the timeout should be greater than 15 seconds for the source and destination virtual node and the route.
	PerRequest interface{} `field:"optional" json:"perRequest" yaml:"perRequest"`
}

