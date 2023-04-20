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
type CfnLicenseProps struct {
	// Configuration for consumption of the license.
	ConsumptionConfiguration interface{} `field:"required" json:"consumptionConfiguration" yaml:"consumptionConfiguration"`
	// License entitlements.
	Entitlements interface{} `field:"required" json:"entitlements" yaml:"entitlements"`
	// Home Region of the license.
	HomeRegion *string `field:"required" json:"homeRegion" yaml:"homeRegion"`
	// License issuer.
	Issuer interface{} `field:"required" json:"issuer" yaml:"issuer"`
	// License name.
	LicenseName *string `field:"required" json:"licenseName" yaml:"licenseName"`
	// Product name.
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// Date and time range during which the license is valid, in ISO8601-UTC format.
	Validity interface{} `field:"required" json:"validity" yaml:"validity"`
	// License beneficiary.
	Beneficiary *string `field:"optional" json:"beneficiary" yaml:"beneficiary"`
	// License metadata.
	LicenseMetadata interface{} `field:"optional" json:"licenseMetadata" yaml:"licenseMetadata"`
	// Product SKU.
	ProductSku *string `field:"optional" json:"productSku" yaml:"productSku"`
	// License status.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

