package awsroute53


// Base properties of an SVCB ServiceMode record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alpn Alpn
//
//   svcbRecordServiceModeProps := &SvcbRecordServiceModeProps{
//   	Alpn: []Alpn{
//   		alpn,
//   	},
//   	Ipv4hint: []*string{
//   		jsii.String("ipv4hint"),
//   	},
//   	Ipv6hint: []*string{
//   		jsii.String("ipv6hint"),
//   	},
//   	Mandatory: []*string{
//   		jsii.String("mandatory"),
//   	},
//   	NoDefaultAlpn: jsii.Boolean(false),
//   	Port: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   	TargetName: jsii.String("targetName"),
//   }
//
type SvcbRecordServiceModeProps struct {
	// Indicates the set of Application-Layer Protocol Negotiation (ALPN) protocol identifiers and associated transport protocols supported by this service endpoint.
	// Default: - No ALPN protocol identifiers.
	//
	Alpn *[]Alpn `field:"optional" json:"alpn" yaml:"alpn"`
	// Conveys that clients may use to reach the service.
	// Default: - No hints.
	//
	Ipv4hint *[]*string `field:"optional" json:"ipv4hint" yaml:"ipv4hint"`
	// Conveys that clients may use to reach the service.
	// Default: - No hints.
	//
	Ipv6hint *[]*string `field:"optional" json:"ipv6hint" yaml:"ipv6hint"`
	// Indicates mandatory keys.
	// Default: - No mandatory keys.
	//
	Mandatory *[]*string `field:"optional" json:"mandatory" yaml:"mandatory"`
	// Indicates no default ALPN protocol identifiers.
	//
	// The `alpn` parameter must be supplied together.
	// Default: false.
	//
	NoDefaultAlpn *bool `field:"optional" json:"noDefaultAlpn" yaml:"noDefaultAlpn"`
	// The alternative port number.
	// Default: - Use the default port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The priority.
	// Default: 1.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The domain name of the alternative endpoint.
	// Default: '.' - The record name of the record itself
	//
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
}

