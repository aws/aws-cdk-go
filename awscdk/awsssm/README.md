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
stringValue := ssm.stringParameter.fromStringParameterAttributes(this, jsii.String("MyValue"), &stringParameterAttributes{
	parameterName: jsii.String("/My/Public/Parameter"),
}).stringValue
stringValueVersionFromToken := ssm.stringParameter.fromStringParameterAttributes(this, jsii.String("MyValueVersionFromToken"), &stringParameterAttributes{
	parameterName: jsii.String("/My/Public/Parameter"),
	// parameter version from token
	version: parameterVersion,
}).stringValue

// Retrieve a specific version of the secret (SecureString) parameter.
// 'version' is always required.
secretValue := ssm.stringParameter.fromSecureStringParameterAttributes(this, jsii.String("MySecureValue"), &secureStringParameterAttributes{
	parameterName: jsii.String("/My/Secret/Parameter"),
	version: jsii.Number(5),
})
secretValueVersionFromToken := ssm.stringParameter.fromSecureStringParameterAttributes(this, jsii.String("MySecureValueVersionFromToken"), &secureStringParameterAttributes{
	parameterName: jsii.String("/My/Secret/Parameter"),
	// parameter version from token
	version: parameterVersion,
})
```

## Creating new SSM Parameters in your CDK app

You can create either `ssm.StringParameter` or `ssm.StringListParameter`s in
a CDK app. These are public (not secret) values. Parameters of type
*SecureString* cannot be created directly from a CDK application; if you want
to provision secrets automatically, use Secrets Manager Secrets (see the
`@aws-cdk/aws-secretsmanager` package).

```go
ssm.NewStringParameter(this, jsii.String("Parameter"), &stringParameterProps{
	allowedPattern: jsii.String(".*"),
	description: jsii.String("The value Foo"),
	parameterName: jsii.String("FooParameter"),
	stringValue: jsii.String("Foo"),
	tier: ssm.parameterTier_ADVANCED,
})
```

```go
// Create a new SSM Parameter holding a String
param := ssm.NewStringParameter(stack, jsii.String("StringParameter"), &stringParameterProps{
	// description: 'Some user-friendly description',
	// name: 'ParameterName',
	stringValue: jsii.String("Initial parameter value"),
})

// Grant read access to some Role
param.grantRead(role)

// Create a new SSM Parameter holding a StringList
listParameter := ssm.NewStringListParameter(stack, jsii.String("StringListParameter"), &stringListParameterProps{
	// description: 'Some user-friendly description',
	// name: 'ParameterName',
	stringListValue: []*string{
		jsii.String("Initial parameter value A"),
		jsii.String("Initial parameter value B"),
	},
})
```

When specifying an `allowedPattern`, the values provided as string literals
are validated against the pattern and an exception is raised if a value
provided does not comply.
