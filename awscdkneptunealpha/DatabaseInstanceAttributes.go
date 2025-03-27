package awscdkneptunealpha


// Properties that describe an existing instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import neptune_alpha "github.com/aws/aws-cdk-go/awscdkneptunealpha"
//
//   databaseInstanceAttributes := &DatabaseInstanceAttributes{
//   	InstanceEndpointAddress: jsii.String("instanceEndpointAddress"),
//   	InstanceIdentifier: jsii.String("instanceIdentifier"),
//   	Port: jsii.Number(123),
//   }
//
// Experimental.
type DatabaseInstanceAttributes struct {
	// The endpoint address.
	// Experimental.
	InstanceEndpointAddress *string `field:"required" json:"instanceEndpointAddress" yaml:"instanceEndpointAddress"`
	// The instance identifier.
	// Experimental.
	InstanceIdentifier *string `field:"required" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The database port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

