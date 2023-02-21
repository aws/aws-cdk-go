package awsiotsitewise


// Contains an asset attribute property.
//
// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#attributes) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &attributeProperty{
//   	defaultValue: jsii.String("defaultValue"),
//   }
//
type CfnAssetModel_AttributeProperty struct {
	// The default value of the asset model property attribute.
	//
	// All assets that you create from the asset model contain this attribute value. You can update an attribute's value after you create an asset. For more information, see [Updating attribute values](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-attribute-values.html) in the *AWS IoT SiteWise User Guide* .
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}

