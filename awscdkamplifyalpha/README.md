# AWS Amplify Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

The AWS Amplify Console provides a Git-based workflow for deploying and hosting fullstack serverless web applications. A fullstack serverless app consists of a backend built with cloud resources such as GraphQL or REST APIs, file and data storage, and a frontend built with single page application frameworks such as React, Angular, Vue, or Gatsby.

## Setting up an app with branches, custom rules and a domain

To set up an Amplify Console app, define an `App`:

```go
import codebuild "github.com/aws/aws-cdk-go/awscdk"


amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
		owner: jsii.String("<user>"),
		repository: jsii.String("<repo>"),
		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	}),
	buildSpec: codebuild.buildSpec.fromObjectToYaml(map[string]interface{}{
		// Alternatively add a `amplify.yml` to the repo
		"version": jsii.String("1.0"),
		"frontend": map[string]map[string]map[string][]*string{
			"phases": map[string]map[string][]*string{
				"preBuild": map[string][]*string{
					"commands": []*string{
						jsii.String("yarn"),
					},
				},
				"build": map[string][]*string{
					"commands": []*string{
						jsii.String("yarn build"),
					},
				},
			},
			"artifacts": map[string]interface{}{
				"baseDirectory": jsii.String("public"),
				"files": -jsii.String("**/*"),
			},
		},
	}),
})
```

To connect your `App` to GitLab, use the `GitLabSourceCodeProvider`:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
	sourceCodeProvider: amplify.NewGitLabSourceCodeProvider(&gitLabSourceCodeProviderProps{
		owner: jsii.String("<user>"),
		repository: jsii.String("<repo>"),
		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-gitlab-token")),
	}),
})
```

To connect your `App` to CodeCommit, use the `CodeCommitSourceCodeProvider`:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"


repository := codecommit.NewRepository(this, jsii.String("Repo"), &repositoryProps{
	repositoryName: jsii.String("my-repo"),
})

amplifyApp := amplify.NewApp(this, jsii.String("App"), &appProps{
	sourceCodeProvider: amplify.NewCodeCommitSourceCodeProvider(&codeCommitSourceCodeProviderProps{
		repository: repository,
	}),
})
```

The IAM role associated with the `App` will automatically be granted the permission
to pull the CodeCommit repository.

Add branches:

```go
var amplifyApp app


master := amplifyApp.addBranch(jsii.String("master")) // `id` will be used as repo branch name
dev := amplifyApp.addBranch(jsii.String("dev"), &branchOptions{
	performanceMode: jsii.Boolean(true),
})
dev.addEnvironment(jsii.String("STAGE"), jsii.String("dev"))
```

Auto build and pull request preview are enabled by default.

Add custom rules for redirection:

```go
var amplifyApp app

amplifyApp.addCustomRule(map[string]interface{}{
	"source": jsii.String("/docs/specific-filename.html"),
	"target": jsii.String("/documents/different-filename.html"),
	"status": amplify.RedirectStatus_TEMPORARY_REDIRECT,
})
```

When working with a single page application (SPA), use the
`CustomRule.SINGLE_PAGE_APPLICATION_REDIRECT` to set up a 200
rewrite for all files to `index.html` except for the following
file extensions: css, gif, ico, jpg, js, png, txt, svg, woff,
ttf, map, json, webmanifest.

```go
var mySinglePageApp app


mySinglePageApp.addCustomRule(amplify.customRule_SINGLE_PAGE_APPLICATION_REDIRECT())
```

Add a domain and map sub domains to branches:

```go
var amplifyApp app
var master branch
var dev branch


domain := amplifyApp.addDomain(jsii.String("example.com"), &domainOptions{
	enableAutoSubdomain: jsii.Boolean(true),
	 // in case subdomains should be auto registered for branches
	autoSubdomainCreationPatterns: []*string{
		jsii.String("*"),
		jsii.String("pr*"),
	},
})
domain.mapRoot(master) // map master branch to domain root
domain.mapSubDomain(master, jsii.String("www"))
domain.mapSubDomain(dev)
```

## Restricting access

Password protect the app with basic auth by specifying the `basicAuth` prop.

Use `BasicAuth.fromCredentials` when referencing an existing secret:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
		owner: jsii.String("<user>"),
		repository: jsii.String("<repo>"),
		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	}),
	basicAuth: amplify.basicAuth.fromCredentials(jsii.String("username"), awscdk.SecretValue.secretsManager(jsii.String("my-github-token"))),
})
```

Use `BasicAuth.fromGeneratedPassword` to generate a password in Secrets Manager:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
		owner: jsii.String("<user>"),
		repository: jsii.String("<repo>"),
		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	}),
	basicAuth: amplify.basicAuth.fromGeneratedPassword(jsii.String("username")),
})
```

Basic auth can be added to specific branches:

```go
var amplifyApp app

amplifyApp.addBranch(jsii.String("feature/next"), &branchOptions{
	basicAuth: amplify.basicAuth.fromGeneratedPassword(jsii.String("username")),
})
```

## Automatically creating and deleting branches

Use the `autoBranchCreation` and `autoBranchDeletion` props to control creation/deletion
of branches:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
		owner: jsii.String("<user>"),
		repository: jsii.String("<repo>"),
		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	}),
	autoBranchCreation: &autoBranchCreation{
		 // Automatically connect branches that match a pattern set
		patterns: []*string{
			jsii.String("feature/*"),
			jsii.String("test/*"),
		},
	},
	autoBranchDeletion: jsii.Boolean(true),
})
```

## Adding custom response headers

Use the `customResponseHeaders` prop to configure custom response headers for an Amplify app:

```go
amplifyApp := amplify.NewApp(this, jsii.String("App"), &appProps{
	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
		owner: jsii.String("<user>"),
		repository: jsii.String("<repo>"),
		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
	}),
	customResponseHeaders: []customResponseHeader{
		&customResponseHeader{
			pattern: jsii.String("*.json"),
			headers: map[string]*string{
				"custom-header-name-1": jsii.String("custom-header-value-1"),
				"custom-header-name-2": jsii.String("custom-header-value-2"),
			},
		},
		&customResponseHeader{
			pattern: jsii.String("/path/*"),
			headers: map[string]*string{
				"custom-header-name-1": jsii.String("custom-header-value-2"),
			},
		},
	},
})
```

## Deploying Assets

`sourceCodeProvider` is optional; when this is not specified the Amplify app can be deployed to using `.zip` packages. The `asset` property can be used to deploy S3 assets to Amplify as part of the CDK:

```go
import assets "github.com/aws/aws-cdk-go/awscdk"

var asset asset
var amplifyApp app

branch := amplifyApp.addBranch(jsii.String("dev"), &branchOptions{
	asset: asset,
})
```
