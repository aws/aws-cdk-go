package awsdynamodb


// Represents the settings used to enable or disable Time to Live (TTL) for the specified table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeToLiveSpecificationProperty := &timeToLiveSpecificationProperty{
//   	attributeName: jsii.String("attributeName"),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnTable_TimeToLiveSpecificationProperty struct {
	// The name of the TTL attribute used to store the expiration time for items in the table.
	//
	// > To update this property, you must first disable TTL then enable TTL with the new attribute name.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

