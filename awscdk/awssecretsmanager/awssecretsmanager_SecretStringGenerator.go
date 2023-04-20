package awssecretsmanager


// Configuration to generate secrets such as passwords automatically.
//
// Example:
//   // Default secret
//   secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
//   // Using the default secret
//   // Using the default secret
//   iam.NewUser(this, jsii.String("User"), &UserProps{
//   	Password: secret.secretValue,
//   })
//   // Templated secret
//   templatedSecret := secretsmanager.NewSecret(this, jsii.String("TemplatedSecret"), &SecretProps{
//   	GenerateSecretString: &SecretStringGenerator{
//   		SecretStringTemplate: jSON.stringify(map[string]*string{
//   			"username": jsii.String("user"),
//   		}),
//   		GenerateStringKey: jsii.String("password"),
//   	},
//   })
//   // Using the templated secret
//   // Using the templated secret
//   iam.NewUser(this, jsii.String("OtherUser"), &UserProps{
//   	UserName: templatedSecret.secretValueFromJson(jsii.String("username")).ToString(),
//   	Password: templatedSecret.secretValueFromJson(jsii.String("password")),
//   })
//
// Experimental.
type SecretStringGenerator struct {
	// A string that includes characters that shouldn't be included in the generated password.
	//
	// The string can be a minimum
	// of ``0`` and a maximum of ``4096`` characters long.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies that the generated password shouldn't include lowercase letters.
	// Experimental.
	ExcludeLowercase *bool `field:"optional" json:"excludeLowercase" yaml:"excludeLowercase"`
	// Specifies that the generated password shouldn't include digits.
	// Experimental.
	ExcludeNumbers *bool `field:"optional" json:"excludeNumbers" yaml:"excludeNumbers"`
	// Specifies that the generated password shouldn't include punctuation characters.
	// Experimental.
	ExcludePunctuation *bool `field:"optional" json:"excludePunctuation" yaml:"excludePunctuation"`
	// Specifies that the generated password shouldn't include uppercase letters.
	// Experimental.
	ExcludeUppercase *bool `field:"optional" json:"excludeUppercase" yaml:"excludeUppercase"`
	// The JSON key name that's used to add the generated password to the JSON structure specified by the ``secretStringTemplate`` parameter.
	//
	// If you specify ``generateStringKey`` then ``secretStringTemplate``
	// must be also be specified.
	// Experimental.
	GenerateStringKey *string `field:"optional" json:"generateStringKey" yaml:"generateStringKey"`
	// Specifies that the generated password can include the space character.
	// Experimental.
	IncludeSpace *bool `field:"optional" json:"includeSpace" yaml:"includeSpace"`
	// The desired length of the generated password.
	// Experimental.
	PasswordLength *float64 `field:"optional" json:"passwordLength" yaml:"passwordLength"`
	// Specifies whether the generated password must include at least one of every allowed character type.
	// Experimental.
	RequireEachIncludedType *bool `field:"optional" json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
	// A properly structured JSON string that the generated password can be added to.
	//
	// The ``generateStringKey`` is
	// combined with the generated random string and inserted into the JSON structure that's specified by this parameter.
	// The merged JSON string is returned as the completed SecretString of the secret. If you specify ``secretStringTemplate``
	// then ``generateStringKey`` must be also be specified.
	// Experimental.
	SecretStringTemplate *string `field:"optional" json:"secretStringTemplate" yaml:"secretStringTemplate"`
}

