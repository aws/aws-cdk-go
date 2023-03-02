package awsappmesh


// An object that represents types of timeouts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tcpTimeoutProperty := &tcpTimeoutProperty{
//   	idle: &durationProperty{
//   		unit: jsii.String("unit"),
//   		value: jsii.Number(123),
//   	},
//   }
//
type CfnRoute_TcpTimeoutProperty struct {
	// An object that represents an idle timeout.
	//
	// An idle timeout bounds the amount of time that a connection may be idle. The default value is none.
	Idle interface{} `field:"optional" json:"idle" yaml:"idle"`
}

