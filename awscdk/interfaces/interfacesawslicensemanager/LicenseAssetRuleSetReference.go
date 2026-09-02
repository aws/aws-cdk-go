package interfacesawslicensemanager


// A reference to a LicenseAssetRuleSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseAssetRuleSetReference := &LicenseAssetRuleSetReference{
//   	LicenseAssetRulesetArn: jsii.String("licenseAssetRulesetArn"),
//   }
//
type LicenseAssetRuleSetReference struct {
	// The LicenseAssetRulesetArn of the LicenseAssetRuleSet resource.
	LicenseAssetRulesetArn *string `field:"required" json:"licenseAssetRulesetArn" yaml:"licenseAssetRulesetArn"`
}

