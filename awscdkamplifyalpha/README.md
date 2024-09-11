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


amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
		Owner: jsii.String("<user>"),
		Repository: jsii.String("<repo>"),
		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	}),
	BuildSpec: codebuild.BuildSpec_FromObjectToYaml(map[string]interface{}{
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
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	SourceCodeProvider: amplify.NewGitLabSourceCodeProvider(&GitLabSourceCodeProviderProps{
		Owner: jsii.String("<user>"),
		Repository: jsii.String("<repo>"),
		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-gitlab-token")),
	}),
})
```

To connect your `App` to CodeCommit, use the `CodeCommitSourceCodeProvider`:

```go
import codecommit "github.com/aws/aws-cdk-go/awscdk"


repository := codecommit.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
	RepositoryName: jsii.String("my-repo"),
})

amplifyApp := amplify.NewApp(this, jsii.String("App"), &AppProps{
	SourceCodeProvider: amplify.NewCodeCommitSourceCodeProvider(&CodeCommitSourceCodeProviderProps{
		Repository: *Repository,
	}),
})
```

The IAM role associated with the `App` will automatically be granted the permission
to pull the CodeCommit repository.

Add branches:

```go
var amplifyApp app


main := amplifyApp.AddBranch(jsii.String("main")) // `id` will be used as repo branch name
dev := amplifyApp.AddBranch(jsii.String("dev"), &BranchOptions{
	PerformanceMode: jsii.Boolean(true),
})
dev.AddEnvironment(jsii.String("STAGE"), jsii.String("dev"))
```

Auto build and pull request preview are enabled by default.

Add custom rules for redirection:

```go
var amplifyApp app

amplifyApp.AddCustomRule(map[string]interface{}{
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


mySinglePageApp.AddCustomRule(amplify.CustomRule_SINGLE_PAGE_APPLICATION_REDIRECT())
```

Add a domain and map sub domains to branches:

```go
var amplifyApp app
var main branch
var dev branch


domain := amplifyApp.AddDomain(jsii.String("example.com"), &DomainOptions{
	EnableAutoSubdomain: jsii.Boolean(true),
	 // in case subdomains should be auto registered for branches
	AutoSubdomainCreationPatterns: []*string{
		jsii.String("*"),
		jsii.String("pr*"),
	},
})
domain.MapRoot(main) // map main branch to domain root
domain.MapSubDomain(main, jsii.String("www"))
domain.MapSubDomain(dev)
```

To specify a custom certificate for your custom domain use the `customCertificate` property:

```go
var customCertificate certificate
var amplifyApp app


domain := amplifyApp.AddDomain(jsii.String("example.com"), &DomainOptions{
	CustomCertificate: CustomCertificate,
})
```

## Restricting access

Password protect the app with basic auth by specifying the `basicAuth` prop.

Use `BasicAuth.fromCredentials` when referencing an existing secret:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
		Owner: jsii.String("<user>"),
		Repository: jsii.String("<repo>"),
		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	}),
	BasicAuth: amplify.BasicAuth_FromCredentials(jsii.String("username"), awscdk.SecretValue_*SecretsManager(jsii.String("my-github-token"))),
})
```

Use `BasicAuth.fromGeneratedPassword` to generate a password in Secrets Manager:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
		Owner: jsii.String("<user>"),
		Repository: jsii.String("<repo>"),
		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	}),
	BasicAuth: amplify.BasicAuth_FromGeneratedPassword(jsii.String("username")),
})
```

Basic auth can be added to specific branches:

```go
var amplifyApp app

amplifyApp.AddBranch(jsii.String("feature/next"), &BranchOptions{
	BasicAuth: amplify.BasicAuth_FromGeneratedPassword(jsii.String("username")),
})
```

## Automatically creating and deleting branches

Use the `autoBranchCreation` and `autoBranchDeletion` props to control creation/deletion
of branches:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
		Owner: jsii.String("<user>"),
		Repository: jsii.String("<repo>"),
		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	}),
	AutoBranchCreation: &AutoBranchCreation{
		 // Automatically connect branches that match a pattern set
		Patterns: []*string{
			jsii.String("feature/*"),
			jsii.String("test/*"),
		},
	},
	AutoBranchDeletion: jsii.Boolean(true),
})
```

## Adding custom response headers

Use the `customResponseHeaders` prop to configure custom response headers for an Amplify app:

```go
amplifyApp := amplify.NewApp(this, jsii.String("App"), &AppProps{
	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
		Owner: jsii.String("<user>"),
		Repository: jsii.String("<repo>"),
		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
	}),
	CustomResponseHeaders: []customResponseHeader{
		&customResponseHeader{
			Pattern: jsii.String("*.json"),
			Headers: map[string]*string{
				"custom-header-name-1": jsii.String("custom-header-value-1"),
				"custom-header-name-2": jsii.String("custom-header-value-2"),
			},
		},
		&customResponseHeader{
			Pattern: jsii.String("/path/*"),
			Headers: map[string]*string{
				"custom-header-name-1": jsii.String("custom-header-value-2"),
			},
		},
	},
})
```

## Configure server side rendering when hosting app

Setting the `platform` field on the Amplify `App` construct can be used to control whether the app will host only static assets or server side rendered assets in addition to static. By default, the value is set to `WEB` (static only), however, server side rendering can be turned on by setting to `WEB_COMPUTE` as follows:

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	Platform: amplify.Platform_WEB_COMPUTE,
})
```

## Cache Config

Amplify uses Amazon CloudFront to manage the caching configuration for your hosted applications. A cache configuration is applied to each app to optimize for the best performance.

Setting the `cacheConfigType` field on the Amplify `App` construct can be used to control cache configguration. By default, the value is set to `AMPLIFY_MANAGED`. If you want to exclude all cookies from the cache key, set `AMPLIFY_MANAGED_NO_COOKIES`.

For more information, see [Managing the cache configuration for an app](https://docs.aws.amazon.com/amplify/latest/userguide/caching.html).

```go
amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
	CacheConfigType: amplify.CacheConfigType_AMPLIFY_MANAGED_NO_COOKIES,
})
```

## Deploying Assets

`sourceCodeProvider` is optional; when this is not specified the Amplify app can be deployed to using `.zip` packages. The `asset` property can be used to deploy S3 assets to Amplify as part of the CDK:

```go
import assets "github.com/aws/aws-cdk-go/awscdk"

var asset asset
var amplifyApp app

branch := amplifyApp.AddBranch(jsii.String("dev"), &BranchOptions{
	Asset: asset,
})
```
