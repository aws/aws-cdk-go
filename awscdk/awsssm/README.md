# AWS Systems Manager Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Using existing SSM Parameters in your CDK app

You can reference existing SSM Parameter Store values that you want to use in
your CDK app by using `ssm.StringParameter.fromStringParameterAttributes`:

```go
// Example automatically generated from non-compiling source. May contain errors.
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

You can also reference an existing SSM Parameter Store value that matches an
[AWS specific parameter type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html#aws-specific-parameter-types):

```go
// Example automatically generated from non-compiling source. May contain errors.
ssm.stringParameter.valueForTypedStringParameterV2(stack, jsii.String("/My/Public/Parameter"), ssm.parameterValueType_AWS_EC2_IMAGE_ID)
```

To do the same for a SSM Parameter Store value that is stored as a list:

```go
// Example automatically generated from non-compiling source. May contain errors.
ssm.stringListParameter.valueForTypedListParameter(stack, jsii.String("/My/Public/Parameter"), ssm.parameterValueType_AWS_EC2_IMAGE_ID)
```

### Lookup existing parameters

You can also use an existing parameter by looking up the parameter from the AWS environment.
This method uses AWS API calls to lookup the value from SSM during synthesis.

```go
// Example automatically generated from non-compiling source. May contain errors.
stringValue := ssm.stringParameter.valueFromLookup(stack, jsii.String("/My/Public/Parameter"))
```

When using `valueFromLookup` an initial value of 'dummy-value-for-${parameterName}'
(`dummy-value-for-/My/Public/Parameter` in the above example)
is returned prior to the lookup being performed. This can lead to errors if you are using this
value in places that require a certain format. For example if you have stored the ARN for a SNS
topic in a SSM Parameter which you want to lookup and provide to `Topic.fromTopicArn()`

```go
arnLookup := ssm.stringParameter.valueFromLookup(this, jsii.String("/my/topic/arn"))
sns.topic.fromTopicArn(this, jsii.String("Topic"), arnLookup)
```

Initially `arnLookup` will be equal to `dummy-value-for-/my/topic/arn` which will cause
`Topic.fromTopicArn` to throw an error indicating that the value is not in `arn` format.

For these use cases you need to handle the `dummy-value` in your code. For example:

```go
arnLookup := ssm.stringParameter.valueFromLookup(this, jsii.String("/my/topic/arn"))
var arnLookupValue string
if arnLookup.includes(jsii.String("dummy-value")) {
	arnLookupValue = this.formatArn(&arnComponents{
		service: jsii.String("sns"),
		resource: jsii.String("topic"),
		resourceName: arnLookup,
	})
} else {
	arnLookupValue = arnLookup
}

sns.topic.fromTopicArn(this, jsii.String("Topic"), arnLookupValue)
```

Alternatively, if the property supports tokens you can convert the parameter value into a token
to be resolved *after* the lookup has been completed.

```go
arnLookup := ssm.stringParameter.valueFromLookup(this, jsii.String("/my/role/arn"))
iam.role.fromRoleArn(this, jsii.String("role"), awscdk.Lazy.string(map[string]produce{
	"produce": () => arnLookup,
}))
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
