package awslicensemanager


// Describes a resource entitled for use with a license.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entitlementProperty := &entitlementProperty{
//   	name: jsii.String("name"),
//   	unit: jsii.String("unit"),
//
//   	// the properties below are optional
//   	allowCheckIn: jsii.Boolean(false),
//   	maxCount: jsii.Number(123),
//   	overage: jsii.Boolean(false),
//   	value: jsii.String("value"),
//   }
//
type CfnLicense_EntitlementProperty struct {
	// Entitlement name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Entitlement unit.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Indicates whether check-ins are allowed.
	AllowCheckIn interface{} `field:"optional" json:"allowCheckIn" yaml:"allowCheckIn"`
	// Maximum entitlement count.
	//
	// Use if the unit is not None.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Indicates whether overages are allowed.
	Overage interface{} `field:"optional" json:"overage" yaml:"overage"`
	// Entitlement resource.
	//
	// Use only if the unit is None.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

