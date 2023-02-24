# Amazon Lambda Python Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This library provides constructs for Python Lambda functions.

To use this module, you will need to have Docker installed.

## Python Function

Define a `PythonFunction`:

```go
python.NewPythonFunction(this, jsii.String("MyFunction"), &PythonFunctionProps{
	Entry: jsii.String("/path/to/my/function"),
	 // required
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	 // required
	Index: jsii.String("my_index.py"),
	 // optional, defaults to 'index.py'
	Handler: jsii.String("my_exported_func"),
})
```

All other properties of `lambda.Function` are supported, see also the [AWS Lambda construct library](https://github.com/aws/aws-cdk/tree/main/packages/%40aws-cdk/aws-lambda).

## Python Layer

You may create a python-based lambda layer with `PythonLayerVersion`. If `PythonLayerVersion` detects a `requirements.txt`
or `Pipfile` or `poetry.lock` with the associated `pyproject.toml` at the entry path, then `PythonLayerVersion` will include the dependencies inline with your code in the
layer.

Define a `PythonLayerVersion`:

```go
python.NewPythonLayerVersion(this, jsii.String("MyLayer"), &PythonLayerVersionProps{
	Entry: jsii.String("/path/to/my/layer"),
})
```

A layer can also be used as a part of a `PythonFunction`:

```go
python.NewPythonFunction(this, jsii.String("MyFunction"), &PythonFunctionProps{
	Entry: jsii.String("/path/to/my/function"),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Layers: []iLayerVersion{
		python.NewPythonLayerVersion(this, jsii.String("MyLayer"), &PythonLayerVersionProps{
			Entry: jsii.String("/path/to/my/layer"),
		}),
	},
})
```

## Packaging

If `requirements.txt`, `Pipfile` or `poetry.lock` exists at the entry path, the construct will handle installing all required modules in a [Lambda compatible Docker container](https://gallery.ecr.aws/sam/build-python3.7) according to the `runtime` and with the Docker platform based on the target architecture of the Lambda function.

Python bundles are only recreated and published when a file in a source directory has changed.
Therefore (and as a general best-practice), it is highly recommended to commit a lockfile with a
list of all transitive dependencies and their exact versions. This will ensure that when any dependency version is updated, the bundle asset is recreated and uploaded.

To that end, we recommend using [`pipenv`] or [`poetry`] which have lockfile support.

* [`pipenv`](https://pipenv-fork.readthedocs.io/en/latest/basics.html#example-pipfile-lock)
* [`poetry`](https://python-poetry.org/docs/basic-usage/#commit-your-poetrylock-file-to-version-control)

Packaging is executed using the `Packaging` class, which:

1. Infers the packaging type based on the files present.
2. If it sees a `Pipfile` or a `poetry.lock` file, it exports it to a compatible `requirements.txt` file with credentials (if they're available in the source files or in the bundling container).
3. Installs dependencies using `pip`.
4. Copies the dependencies into an asset that is bundled for the Lambda package.

**Lambda with a requirements.txt**

```plaintext
.
├── lambda_function.py # exports a function named 'handler'
├── requirements.txt # has to be present at the entry path
```

**Lambda with a Pipfile**

```plaintext
.
├── lambda_function.py # exports a function named 'handler'
├── Pipfile # has to be present at the entry path
├── Pipfile.lock # your lock file
```

**Lambda with a poetry.lock**

```plaintext
.
├── lambda_function.py # exports a function named 'handler'
├── pyproject.toml # your poetry project definition
├── poetry.lock # your poetry lock file has to be present at the entry path
```

**Excluding source files**

You can exclude files from being copied using the optional bundling string array parameter `assetExcludes`

```go
python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String("/path/to/poetry-function"),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		// translates to `rsync --exclude='.venv'`
		AssetExcludes: []*string{
			jsii.String(".venv"),
		},
	},
})
```

## Custom Bundling

Custom bundling can be performed by passing in additional build arguments that point to index URLs to private repos, or by using an entirely custom Docker images for bundling dependencies. The build args currently supported are:

* `PIP_INDEX_URL`
* `PIP_EXTRA_INDEX_URL`
* `HTTPS_PROXY`

Additional build args for bundling that refer to PyPI indexes can be specified as:

```go
entry := "/path/to/function"
image := awscdk.DockerImage_FromBuild(entry)

python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		BuildArgs: map[string]*string{
			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
		},
	},
})
```

If using a custom Docker image for bundling, the dependencies are installed with `pip`, `pipenv` or `poetry` by using the `Packaging` class. A different bundling Docker image that is in the same directory as the function can be specified as:

```go
entry := "/path/to/function"
image := awscdk.DockerImage_FromBuild(entry)

python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		Image: *Image,
	},
})
```

You can set additional Docker options to configure the build environment:

```go
entry := "/path/to/function"

python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		Network: jsii.String("host"),
		SecurityOpt: jsii.String("no-new-privileges"),
		User: jsii.String("user:group"),
		VolumesFrom: []*string{
			jsii.String("777f7dc92da7"),
		},
		Volumes: []dockerVolume{
			&dockerVolume{
				HostPath: jsii.String("/host-path"),
				ContainerPath: jsii.String("/container-path"),
			},
		},
	},
})
```

## Custom Bundling with Code Artifact

To use a Code Artifact PyPI repo, the `PIP_INDEX_URL` for bundling the function can be customized (requires AWS CLI in the build environment):

```go
import "github.com/aws-samples/dummy/child_process"


entry := "/path/to/function"
image := awscdk.DockerImage_FromBuild(entry)

domain := "my-domain"
domainOwner := "111122223333"
repoName := "my_repo"
region := "us-east-1"
codeArtifactAuthToken := child_process.ExecSync(fmt.Sprintf("aws codeartifact get-authorization-token --domain %v --domain-owner %v --query authorizationToken --output text", domain, domainOwner)).toString().trim()

indexUrl := fmt.Sprintf("https://aws:%v@%v-%v.d.codeartifact.%v.amazonaws.com/pypi/%v/simple/", codeArtifactAuthToken, domain, domainOwner, region, repoName)

python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		Environment: map[string]*string{
			"PIP_INDEX_URL": indexUrl,
		},
	},
})
```

The index URL or the token are only used during bundling and thus not included in the final asset. Setting only environment variable for `PIP_INDEX_URL` or `PIP_EXTRA_INDEX_URL` should work for accesing private Python repositories with `pip`, `pipenv` and `poetry` based dependencies.

If you also want to use the Code Artifact repo for building the base Docker image for bundling, use `buildArgs`. However, note that setting custom build args for bundling will force the base bundling image to be rebuilt every time (i.e. skip the Docker cache). Build args can be customized as:

```go
import "github.com/aws-samples/dummy/child_process"


entry := "/path/to/function"
image := awscdk.DockerImage_FromBuild(entry)

domain := "my-domain"
domainOwner := "111122223333"
repoName := "my_repo"
region := "us-east-1"
codeArtifactAuthToken := child_process.ExecSync(fmt.Sprintf("aws codeartifact get-authorization-token --domain %v --domain-owner %v --query authorizationToken --output text", domain, domainOwner)).toString().trim()

indexUrl := fmt.Sprintf("https://aws:%v@%v-%v.d.codeartifact.%v.amazonaws.com/pypi/%v/simple/", codeArtifactAuthToken, domain, domainOwner, region, repoName)

python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		BuildArgs: map[string]*string{
			"PIP_INDEX_URL": indexUrl,
		},
	},
})
```

## Command hooks

It is  possible to run additional commands by specifying the `commandHooks` prop:

```go
entry := "/path/to/function"
python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		CommandHooks: map[string]interface{}{
			// run tests
			(MethodDeclaration beforeBundling(inputDir: string): string[] {
			        return ['pytest'];
			      }
					beforeBundling
					inputDir *string
					string[]
					{
						return []*string{
							jsii.String("pytest"),
						}
					}),
			(MethodDeclaration afterBundling(inputDir: string): string[] {
			        return ['pylint'];
			      }
					afterBundling
					inputDir *string
					string[]
					{
						return []*string{
							jsii.String("pylint"),
						}
					}),
		},
	},
})
```

The following hooks are available:

* `beforeBundling`: runs before all bundling commands
* `afterBundling`: runs after all bundling commands

They all receive the directory containing the dependencies file (`inputDir`) and the
directory where the bundled asset will be output (`outputDir`). They must return
an array of commands to run. Commands are chained with `&&`.

The commands will run in the environment in which bundling occurs: inside the
container for Docker bundling or on the host OS for local bundling.

## Docker based bundling in complex Docker configurations

By default the input and output of Docker based bundling is handled via bind mounts.
In situtations where this does not work, like Docker-in-Docker setups or when using a remote Docker socket, you can configure an alternative, but slower, variant that also works in these situations.

```go
entry := "/path/to/function"

python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
	Entry: jsii.String(Entry),
	Runtime: awscdk.Runtime_PYTHON_3_8(),
	Bundling: &BundlingOptions{
		BundlingFileAccess: awscdk.BundlingFileAccess_VOLUME_COPY,
	},
})
```
