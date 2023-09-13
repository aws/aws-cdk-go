package awslambda


// Identification of an account (or organization) that is allowed to access a Lambda Layer Version.
//
// Example:
//   layer := lambda.NewLayerVersion(stack, jsii.String("MyLayer"), &LayerVersionProps{
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("layer-code"))),
//   	CompatibleRuntimes: []runtime{
//   		lambda.*runtime_NODEJS_LATEST(),
//   	},
//   	License: jsii.String("Apache-2.0"),
//   	Description: jsii.String("A layer to test the L2 construct"),
//   })
//
//   // To grant usage by other AWS accounts
//   layer.addPermission(jsii.String("remote-account-grant"), &LayerVersionPermission{
//   	AccountId: awsAccountId,
//   })
//
//   // To grant usage to all accounts in some AWS Ogranization
//   // layer.grantUsage({ accountId: '*', organizationId });
//
//   // To grant usage to all accounts in some AWS Ogranization
//   // layer.grantUsage({ accountId: '*', organizationId });
//   lambda.NewFunction(stack, jsii.String("MyLayeredLambda"), &FunctionProps{
//   	Code: lambda.NewInlineCode(jsii.String("foo")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.*runtime_NODEJS_LATEST(),
//   	Layers: []iLayerVersion{
//   		layer,
//   	},
//   })
//
type LayerVersionPermission struct {
	// The AWS Account id of the account that is authorized to use a Lambda Layer Version.
	//
	// The wild-card ``'*'`` can be
	// used to grant access to "any" account (or any account in an organization when ``organizationId`` is specified).
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ID of the AWS Organization to which the grant is restricted.
	//
	// Can only be specified if ``accountId`` is ``'*'``.
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

