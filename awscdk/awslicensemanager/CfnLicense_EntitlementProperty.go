package awslicensemanager


// Describes a resource entitled for use with a license.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entitlementProperty := &EntitlementProperty{
//   	Name: jsii.String("name"),
//   	Unit: jsii.String("unit"),
//
//   	// the properties below are optional
//   	AllowCheckIn: jsii.Boolean(false),
//   	MaxCount: jsii.Number(123),
//   	Overage: jsii.Boolean(false),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html
//
type CfnLicense_EntitlementProperty struct {
	// Entitlement name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Entitlement unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Indicates whether check-ins are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-allowcheckin
	//
	AllowCheckIn interface{} `field:"optional" json:"allowCheckIn" yaml:"allowCheckIn"`
	// Maximum entitlement count.
	//
	// Use if the unit is not None.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-maxcount
	//
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Indicates whether overages are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-overage
	//
	Overage interface{} `field:"optional" json:"overage" yaml:"overage"`
	// Entitlement resource.
	//
	// Use only if the unit is None.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-license-entitlement.html#cfn-licensemanager-license-entitlement-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

