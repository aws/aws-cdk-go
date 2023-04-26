# AWS Systems Manager Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Installation

Install the module:

```console
$ npm i @aws-cdk/aws-ssm
```

Import it into your code:

```go
import ssm "github.com/aws/aws-cdk-go/awscdk"
```

## Using existing SSM Parameters in your CDK app

You can reference existing SSM Parameter Store values that you want to use in
your CDK app by using `ssm.StringParameter.fromStringParameterAttributes`:

```go
// Retrieve the latest value of the non-secret parameter
// with name "/My/String/Parameter".
stringValue := ssm.StringParameter_FromStringParameterAttributes(this, jsii.String("MyValue"), &StringParameterAttributes{
	ParameterName: jsii.String("/My/Public/Parameter"),
}).StringValue
stringValueVersionFromToken := ssm.StringParameter_FromStringParameterAttributes(this, jsii.String("MyValueVersionFromToken"), &StringParameterAttributes{
	ParameterName: jsii.String("/My/Public/Parameter"),
	// parameter version from token
	Version: parameterVersion,
}).StringValue

// Retrieve a specific version of the secret (SecureString) parameter.
// 'version' is always required.
secretValue := ssm.StringParameter_FromSecureStringParameterAttributes(this, jsii.String("MySecureValue"), &SecureStringParameterAttributes{
	ParameterName: jsii.String("/My/Secret/Parameter"),
	Version: jsii.Number(5),
})
secretValueVersionFromToken := ssm.StringParameter_FromSecureStringParameterAttributes(this, jsii.String("MySecureValueVersionFromToken"), &SecureStringParameterAttributes{
	ParameterName: jsii.String("/My/Secret/Parameter"),
	// parameter version from token
	Version: parameterVersion,
})
```

## Creating new SSM Parameters in your CDK app

You can create either `ssm.StringParameter` or `ssm.StringListParameter`s in
a CDK app. These are public (not secret) values. Parameters of type
*SecureString* cannot be created directly from a CDK application; if you want
to provision secrets automatically, use Secrets Manager Secrets (see the
`@aws-cdk/aws-secretsmanager` package).

```go
ssm.NewStringParameter(this, jsii.String("Parameter"), &StringParameterProps{
	AllowedPattern: jsii.String(".*"),
	Description: jsii.String("The value Foo"),
	ParameterName: jsii.String("FooParameter"),
	StringValue: jsii.String("Foo"),
	Tier: ssm.ParameterTier_ADVANCED,
})
```

```go
// Create a new SSM Parameter holding a String
param := ssm.NewStringParameter(stack, jsii.String("StringParameter"), &StringParameterProps{
	// description: 'Some user-friendly description',
	// name: 'ParameterName',
	StringValue: jsii.String("Initial parameter value"),
})

// Grant read access to some Role
param.grantRead(role)

// Create a new SSM Parameter holding a StringList
listParameter := ssm.NewStringListParameter(stack, jsii.String("StringListParameter"), &StringListParameterProps{
	// description: 'Some user-friendly description',
	// name: 'ParameterName',
	StringListValue: []*string{
		jsii.String("Initial parameter value A"),
		jsii.String("Initial parameter value B"),
	},
})
```

When specifying an `allowedPattern`, the values provided as string literals
are validated against the pattern and an exception is raised if a value
provided does not comply.
