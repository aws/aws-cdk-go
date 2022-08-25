package awsopensearchservice


// Specifies options for fine-grained access control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedSecurityOptionsInputProperty := &advancedSecurityOptionsInputProperty{
//   	enabled: jsii.Boolean(false),
//   	internalUserDatabaseEnabled: jsii.Boolean(false),
//   	masterUserOptions: &masterUserOptionsProperty{
//   		masterUserArn: jsii.String("masterUserArn"),
//   		masterUserName: jsii.String("masterUserName"),
//   		masterUserPassword: jsii.String("masterUserPassword"),
//   	},
//   }
//
type CfnDomain_AdvancedSecurityOptionsInputProperty struct {
	// True to enable fine-grained access control.
	//
	// You must also enable encryption of data at rest and node-to-node encryption. See [Fine-grained access control in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html) .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// True to enable the internal user database.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// Specifies information about the master user.
	MasterUserOptions interface{} `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

