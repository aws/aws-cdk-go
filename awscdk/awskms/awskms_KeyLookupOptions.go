package awskms


// Properties for looking up an existing Key.
//
// Example:
//   myKeyLookup := kms.key.fromLookup(this, jsii.String("MyKeyLookup"), &keyLookupOptions{
//   	aliasName: jsii.String("alias/KeyAlias"),
//   })
//
//   role := iam.NewRole(this, jsii.String("MyRole"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//   myKeyLookup.grantEncryptDecrypt(role)
//
type KeyLookupOptions struct {
	// The alias name of the Key.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
}

