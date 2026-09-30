package awsfms


// An individual Firewall Manager application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appProperty := &AppProperty{
//   	AppName: jsii.String("appName"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-applicationslist-app.html
//
type CfnApplicationsList_AppProperty struct {
	// The application's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-applicationslist-app.html#cfn-fms-applicationslist-app-appname
	//
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The application's port number, for example 80.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-applicationslist-app.html#cfn-fms-applicationslist-app-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The IP protocol name or number.
	//
	// The name can be one of tcp, udp, or icmp. For information on possible numbers, see Protocol Numbers (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-applicationslist-app.html#cfn-fms-applicationslist-app-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

