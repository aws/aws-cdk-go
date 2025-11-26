package previewawsguarddutyevents


// Type definition for LocalPortDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localPortDetails := &LocalPortDetails{
//   	Port: []*string{
//   		jsii.String("port"),
//   	},
//   	PortName: []*string{
//   		jsii.String("portName"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_LocalPortDetails struct {
	// port property.
	//
	// Specify an array of string values to match this event if the actual value of port is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Port *[]*string `field:"optional" json:"port" yaml:"port"`
	// portName property.
	//
	// Specify an array of string values to match this event if the actual value of portName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PortName *[]*string `field:"optional" json:"portName" yaml:"portName"`
}

