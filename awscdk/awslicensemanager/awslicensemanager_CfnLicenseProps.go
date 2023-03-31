package awslicensemanager


// Properties for defining a `CfnLicense`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLicenseProps := &cfnLicenseProps{
//   	consumptionConfiguration: &consumptionConfigurationProperty{
//   		borrowConfiguration: &borrowConfigurationProperty{
//   			allowEarlyCheckIn: jsii.Boolean(false),
//   			maxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		provisionalConfiguration: &provisionalConfigurationProperty{
//   			maxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		renewType: jsii.String("renewType"),
//   	},
//   	entitlements: []interface{}{
//   		&entitlementProperty{
//   			name: jsii.String("name"),
//   			unit: jsii.String("unit"),
//
//   			// the properties below are optional
//   			allowCheckIn: jsii.Boolean(false),
//   			maxCount: jsii.Number(123),
//   			overage: jsii.Boolean(false),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	homeRegion: jsii.String("homeRegion"),
//   	issuer: &issuerDataProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		signKey: jsii.String("signKey"),
//   	},
//   	licenseName: jsii.String("licenseName"),
//   	productName: jsii.String("productName"),
//   	validity: &validityDateFormatProperty{
//   		begin: jsii.String("begin"),
//   		end: jsii.String("end"),
//   	},
//
//   	// the properties below are optional
//   	beneficiary: jsii.String("beneficiary"),
//   	licenseMetadata: []interface{}{
//   		&metadataProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	productSku: jsii.String("productSku"),
//   	status: jsii.String("status"),
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

