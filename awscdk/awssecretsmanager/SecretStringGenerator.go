package awssecretsmanager


// Configuration to generate secrets such as passwords automatically.
//
// Example:
//   var vpc iVpc
//
//
//   instance1 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance1"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_POSTGRES(),
//   	// Generate the secret with admin username `postgres` and random password
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("postgres")),
//   	Vpc: Vpc,
//   })
//   // Templated secret with username and password fields
//   templatedSecret := secretsmanager.NewSecret(this, jsii.String("TemplatedSecret"), &SecretProps{
//   	GenerateSecretString: &SecretStringGenerator{
//   		SecretStringTemplate: jSON.stringify(map[string]*string{
//   			"username": jsii.String("postgres"),
//   		}),
//   		GenerateStringKey: jsii.String("password"),
//   		ExcludeCharacters: jsii.String("/@\""),
//   	},
//   })
//   // Using the templated secret as credentials
//   instance2 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance2"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_POSTGRES(),
//   	Credentials: map[string]interface{}{
//   		"username": templatedSecret.secretValueFromJson(jsii.String("username")).toString(),
//   		"password": templatedSecret.secretValueFromJson(jsii.String("password")),
//   	},
//   	Vpc: Vpc,
//   })
//
type SecretStringGenerator struct {
	// A string that includes characters that shouldn't be included in the generated password.
	//
	// The string can be a minimum
	// of ``0`` and a maximum of ``4096`` characters long.
	// Default: no exclusions.
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies that the generated password shouldn't include lowercase letters.
	// Default: false.
	//
	ExcludeLowercase *bool `field:"optional" json:"excludeLowercase" yaml:"excludeLowercase"`
	// Specifies that the generated password shouldn't include digits.
	// Default: false.
	//
	ExcludeNumbers *bool `field:"optional" json:"excludeNumbers" yaml:"excludeNumbers"`
	// Specifies that the generated password shouldn't include punctuation characters.
	// Default: false.
	//
	ExcludePunctuation *bool `field:"optional" json:"excludePunctuation" yaml:"excludePunctuation"`
	// Specifies that the generated password shouldn't include uppercase letters.
	// Default: false.
	//
	ExcludeUppercase *bool `field:"optional" json:"excludeUppercase" yaml:"excludeUppercase"`
	// The JSON key name that's used to add the generated password to the JSON structure specified by the ``secretStringTemplate`` parameter.
	//
	// If you specify ``generateStringKey`` then ``secretStringTemplate``
	// must be also be specified.
	GenerateStringKey *string `field:"optional" json:"generateStringKey" yaml:"generateStringKey"`
	// Specifies that the generated password can include the space character.
	// Default: false.
	//
	IncludeSpace *bool `field:"optional" json:"includeSpace" yaml:"includeSpace"`
	// The desired length of the generated password.
	// Default: 32.
	//
	PasswordLength *float64 `field:"optional" json:"passwordLength" yaml:"passwordLength"`
	// Specifies whether the generated password must include at least one of every allowed character type.
	// Default: true.
	//
	RequireEachIncludedType *bool `field:"optional" json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
	// A properly structured JSON string that the generated password can be added to.
	//
	// The ``generateStringKey`` is
	// combined with the generated random string and inserted into the JSON structure that's specified by this parameter.
	// The merged JSON string is returned as the completed SecretString of the secret. If you specify ``secretStringTemplate``
	// then ``generateStringKey`` must be also be specified.
	SecretStringTemplate *string `field:"optional" json:"secretStringTemplate" yaml:"secretStringTemplate"`
}

