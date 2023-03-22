package awsdynamodb


// Represents the settings used to enable or disable Time to Live (TTL) for the specified table.
//
// All replicas will have the same time to live configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeToLiveSpecificationProperty := &timeToLiveSpecificationProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	attributeName: jsii.String("attributeName"),
//   }
//
type CfnGlobalTable_TimeToLiveSpecificationProperty struct {
	// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the attribute used to store the expiration time for items in the table.
	//
	// Currently, you cannot directly change the attribute name used to evaluate time to live. In order to do so, you must first disable time to live, and then re-enable it with the new attribute name. It can take up to one hour for changes to time to live to take effect. If you attempt to modify time to live within that time window, your stack operation might be delayed.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
}

