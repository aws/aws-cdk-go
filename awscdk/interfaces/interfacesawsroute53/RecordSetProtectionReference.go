package interfacesawsroute53


// A reference to a RecordSetProtection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordSetProtectionReference := &RecordSetProtectionReference{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	RecordName: jsii.String("recordName"),
//   	RecordType: jsii.String("recordType"),
//   }
//
type RecordSetProtectionReference struct {
	// The HostedZoneId of the RecordSetProtection resource.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The RecordName of the RecordSetProtection resource.
	RecordName *string `field:"required" json:"recordName" yaml:"recordName"`
	// The RecordType of the RecordSetProtection resource.
	RecordType *string `field:"required" json:"recordType" yaml:"recordType"`
}

