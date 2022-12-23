package awslakeformation


// Properties for defining a `CfnDataLakeSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataLakeSettingsProps := &cfnDataLakeSettingsProps{
//   	admins: []interface{}{
//   		&dataLakePrincipalProperty{
//   			dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	trustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   }
//
type CfnDataLakeSettingsProps struct {
	// A list of AWS Lake Formation principals.
	Admins interface{} `field:"optional" json:"admins" yaml:"admins"`
	// An array of UTF-8 strings.
	//
	// A list of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs). The user ARNs can be logged in the resource owner's CloudTrail log. You may want to specify this property when you are in a high-trust boundary, such as the same team or company.
	TrustedResourceOwners *[]*string `field:"optional" json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

