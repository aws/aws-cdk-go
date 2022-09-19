package cloudassemblyschema


// Query input for looking up a KMS Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyContextQuery := &keyContextQuery{
//   	account: jsii.String("account"),
//   	aliasName: jsii.String("aliasName"),
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
type KeyContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Alias name used to search the Key.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

