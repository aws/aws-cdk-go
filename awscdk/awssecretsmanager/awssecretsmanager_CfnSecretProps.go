package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSecret`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecretProps := &cfnSecretProps{
//   	description: jsii.String("description"),
//   	generateSecretString: &generateSecretStringProperty{
//   		excludeCharacters: jsii.String("excludeCharacters"),
//   		excludeLowercase: jsii.Boolean(false),
//   		excludeNumbers: jsii.Boolean(false),
//   		excludePunctuation: jsii.Boolean(false),
//   		excludeUppercase: jsii.Boolean(false),
//   		generateStringKey: jsii.String("generateStringKey"),
//   		includeSpace: jsii.Boolean(false),
//   		passwordLength: jsii.Number(123),
//   		requireEachIncludedType: jsii.Boolean(false),
//   		secretStringTemplate: jsii.String("secretStringTemplate"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	name: jsii.String("name"),
//   	replicaRegions: []interface{}{
//   		&replicaRegionProperty{
//   			region: jsii.String("region"),
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	secretString: jsii.String("secretString"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSecretProps struct {
	// The description of the secret.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that specifies how to generate a password to encrypt and store in the secret.
	//
	// Either `GenerateSecretString` or `SecretString` must have a value, but not both. They cannot both be empty.
	//
	// We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
	GenerateSecretString interface{} `field:"optional" json:"generateSecretString" yaml:"generateSecretString"`
	// The ARN, key ID, or alias of the AWS KMS key that Secrets Manager uses to encrypt the secret value in the secret.
	//
	// To use a AWS KMS key in a different account, use the key ARN or the alias ARN.
	//
	// If you don't specify this value, then Secrets Manager uses the key `aws/secretsmanager` . If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//
	// If the secret is in a different AWS account from the credentials calling the API, then you can't use `aws/secretsmanager` to encrypt the secret, and you must create and use a customer managed AWS KMS key.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the new secret.
	//
	// The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
	//
	// Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A custom type that specifies a `Region` and the `KmsKeyId` for a replica secret.
	ReplicaRegions interface{} `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// The text to encrypt and store in the secret.
	//
	// We recommend you use a JSON structure of key/value pairs for your secret value.
	//
	// Either `GenerateSecretString` or `SecretString` must have a value, but not both. They cannot both be empty. We recommend that you use the `GenerateSecretString` property to generate a random password.
	SecretString *string `field:"optional" json:"secretString" yaml:"secretString"`
	// A list of tags to attach to the secret.
	//
	// Each tag is a key and value pair of strings in a JSON text string, for example:
	//
	// `[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]`
	//
	// Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//
	// If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an `Access Denied` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2) .
	//
	// For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json) . If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//
	// The following restrictions apply to tags:
	//
	// - Maximum number of tags per secret: 50
	// - Maximum key length: 127 Unicode characters in UTF-8
	// - Maximum value length: 255 Unicode characters in UTF-8
	// - Tag keys and values are case sensitive.
	// - Do not use the `aws:` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	// - If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

