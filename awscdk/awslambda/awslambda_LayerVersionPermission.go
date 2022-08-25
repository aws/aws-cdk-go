package awslambda


// Identification of an account (or organization) that is allowed to access a Lambda Layer Version.
//
// Example:
//   layer := lambda.NewLayerVersion(stack, jsii.String("MyLayer"), &layerVersionProps{
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("layer-code"))),
//   	compatibleRuntimes: []runtime{
//   		lambda.*runtime_NODEJS_14_X(),
//   	},
//   	license: jsii.String("Apache-2.0"),
//   	description: jsii.String("A layer to test the L2 construct"),
//   })
//
//   // To grant usage by other AWS accounts
//   layer.addPermission(jsii.String("remote-account-grant"), &layerVersionPermission{
//   	accountId: awsAccountId,
//   })
//
//   // To grant usage to all accounts in some AWS Ogranization
//   // layer.grantUsage({ accountId: '*', organizationId });
//
//   // To grant usage to all accounts in some AWS Ogranization
//   // layer.grantUsage({ accountId: '*', organizationId });
//   lambda.NewFunction(stack, jsii.String("MyLayeredLambda"), &functionProps{
//   	code: lambda.NewInlineCode(jsii.String("foo")),
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.*runtime_NODEJS_14_X(),
//   	layers: []iLayerVersion{
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

