package previewawslicensemanagermixins


// Properties for CfnLicensePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLicenseMixinProps := &CfnLicenseMixinProps{
//   	Beneficiary: jsii.String("beneficiary"),
//   	ConsumptionConfiguration: &ConsumptionConfigurationProperty{
//   		BorrowConfiguration: &BorrowConfigurationProperty{
//   			AllowEarlyCheckIn: jsii.Boolean(false),
//   			MaxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		ProvisionalConfiguration: &ProvisionalConfigurationProperty{
//   			MaxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		RenewType: jsii.String("renewType"),
//   	},
//   	Entitlements: []interface{}{
//   		&EntitlementProperty{
//   			AllowCheckIn: jsii.Boolean(false),
//   			MaxCount: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Overage: jsii.Boolean(false),
//   			Unit: jsii.String("unit"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	HomeRegion: jsii.String("homeRegion"),
//   	Issuer: &IssuerDataProperty{
//   		Name: jsii.String("name"),
//   		SignKey: jsii.String("signKey"),
//   	},
//   	LicenseMetadata: []interface{}{
//   		&MetadataProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	LicenseName: jsii.String("licenseName"),
//   	ProductName: jsii.String("productName"),
//   	ProductSku: jsii.String("productSku"),
//   	Status: jsii.String("status"),
//   	Validity: &ValidityDateFormatProperty{
//   		Begin: jsii.String("begin"),
//   		End: jsii.String("end"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html
//
type CfnLicenseMixinProps struct {
	// License beneficiary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-beneficiary
	//
	Beneficiary *string `field:"optional" json:"beneficiary" yaml:"beneficiary"`
	// Configuration for consumption of the license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-consumptionconfiguration
	//
	ConsumptionConfiguration interface{} `field:"optional" json:"consumptionConfiguration" yaml:"consumptionConfiguration"`
	// License entitlements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-entitlements
	//
	Entitlements interface{} `field:"optional" json:"entitlements" yaml:"entitlements"`
	// Home Region of the license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-homeregion
	//
	HomeRegion *string `field:"optional" json:"homeRegion" yaml:"homeRegion"`
	// License issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-issuer
	//
	Issuer interface{} `field:"optional" json:"issuer" yaml:"issuer"`
	// License metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensemetadata
	//
	LicenseMetadata interface{} `field:"optional" json:"licenseMetadata" yaml:"licenseMetadata"`
	// License name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensename
	//
	LicenseName *string `field:"optional" json:"licenseName" yaml:"licenseName"`
	// Product name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productname
	//
	ProductName *string `field:"optional" json:"productName" yaml:"productName"`
	// Product SKU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productsku
	//
	ProductSku *string `field:"optional" json:"productSku" yaml:"productSku"`
	// License status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Date and time range during which the license is valid, in ISO8601-UTC format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-validity
	//
	Validity interface{} `field:"optional" json:"validity" yaml:"validity"`
}

