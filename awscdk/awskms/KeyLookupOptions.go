package awskms


// Properties for looking up an existing Key.
//
// Example:
//   myKeyLookup := kms.Key_FromLookup(this, jsii.String("MyKeyLookup"), &KeyLookupOptions{
//   	AliasName: jsii.String("alias/KeyAlias"),
//   })
//
//   role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//   myKeyLookup.GrantEncryptDecrypt(role)
//
type KeyLookupOptions struct {
	// The alias name of the Key.
	//
	// Must be in the format `alias/<AliasName>`.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
}

