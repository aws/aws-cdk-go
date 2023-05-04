package awsdevicefarm


// Represents a condition for a device pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	Attribute: jsii.String("attribute"),
//   	Operator: jsii.String("operator"),
//   	Value: jsii.String("value"),
//   }
//
type CfnDevicePool_RuleProperty struct {
	// The rule's stringified attribute. For example, specify the value as `"\"abc\""` .
	//
	// The supported operators for each attribute are provided in the following list.
	//
	// - **APPIUM_VERSION** - The Appium version for the test.
	//
	// Supported operators: `CONTAINS`
	// - **ARN** - The Amazon Resource Name (ARN) of the device (for example, `arn:aws:devicefarm:us-west-2::device:12345Example` .
	//
	// Supported operators: `EQUALS` , `IN` , `NOT_IN`
	// - **AVAILABILITY** - The current availability of the device. Valid values are AVAILABLE, HIGHLY_AVAILABLE, BUSY, or TEMPORARY_NOT_AVAILABLE.
	//
	// Supported operators: `EQUALS`
	// - **FLEET_TYPE** - The fleet type. Valid values are PUBLIC or PRIVATE.
	//
	// Supported operators: `EQUALS`
	// - **FORM_FACTOR** - The device form factor. Valid values are PHONE or TABLET.
	//
	// Supported operators: `EQUALS` , `IN` , `NOT_IN`
	// - **INSTANCE_ARN** - The Amazon Resource Name (ARN) of the device instance.
	//
	// Supported operators: `IN` , `NOT_IN`
	// - **INSTANCE_LABELS** - The label of the device instance.
	//
	// Supported operators: `CONTAINS`
	// - **MANUFACTURER** - The device manufacturer (for example, Apple).
	//
	// Supported operators: `EQUALS` , `IN` , `NOT_IN`
	// - **MODEL** - The device model, such as Apple iPad Air 2 or Google Pixel.
	//
	// Supported operators: `CONTAINS` , `EQUALS` , `IN` , `NOT_IN`
	// - **OS_VERSION** - The operating system version (for example, 10.3.2).
	//
	// Supported operators: `EQUALS` , `GREATER_THAN` , `GREATER_THAN_OR_EQUALS` , `IN` , `LESS_THAN` , `LESS_THAN_OR_EQUALS` , `NOT_IN`
	// - **PLATFORM** - The device platform. Valid values are ANDROID or IOS.
	//
	// Supported operators: `EQUALS` , `IN` , `NOT_IN`
	// - **REMOTE_ACCESS_ENABLED** - Whether the device is enabled for remote access. Valid values are TRUE or FALSE.
	//
	// Supported operators: `EQUALS`
	// - **REMOTE_DEBUG_ENABLED** - Whether the device is enabled for remote debugging. Valid values are TRUE or FALSE.
	//
	// Supported operators: `EQUALS`
	//
	// Because remote debugging is [no longer supported](https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html) , this filter is ignored.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Specifies how Device Farm compares the rule's attribute to the value.
	//
	// For the operators that are supported by each attribute, see the attribute descriptions.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The rule's value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

