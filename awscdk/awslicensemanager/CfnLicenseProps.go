package awslicensemanager


// Properties for defining a `CfnLicense`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLicenseProps := &CfnLicenseProps{
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
//   			Name: jsii.String("name"),
//   			Unit: jsii.String("unit"),
//
//   			// the properties below are optional
//   			AllowCheckIn: jsii.Boolean(false),
//   			MaxCount: jsii.Number(123),
//   			Overage: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	HomeRegion: jsii.String("homeRegion"),
//   	Issuer: &IssuerDataProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		SignKey: jsii.String("signKey"),
//   	},
//   	LicenseName: jsii.String("licenseName"),
//   	ProductName: jsii.String("productName"),
//   	Validity: &ValidityDateFormatProperty{
//   		Begin: jsii.String("begin"),
//   		End: jsii.String("end"),
//   	},
//
//   	// the properties below are optional
//   	Beneficiary: jsii.String("beneficiary"),
//   	LicenseMetadata: []interface{}{
//   		&MetadataProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ProductSku: jsii.String("productSku"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html
//
type CfnLicenseProps struct {
	// Configuration for consumption of the license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-consumptionconfiguration
	//
	ConsumptionConfiguration interface{} `field:"required" json:"consumptionConfiguration" yaml:"consumptionConfiguration"`
	// License entitlements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-entitlements
	//
	Entitlements interface{} `field:"required" json:"entitlements" yaml:"entitlements"`
	// Home Region of the license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-homeregion
	//
	HomeRegion *string `field:"required" json:"homeRegion" yaml:"homeRegion"`
	// License issuer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-issuer
	//
	Issuer interface{} `field:"required" json:"issuer" yaml:"issuer"`
	// License name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensename
	//
	LicenseName *string `field:"required" json:"licenseName" yaml:"licenseName"`
	// Product name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productname
	//
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// Date and time range during which the license is valid, in ISO8601-UTC format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-validity
	//
	Validity interface{} `field:"required" json:"validity" yaml:"validity"`
	// License beneficiary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-beneficiary
	//
	Beneficiary *string `field:"optional" json:"beneficiary" yaml:"beneficiary"`
	// License metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensemetadata
	//
	LicenseMetadata interface{} `field:"optional" json:"licenseMetadata" yaml:"licenseMetadata"`
	// Product SKU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productsku
	//
	ProductSku *string `field:"optional" json:"productSku" yaml:"productSku"`
	// License status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

