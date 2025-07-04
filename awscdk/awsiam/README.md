# AWS Identity and Access Management Construct Library

## Security and Safety Dev Guide

For a detailed guide on CDK security and safety please see the [CDK Security And
Safety Dev Guide](https://github.com/aws/aws-cdk/wiki/Security-And-Safety-Dev-Guide)

The guide will cover topics like:

* What permissions to extend to CDK deployments
* How to control the permissions of CDK deployments via IAM identities and policies
* How to use CDK to configure the IAM identities and policies of deployed applications
* Using Permissions Boundaries with CDK

## Overview

Define a role and add permissions to it. This will automatically create and
attach an IAM policy to the role:

```go
role := awscdk.NewRole(this, jsii.String("MyRole"), &RoleProps{
	AssumedBy: awscdk.NewServicePrincipal(jsii.String("sns.amazonaws.com")),
})

role.AddToPolicy(awscdk.NewPolicyStatement(&PolicyStatementProps{
	Resources: []*string{
		jsii.String("*"),
	},
	Actions: []*string{
		jsii.String("lambda:InvokeFunction"),
	},
}))
```

Define a policy and attach it to groups, users and roles. Note that it is possible to attach
the policy either by calling `xxx.attachInlinePolicy(policy)` or `policy.attachToXxx(xxx)`.

```go
user := awscdk.NewUser(this, jsii.String("MyUser"), &UserProps{
	Password: awscdk.SecretValue_PlainText(jsii.String("1234")),
})
group := awscdk.NewGroup(this, jsii.String("MyGroup"))

policy := awscdk.NewPolicy(this, jsii.String("MyPolicy"))
policy.AttachToUser(user)
group.attachInlinePolicy(policy)
```

Managed policies can be attached using `xxx.addManagedPolicy(ManagedPolicy.fromAwsManagedPolicyName(policyName))`:

```go
group := awscdk.NewGroup(this, jsii.String("MyGroup"))
group.AddManagedPolicy(awscdk.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AdministratorAccess")))
```

## Granting permissions to resources

Many of the AWS CDK resources have `grant*` methods that allow you to grant other resources access to that resource. As an example, the following code gives a Lambda function write permissions (Put, Update, Delete) to a DynamoDB table.

```go
var fn function
var table table


table.GrantWriteData(fn)
```

The more generic `grant` method allows you to give specific permissions to a resource:

```go
var fn function
var table table


table.Grant(fn, jsii.String("dynamodb:PutItem"))
```

The `grant*` methods accept an `IGrantable` object. This interface is implemented by IAM principal resources (groups, users and roles), policies, managed policies and resources that assume a role such as a Lambda function, EC2 instance or a Codebuild project.

You can find which `grant*` methods exist for a resource in the [AWS CDK API Reference](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-construct-library.html).

## Roles

Many AWS resources require *Roles* to operate. These Roles define the AWS API
calls an instance or other AWS service is allowed to make.

Creating Roles and populating them with the right permissions *Statements* is
a necessary but tedious part of setting up AWS infrastructure. In order to
help you focus on your business logic, CDK will take care of creating
roles and populating them with least-privilege permissions automatically.

All constructs that require Roles will create one for you if don't specify
one at construction time. Permissions will be added to that role
automatically if you associate the construct with other constructs from the
AWS Construct Library (for example, if you tell an *AWS CodePipeline* to trigger
an *AWS Lambda Function*, the Pipeline's Role will automatically get
`lambda:InvokeFunction` permissions on that particular Lambda Function),
or if you explicitly grant permissions using `grant` functions (see the
previous section).

### Opting out of automatic permissions management

You may prefer to manage a Role's permissions yourself instead of having the
CDK automatically manage them for you. This may happen in one of the
following cases:

* You don't like the permissions that CDK automatically generates and
  want to substitute your own set.
* The least-permissions policy that the CDK generates is becoming too
  big for IAM to store, and you need to add some wildcards to keep the
  policy size down.

To prevent constructs from updating your Role's policy, pass the object
returned by `myRole.withoutPolicyUpdates()` instead of `myRole` itself.

For example, to have an AWS CodePipeline *not* automatically add the required
permissions to trigger the expected targets, do the following:

```go
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("codepipeline.amazonaws.com")),
	// custom description if desired
	Description: jsii.String("This is a custom role..."),
})

codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
	// Give the Pipeline an immutable view of the Role
	Role: role.WithoutPolicyUpdates(),
})

// You now have to manage the Role policies yourself
role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
	},
	Resources: []*string{
	},
}))
```

### Using existing roles

If there are Roles in your account that have already been created which you
would like to use in your CDK application, you can use `Role.fromRoleArn` to
import them, as follows:

```go
role := iam.Role_FromRoleArn(this, jsii.String("Role"), jsii.String("arn:aws:iam::123456789012:role/MyExistingRole"), &FromRoleArnOptions{
	// Set 'mutable' to 'false' to use the role as-is and prevent adding new
	// policies to it. The default is 'true', which means the role may be
	// modified as part of the deployment.
	Mutable: jsii.Boolean(false),
})
```

If you want to lookup roles that actually exist in your account, you can use `Role.fromLookup()`.

```go
role := iam.Role_FromLookup(this, jsii.String("Role"), &RoleLookupOptions{
	RoleName: jsii.String("MyExistingRole"),
})
```

### Customizing role creation

It is best practice to allow CDK to manage IAM roles and permissions. You can prevent CDK from
creating roles by using the `customizeRoles` method for special cases. One such case is using CDK in
an environment where role creation is not allowed or needs to be managed through a process outside
of the CDK application.

An example of how to opt in to this behavior is below:

```go
var stack stack

iam.Role_CustomizeRoles(stack)
```

CDK will not create any IAM roles or policies with the `stack` scope. `cdk synth` will fail and
it will generate a policy report to the cloud assembly (i.e. cdk.out). The `iam-policy-report.txt`
report will contain a list of IAM roles and associated permissions that would have been created.
This report can be used to create the roles with the appropriate permissions outside of
the CDK application.

Once the missing roles have been created, their names can be added to the `usePrecreatedRoles`
property, like shown below:

```go
var app app

stack := awscdk.Newstack(app, jsii.String("MyStack"))
iam.Role_CustomizeRoles(this, &CustomizeRolesOptions{
	UsePrecreatedRoles: map[string]*string{
		"MyStack/MyRole": jsii.String("my-precreated-role-name"),
	},
})

iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("sns.amazonaws.com")),
})
```

If any IAM policies reference deploy time values (i.e. ARN of a resource that hasn't been created
yet) you will have to modify the generated report to be more generic. For example, given the
following CDK code:

```go
var app app

stack := awscdk.Newstack(app, jsii.String("MyStack"))
iam.Role_CustomizeRoles(stack)

fn := lambda.NewFunction(this, jsii.String("MyLambda"), &FunctionProps{
	Code: lambda.NewInlineCode(jsii.String("foo")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
})

bucket := s3.NewBucket(this, jsii.String("Bucket"))
bucket.GrantRead(fn)
```

The following report will be generated.

```txt
<missing role> (MyStack/MyLambda/ServiceRole)

AssumeRole Policy:
[
  {
    "Action": "sts:AssumeRole",
    "Effect": "Allow",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    }
  }
]

Managed Policy ARNs:
[
  "arn:(PARTITION):iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
]

Managed Policies Statements:
NONE

Identity Policy Statements:
[
  {
    "Action": [
      "s3:GetObject*",
      "s3:GetBucket*",
      "s3:List*"
    ],
    "Effect": "Allow",
    "Resource": [
      "(MyStack/Bucket/Resource.Arn)",
      "(MyStack/Bucket/Resource.Arn)/*"
    ]
  }
]
```

You would then need to create the role with the inline & managed policies in the report and then
come back and update the `customizeRoles` with the role name.

```go
var app app

stack := awscdk.Newstack(app, jsii.String("MyStack"))
iam.Role_CustomizeRoles(this, &CustomizeRolesOptions{
	UsePrecreatedRoles: map[string]*string{
		"MyStack/MyLambda/ServiceRole": jsii.String("my-role-name"),
	},
})
```

For more information on configuring permissions see the [Security And Safety Dev
Guide](https://github.com/aws/aws-cdk/wiki/Security-And-Safety-Dev-Guide)

#### Policy report generation

When `customizeRoles` is used, the `iam-policy-report.txt` report will contain a list
of IAM roles and associated permissions that would have been created. This report is
generated in an attempt to resolve and replace any references with a more user-friendly
value.

The following are some examples of the value that will appear in the report:

```json
"Resource": {
  "Fn::Join": [
    "",
    [
      "arn:",
      {
        "Ref": "AWS::Partition"
      },
      ":iam::",
      {
        "Ref": "AWS::AccountId"
      },
      ":role/Role"
    ]
  ]
}
```

The policy report will instead get:

```json
"Resource": "arn:(PARTITION):iam::(ACCOUNT):role/Role"
```

If IAM policy is referencing a resource attribute:

```json
"Resource": [
  {
    "Fn::GetAtt": [
      "SomeResource",
      "Arn"
    ]
  },
  {
    "Ref": "AWS::NoValue",
  }
]
```

The policy report will instead get:

```json
"Resource": [
  "(Path/To/SomeResource.Arn)"
  "(NOVALUE)"
]
```

The following pseudo parameters will be converted:

1. `{ 'Ref': 'AWS::AccountId' }` -> `(ACCOUNT)
2. `{ 'Ref': 'AWS::Partition' }` -> `(PARTITION)
3. `{ 'Ref': 'AWS::Region' }` -> `(REGION)
4. `{ 'Ref': 'AWS::NoValue' }` -> `(NOVALUE)

#### Generating a permissions report

It is also possible to generate the report *without* preventing the role/policy creation.

```go
var stack stack

iam.Role_CustomizeRoles(this, &CustomizeRolesOptions{
	PreventSynthesis: jsii.Boolean(false),
})
```

## Configuring an ExternalId

If you need to create Roles that will be assumed by third parties, it is generally a good idea to [require an `ExternalId`
to assume them](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html).  Configuring
an `ExternalId` works like this:

```go
role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
	ExternalIds: []*string{
		jsii.String("SUPPLY-ME"),
	},
})
```

## SourceArn and SourceAccount

If you need to create resource policies using `aws:SourceArn` and `aws:SourceAccount` for cross-service resource access,
use `addSourceArnCondition` and `addSourceAccountCondition` to create the conditions.

See [Cross-service confused deputy prevention for more details](https://docs.aws.amazon.com/IAM/latest/UserGuide/confused-deputy.html#cross-service-confused-deputy-prevention).

## Principals vs Identities

When we say *Principal*, we mean an entity you grant permissions to. This
entity can be an AWS Service, a Role, or something more abstract such as "all
users in this account" or even "all users in this organization". An
*Identity* is an IAM representing a single IAM entity that can have
a policy attached, one of `Role`, `User`, or `Group`.

## IAM Principals

When defining policy statements as part of an AssumeRole policy or as part of a
resource policy, statements would usually refer to a specific IAM principal
under `Principal`.

IAM principals are modeled as classes that derive from the `iam.PolicyPrincipal`
abstract class. Principal objects include principal type (string) and value
(array of string), optional set of conditions and the action that this principal
requires when it is used in an assume role policy document.

To add a principal to a policy statement you can either use the abstract
`statement.addPrincipal`, one of the concrete `addXxxPrincipal` methods:

* `addAwsPrincipal`, `addArnPrincipal` or `new ArnPrincipal(arn)` for `{ "AWS": arn }`
* `addAwsAccountPrincipal` or `new AccountPrincipal(accountId)` for `{ "AWS": account-arn }`
* `addServicePrincipal` or `new ServicePrincipal(service)` for `{ "Service": service }`
* `addAccountRootPrincipal` or `new AccountRootPrincipal()` for `{ "AWS": { "Ref: "AWS::AccountId" } }`
* `addCanonicalUserPrincipal` or `new CanonicalUserPrincipal(id)` for `{ "CanonicalUser": id }`
* `addFederatedPrincipal` or `new FederatedPrincipal(federated, conditions, assumeAction)` for
  `{ "Federated": arn }` and a set of optional conditions and the assume role action to use.
* `addAnyPrincipal` or `new AnyPrincipal` for `{ "AWS": "*" }`

If multiple principals are added to the policy statement, they will be merged together:

```go
statement := iam.NewPolicyStatement()
statement.AddServicePrincipal(jsii.String("cloudwatch.amazonaws.com"))
statement.AddServicePrincipal(jsii.String("ec2.amazonaws.com"))
statement.AddArnPrincipal(jsii.String("arn:aws:boom:boom"))
```

Will result in:

```json
{
  "Principal": {
    "Service": [ "cloudwatch.amazonaws.com", "ec2.amazonaws.com" ],
    "AWS": "arn:aws:boom:boom"
  }
}
```

The `CompositePrincipal` class can also be used to define complex principals, for example:

```go
role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	AssumedBy: iam.NewCompositePrincipal(
	iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	iam.NewAccountPrincipal(jsii.String("1818188181818187272"))),
})
```

The `PrincipalWithConditions` class can be used to add conditions to a
principal, especially those that don't take a `conditions` parameter in their
constructor. The `principal.withConditions()` method can be used to create a
`PrincipalWithConditions` from an existing principal, for example:

```go
principal := iam.NewAccountPrincipal(jsii.String("123456789000")).WithConditions(map[string]interface{}{
	"StringEquals": map[string]*string{
		"foo": jsii.String("baz"),
	},
})
```

> NOTE: If you need to define an IAM condition that uses a token (such as a
> deploy-time attribute of another resource) in a JSON map key, use `CfnJson` to
> render this condition. See [this test](./test/integ.condition-with-ref.ts) for
> an example.

The `WebIdentityPrincipal` class can be used as a principal for web identities like
Cognito, Amazon, Google or Facebook, for example:

```go
principal := iam.NewWebIdentityPrincipal(jsii.String("cognito-identity.amazonaws.com"), map[string]interface{}{
	"StringEquals": map[string]*string{
		"cognito-identity.amazonaws.com:aud": jsii.String("us-east-2:12345678-abcd-abcd-abcd-123456"),
	},
	"ForAnyValue:StringLike": map[string]*string{
		"cognito-identity.amazonaws.com:amr": jsii.String("unauthenticated"),
	},
})
```

If your identity provider is configured to assume a Role with [session
tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html), you
need to call `.withSessionTags()` to add the required permissions to the Role's
policy document:

```go
iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewWebIdentityPrincipal(jsii.String("cognito-identity.amazonaws.com"), map[string]interface{}{
		"StringEquals": map[string]*string{
			"cognito-identity.amazonaws.com:aud": jsii.String("us-east-2:12345678-abcd-abcd-abcd-123456"),
		},
		"ForAnyValue:StringLike": map[string]*string{
			"cognito-identity.amazonaws.com:amr": jsii.String("unauthenticated"),
		},
	}).WithSessionTags(),
})
```

### Granting a principal permission to assume a role

A principal can be granted permission to assume a role using `grantAssumeRole`.

Note that this does not apply to service principals or account principals as they must be added to the role trust policy via `assumeRolePolicy`.

```go
user := iam.NewUser(this, jsii.String("user"))
role := iam.NewRole(this, jsii.String("role"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(this.Account),
})

role.GrantAssumeRole(user)
```

### Granting service and account principals permission to assume a role

Service principals and account principals can be granted permission to assume a role using `assumeRolePolicy` which modifies the role trust policy.

```go
role := iam.NewRole(this, jsii.String("role"), &RoleProps{
	AssumedBy: iam.NewAccountPrincipal(this.Account),
})

role.AssumeRolePolicy.AddStatements(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("sts:AssumeRole"),
	},
	Principals: []iPrincipal{
		iam.NewAccountPrincipal(jsii.String("123456789")),
		iam.NewServicePrincipal(jsii.String("beep-boop.amazonaws.com")),
	},
}))
```

### Fixing the synthesized service principle for services that do not follow the IAM Pattern

In some cases, certain AWS services may not use the standard `<service>.amazonaws.com` pattern for their service principals. For these services, you can define the ServicePrincipal as following where the provided service principle name will be used as is without any changing.

```go
sp := iam.ServicePrincipal_FromStaticServicePrincipleName(jsii.String("elasticmapreduce.amazonaws.com.cn"))
```

This principle can use as normal in defining any role, for example:

```go
emrServiceRole := iam.NewRole(this, jsii.String("EMRServiceRole"), &RoleProps{
	AssumedBy: iam.ServicePrincipal_FromStaticServicePrincipleName(jsii.String("elasticmapreduce.amazonaws.com.cn")),
	ManagedPolicies: []iManagedPolicy{
		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AmazonElasticMapReduceRole")),
	},
})
```

## Parsing JSON Policy Documents

The `PolicyDocument.fromJson` and `PolicyStatement.fromJson` static methods can be used to parse JSON objects. For example:

```go
policyDocument := map[string]interface{}{
	"Version": jsii.String("2012-10-17"),
	"Statement": []interface{}{
		map[string]interface{}{
			"Sid": jsii.String("FirstStatement"),
			"Effect": jsii.String("Allow"),
			"Action": []*string{
				jsii.String("iam:ChangePassword"),
			},
			"Resource": []*string{
				jsii.String("*"),
			},
		},
		map[string]interface{}{
			"Sid": jsii.String("SecondStatement"),
			"Effect": jsii.String("Allow"),
			"Action": []*string{
				jsii.String("s3:ListAllMyBuckets"),
			},
			"Resource": []*string{
				jsii.String("*"),
			},
		},
		map[string]interface{}{
			"Sid": jsii.String("ThirdStatement"),
			"Effect": jsii.String("Allow"),
			"Action": []*string{
				jsii.String("s3:List*"),
				jsii.String("s3:Get*"),
			},
			"Resource": []*string{
				jsii.String("arn:aws:s3:::confidential-data"),
				jsii.String("arn:aws:s3:::confidential-data/*"),
			},
			"Condition": map[string]map[string]*string{
				"Bool": map[string]*string{
					"aws:MultiFactorAuthPresent": jsii.String("true"),
				},
			},
		},
	},
}

customPolicyDocument := iam.PolicyDocument_FromJson(policyDocument)

// You can pass this document as an initial document to a ManagedPolicy
// or inline Policy.
newManagedPolicy := iam.NewManagedPolicy(this, jsii.String("MyNewManagedPolicy"), &ManagedPolicyProps{
	Document: customPolicyDocument,
})
newPolicy := iam.NewPolicy(this, jsii.String("MyNewPolicy"), &PolicyProps{
	Document: customPolicyDocument,
})
```

## Permissions Boundaries

[Permissions
Boundaries](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
can be used as a mechanism to prevent privilege escalation by creating new
`Role`s. Permissions Boundaries are a Managed Policy, attached to Roles or
Users, that represent the *maximum* set of permissions they can have. The
effective set of permissions of a Role (or User) will be the intersection of
the Identity Policy and the Permissions Boundary attached to the Role (or
User). Permissions Boundaries are typically created by account
Administrators, and their use on newly created `Role`s will be enforced by
IAM policies.

### Bootstrap Permissions Boundary

If a permissions boundary has been enforced as part of CDK bootstrap, all IAM
Roles and Users that are created as part of the CDK application must be created
with the permissions boundary attached. The most common scenario will be to
apply the enforced permissions boundary to the entire CDK app. This can be done
either by adding the value to `cdk.json` or directly in the `App` constructor.

For example if your organization has created and is enforcing a permissions
boundary with the name
`cdk-${Qualifier}-PermissionsBoundary`

```json
{
  "context": {
     "@aws-cdk/core:permissionsBoundary": {
	   "name": "cdk-${Qualifier}-PermissionsBoundary"
	 }
  }
}
```

OR

```go
awscdk.NewApp(&AppProps{
	Context: map[string]interface{}{
		awscdk.PERMISSIONS_BOUNDARY_CONTEXT_KEY: map[string]*string{
			"name": jsii.String("cdk-${Qualifier}-PermissionsBoundary"),
		},
	},
})
```

Another scenario might be if your organization enforces different permissions
boundaries for different environments. For example your CDK application may have

* `DevStage` that deploys to a personal dev environment where you have elevated
  privileges
* `BetaStage` that deploys to a beta environment which and has a relaxed
  permissions boundary
* `GammaStage` that deploys to a gamma environment which has the prod
  permissions boundary
* `ProdStage` that deploys to the prod environment and has the prod permissions
  boundary

```go
var app app


awscdk.NewStage(app, jsii.String("DevStage"))

awscdk.NewStage(app, jsii.String("BetaStage"), &StageProps{
	PermissionsBoundary: awscdk.PermissionsBoundary_FromName(jsii.String("beta-permissions-boundary")),
})

awscdk.NewStage(app, jsii.String("GammaStage"), &StageProps{
	PermissionsBoundary: awscdk.PermissionsBoundary_*FromName(jsii.String("prod-permissions-boundary")),
})

awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
	PermissionsBoundary: awscdk.PermissionsBoundary_*FromName(jsii.String("prod-permissions-boundary")),
})
```

The provided name can include placeholders for the partition, region, qualifier, and account
These placeholders will be replaced with the actual values if available. This requires
that the Stack has the environment specified, it does not work with environment.

* '${AWS::Partition}'
* '${AWS::Region}'
* '${AWS::AccountId}'
* '${Qualifier}'

```go
var app app


prodStage := awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
	PermissionsBoundary: awscdk.PermissionsBoundary_FromName(jsii.String("cdk-${Qualifier}-PermissionsBoundary-${AWS::AccountId}-${AWS::Region}")),
})

awscdk.Newstack(prodStage, jsii.String("ProdStack"), &StackProps{
	Synthesizer: awscdk.NewDefaultStackSynthesizer(&DefaultStackSynthesizerProps{
		Qualifier: jsii.String("custom"),
	}),
})
```

For more information on configuring permissions see the [Security And Safety Dev
Guide](https://github.com/aws/aws-cdk/wiki/Security-And-Safety-Dev-Guide)

### Custom Permissions Boundary

It is possible to attach Permissions Boundaries to all Roles created in a construct
tree all at once:

```go
// Directly apply the boundary to a Role you create
var role role

// Apply the boundary to an Role that was implicitly created for you
var fn function

// Remove a Permissions Boundary that is inherited, for example from the Stack level
var customResource customResource
// This imports an existing policy.
boundary := iam.ManagedPolicy_FromManagedPolicyArn(this, jsii.String("Boundary"), jsii.String("arn:aws:iam::123456789012:policy/boundary"))

// This creates a new boundary
boundary2 := iam.NewManagedPolicy(this, jsii.String("Boundary2"), &ManagedPolicyProps{
	Statements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Effect: iam.Effect_DENY,
			Actions: []*string{
				jsii.String("iam:*"),
			},
			Resources: []*string{
				jsii.String("*"),
			},
		}),
	},
})
iam.PermissionsBoundary_Of(role).Apply(boundary)
iam.PermissionsBoundary_Of(fn).Apply(boundary)

// Apply the boundary to all Roles in a stack
iam.PermissionsBoundary_Of(this).Apply(boundary)
iam.PermissionsBoundary_Of(customResource).Clear()
```

## OpenID Connect Providers

OIDC identity providers are entities in IAM that describe an external identity
provider (IdP) service that supports the [OpenID Connect](http://openid.net/connect) (OIDC) standard, such
as Google or Salesforce. You use an IAM OIDC identity provider when you want to
establish trust between an OIDC-compatible IdP and your AWS account. This is
useful when creating a mobile app or web application that requires access to AWS
resources, but you don't want to create custom sign-in code or manage your own
user identities. For more information about this scenario, see [About Web
Identity Federation](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html) and the relevant documentation in the [Amazon Cognito
Identity Pools Developer Guide](https://docs.aws.amazon.com/cognito/latest/developerguide/open-id.html).

The following examples defines an OpenID Connect provider. Two client IDs
(audiences) are will be able to send authentication requests to
[https://openid/connect](https://openid/connect).

It is recommended to use the new `OidcProviderNative` which native CloudFormation resource `AWS::IAM::OIDCProvider` over the old `OpenIdConnectProvider` which uses a custom resource.

```go
nativeProvider := iam.NewOidcProviderNative(this, jsii.String("MyProvider"), &OidcProviderNativeProps{
	Url: jsii.String("https://openid/connect"),
	ClientIds: []*string{
		jsii.String("myclient1"),
		jsii.String("myclient2"),
	},
	Thumbprints: []*string{
		jsii.String("aa00aa1122aa00aa1122aa00aa1122aa00aa1122"),
	},
})
```

For the new `OidcProviderNative`, you must provide at least one thumbprint when creating an IAM OIDC
provider. For example, assume that the OIDC provider is server.example.com
and the provider stores its keys at
https://keys.server.example.com/openid-connect. In that case, the
thumbprint string would be the hex-encoded SHA-1 hash value of the
certificate used by https://keys.server.example.com.

The server certificate thumbprint is the hex-encoded SHA-1 hash value of
the X.509 certificate used by the domain where the OpenID Connect provider
makes its keys available. It is always a 40-character string.

Typically this list includes only one entry. However, IAM lets you have up
to five thumbprints for an OIDC provider. This lets you maintain multiple
thumbprints if the identity provider is rotating certificates.

Obtain the thumbprint of the root certificate authority from the provider's
server as described in https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html

The older `OpenIdConnectProvider` is still supported but it is recommended to use the new `OidcProviderNative` instead.

```go
provider := iam.NewOpenIdConnectProvider(this, jsii.String("MyProvider"), &OpenIdConnectProviderProps{
	Url: jsii.String("https://openid/connect"),
	ClientIds: []*string{
		jsii.String("myclient1"),
		jsii.String("myclient2"),
	},
})
```

For the older `OpenIdConnectProvider`, you can specify an optional list of `thumbprints`. If not specified, the
thumbprint of the root certificate authority (CA) will automatically be obtained
from the host as described
[here](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html).

By default, the custom resource enforces strict security practices by rejecting
any unauthorized connections when downloading CA thumbprints from the issuer URL.
If you need to connect to an unauthorized OIDC identity provider and understand the
implications, you can disable this behavior by setting the feature flag
`IAM_OIDC_REJECT_UNAUTHORIZED_CONNECTIONS` to `false` in your `cdk.context.json`
or `cdk.json`. Visit [CDK Feature Flag](https://docs.aws.amazon.com/cdk/v2/guide/featureflags.html)
for more information on how to configure feature flags.

Once you define an OpenID connect provider, you can use it with AWS services
that expect an IAM OIDC provider. For example, when you define an [Amazon
Cognito identity
pool](https://docs.aws.amazon.com/cognito/latest/developerguide/open-id.html)
you can reference the provider's ARN as follows:

```go
import cognito "github.com/aws/aws-cdk-go/awscdk"

var myProvider openIdConnectProvider

cognito.NewCfnIdentityPool(this, jsii.String("IdentityPool"), &CfnIdentityPoolProps{
	OpenIdConnectProviderArns: []*string{
		myProvider.OpenIdConnectProviderArn,
	},
	// And the other properties for your identity pool
	AllowUnauthenticatedIdentities: jsii.Boolean(false),
})
```

The `OpenIdConnectPrincipal` class can be used as a principal used with a `OpenIdConnectProvider`, for example:

```go
provider := iam.NewOpenIdConnectProvider(this, jsii.String("MyProvider"), &OpenIdConnectProviderProps{
	Url: jsii.String("https://openid/connect"),
	ClientIds: []*string{
		jsii.String("myclient1"),
		jsii.String("myclient2"),
	},
})
principal := iam.NewOpenIdConnectPrincipal(provider)
```

## SAML provider

An IAM SAML 2.0 identity provider is an entity in IAM that describes an external
identity provider (IdP) service that supports the SAML 2.0 (Security Assertion
Markup Language 2.0) standard. You use an IAM identity provider when you want
to establish trust between a SAML-compatible IdP such as Shibboleth or Active
Directory Federation Services and AWS, so that users in your organization can
access AWS resources. IAM SAML identity providers are used as principals in an
IAM trust policy.

```go
iam.NewSamlProvider(this, jsii.String("Provider"), &SamlProviderProps{
	MetadataDocument: iam.SamlMetadataDocument_FromFile(jsii.String("/path/to/saml-metadata-document.xml")),
})
```

The `SamlPrincipal` class can be used as a principal with a `SamlProvider`:

```go
provider := iam.NewSamlProvider(this, jsii.String("Provider"), &SamlProviderProps{
	MetadataDocument: iam.SamlMetadataDocument_FromFile(jsii.String("/path/to/saml-metadata-document.xml")),
})
principal := iam.NewSamlPrincipal(provider, map[string]interface{}{
	"StringEquals": map[string]*string{
		"SAML:iss": jsii.String("issuer"),
	},
})
```

When creating a role for programmatic and AWS Management Console access, use the `SamlConsolePrincipal`
class:

```go
provider := iam.NewSamlProvider(this, jsii.String("Provider"), &SamlProviderProps{
	MetadataDocument: iam.SamlMetadataDocument_FromFile(jsii.String("/path/to/saml-metadata-document.xml")),
})
iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewSamlConsolePrincipal(provider),
})
```

## Users

IAM manages users for your AWS account. To create a new user:

```go
user := iam.NewUser(this, jsii.String("MyUser"))
```

To import an existing user by name [with path](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names):

```go
user := iam.User_FromUserName(this, jsii.String("MyImportedUserByName"), jsii.String("johnsmith"))
```

To import an existing user by ARN:

```go
user := iam.User_FromUserArn(this, jsii.String("MyImportedUserByArn"), jsii.String("arn:aws:iam::123456789012:user/johnsmith"))
```

To import an existing user by attributes:

```go
user := iam.User_FromUserAttributes(this, jsii.String("MyImportedUserByAttributes"), &UserAttributes{
	UserArn: jsii.String("arn:aws:iam::123456789012:user/johnsmith"),
})
```

### Access Keys

The ability for a user to make API calls via the CLI or an SDK is enabled by the user having an
access key pair. To create an access key:

```go
user := iam.NewUser(this, jsii.String("MyUser"))
accessKey := iam.NewAccessKey(this, jsii.String("MyAccessKey"), &AccessKeyProps{
	User: user,
})
```

You can force CloudFormation to rotate the access key by providing a monotonically increasing `serial`
property. Simply provide a higher serial value than any number used previously:

```go
user := iam.NewUser(this, jsii.String("MyUser"))
accessKey := iam.NewAccessKey(this, jsii.String("MyAccessKey"), &AccessKeyProps{
	User: user,
	Serial: jsii.Number(1),
})
```

An access key may only be associated with a single user and cannot be "moved" between users. Changing
the user associated with an access key replaces the access key (and its ID and secret value).

## Groups

An IAM user group is a collection of IAM users. User groups let you specify permissions for multiple users.

```go
group := iam.NewGroup(this, jsii.String("MyGroup"))
```

To import an existing group by ARN:

```go
group := iam.Group_FromGroupArn(this, jsii.String("MyImportedGroupByArn"), jsii.String("arn:aws:iam::account-id:group/group-name"))
```

To import an existing group by name [with path](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names):

```go
group := iam.Group_FromGroupName(this, jsii.String("MyImportedGroupByName"), jsii.String("group-name"))
```

To add a user to a group (both for a new and imported user/group):

```go
user := iam.NewUser(this, jsii.String("MyUser")) // or User.fromUserName(this, 'User', 'johnsmith');
group := iam.NewGroup(this, jsii.String("MyGroup")) // or Group.fromGroupArn(this, 'Group', 'arn:aws:iam::account-id:group/group-name');

user.AddToGroup(group)
// or
group.addUser(user)
```

## Instance Profiles

An IAM instance profile is a container for an IAM role that you can use to pass role information to an EC2 instance when the instance starts. By default, an instance profile must be created with a role:

```go
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
})

instanceProfile := iam.NewInstanceProfile(this, jsii.String("InstanceProfile"), &InstanceProfileProps{
	Role: Role,
})
```

An instance profile can also optionally be created with an instance profile name and/or a [path](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) to the instance profile:

```go
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
})

instanceProfile := iam.NewInstanceProfile(this, jsii.String("InstanceProfile"), &InstanceProfileProps{
	Role: Role,
	InstanceProfileName: jsii.String("MyInstanceProfile"),
	Path: jsii.String("/sample/path/"),
})
```

To import an existing instance profile by name:

```go
instanceProfile := iam.InstanceProfile_FromInstanceProfileName(this, jsii.String("ImportedInstanceProfile"), jsii.String("MyInstanceProfile"))
```

To import an existing instance profile by ARN:

```go
instanceProfile := iam.InstanceProfile_FromInstanceProfileArn(this, jsii.String("ImportedInstanceProfile"), jsii.String("arn:aws:iam::account-id:instance-profile/MyInstanceProfile"))
```

To import an existing instance profile with an associated role:

```go
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
})

instanceProfile := iam.InstanceProfile_FromInstanceProfileAttributes(this, jsii.String("ImportedInstanceProfile"), &InstanceProfileAttributes{
	InstanceProfileArn: jsii.String("arn:aws:iam::account-id:instance-profile/MyInstanceProfile"),
	Role: Role,
})
```

## Features

* Policy name uniqueness is enforced. If two policies by the same name are attached to the same
  principal, the attachment will fail.
* Policy names are not required - the CDK logical ID will be used and ensured to be unique.
* Policies are validated during synthesis to ensure that they have actions, and that policies
  attached to IAM principals specify relevant resources, while policies attached to resources
  specify which IAM principals they apply to.
