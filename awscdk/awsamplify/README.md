# AWS Amplify Construct Library

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


master := amplifyApp.AddBranch(jsii.String("master")) // `id` will be used as repo branch name
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
var master branch
var dev branch


domain := amplifyApp.AddDomain(jsii.String("example.com"), &DomainOptions{
	EnableAutoSubdomain: jsii.Boolean(true),
	 // in case subdomains should be auto registered for branches
	AutoSubdomainCreationPatterns: []*string{
		jsii.String("*"),
		jsii.String("pr*"),
	},
})
domain.MapRoot(master) // map master branch to domain root
domain.MapSubDomain(master, jsii.String("www"))
domain.MapSubDomain(dev)
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
