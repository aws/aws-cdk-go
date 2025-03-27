package awssecretsmanager


// Generates a random password.
//
// We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
//
// *Required permissions:* `secretsmanager:GetRandomPassword` . For more information, see [IAM policy actions for Secrets Manager](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awssecretsmanager.html#awssecretsmanager-actions-as-permissions) and [Authentication and access control in Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generateSecretStringProperty := &GenerateSecretStringProperty{
//   	ExcludeCharacters: jsii.String("excludeCharacters"),
//   	ExcludeLowercase: jsii.Boolean(false),
//   	ExcludeNumbers: jsii.Boolean(false),
//   	ExcludePunctuation: jsii.Boolean(false),
//   	ExcludeUppercase: jsii.Boolean(false),
//   	GenerateStringKey: jsii.String("generateStringKey"),
//   	IncludeSpace: jsii.Boolean(false),
//   	PasswordLength: jsii.Number(123),
//   	RequireEachIncludedType: jsii.Boolean(false),
//   	SecretStringTemplate: jsii.String("secretStringTemplate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html
//
type CfnSecret_GenerateSecretStringProperty struct {
	// A string of the characters that you don't want in the password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludecharacters
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Specifies whether to exclude lowercase letters from the password.
	//
	// If you don't include this switch, the password can contain lowercase letters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludelowercase
	//
	ExcludeLowercase interface{} `field:"optional" json:"excludeLowercase" yaml:"excludeLowercase"`
	// Specifies whether to exclude numbers from the password.
	//
	// If you don't include this switch, the password can contain numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludenumbers
	//
	ExcludeNumbers interface{} `field:"optional" json:"excludeNumbers" yaml:"excludeNumbers"`
	// Specifies whether to exclude the following punctuation characters from the password: `!
	//
	// " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~` . If you don't include this switch, the password can contain punctuation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludepunctuation
	//
	ExcludePunctuation interface{} `field:"optional" json:"excludePunctuation" yaml:"excludePunctuation"`
	// Specifies whether to exclude uppercase letters from the password.
	//
	// If you don't include this switch, the password can contain uppercase letters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-excludeuppercase
	//
	ExcludeUppercase interface{} `field:"optional" json:"excludeUppercase" yaml:"excludeUppercase"`
	// The JSON key name for the key/value pair, where the value is the generated password.
	//
	// This pair is added to the JSON structure specified by the `SecretStringTemplate` parameter. If you specify this parameter, then you must also specify `SecretStringTemplate` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-generatestringkey
	//
	GenerateStringKey *string `field:"optional" json:"generateStringKey" yaml:"generateStringKey"`
	// Specifies whether to include the space character.
	//
	// If you include this switch, the password can contain space characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-includespace
	//
	IncludeSpace interface{} `field:"optional" json:"includeSpace" yaml:"includeSpace"`
	// The length of the password.
	//
	// If you don't include this parameter, the default length is 32 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-passwordlength
	//
	PasswordLength *float64 `field:"optional" json:"passwordLength" yaml:"passwordLength"`
	// Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation.
	//
	// If you don't include this switch, the password contains at least one of every character type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-requireeachincludedtype
	//
	RequireEachIncludedType interface{} `field:"optional" json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
	// A template that the generated string must match.
	//
	// When you make a change to this property, a new secret version is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html#cfn-secretsmanager-secret-generatesecretstring-secretstringtemplate
	//
	SecretStringTemplate *string `field:"optional" json:"secretStringTemplate" yaml:"secretStringTemplate"`
}

