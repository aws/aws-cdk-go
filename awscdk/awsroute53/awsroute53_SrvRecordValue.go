package awsroute53


// Properties for a SRV record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srvRecordValue := &srvRecordValue{
//   	hostName: jsii.String("hostName"),
//   	port: jsii.Number(123),
//   	priority: jsii.Number(123),
//   	weight: jsii.Number(123),
//   }
//
type SrvRecordValue struct {
	// The server host name.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// The port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The priority.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The weight.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

