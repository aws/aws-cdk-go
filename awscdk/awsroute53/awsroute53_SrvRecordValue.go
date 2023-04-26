package awsroute53


// Properties for a SRV record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srvRecordValue := &SrvRecordValue{
//   	HostName: jsii.String("hostName"),
//   	Port: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   	Weight: jsii.Number(123),
//   }
//
// Experimental.
type SrvRecordValue struct {
	// The server host name.
	// Experimental.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// The port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The priority.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The weight.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

