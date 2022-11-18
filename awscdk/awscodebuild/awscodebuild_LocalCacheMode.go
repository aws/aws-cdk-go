package awscodebuild


// Local cache modes to enable for the CodeBuild Project.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	source: codebuild.source.gitHubEnterprise(&gitHubEnterpriseSourceProps{
//   		httpsCloneUrl: jsii.String("https://my-github-enterprise.com/owner/repo"),
//   	}),
//
//   	// Enable Docker AND custom caching
//   	cache: codebuild.cache.local(codebuild.localCacheMode_DOCKER_LAYER, codebuild.*localCacheMode_CUSTOM),
//
//   	// BuildSpec with a 'cache' section necessary for 'CUSTOM' caching. This can
//   	// also come from 'buildspec.yml' in your source.
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("..."),
//   				},
//   			},
//   		},
//   		"cache": map[string][]*string{
//   			"paths": []*string{
//   				jsii.String("/root/cachedir/**/*"),
//   			},
//   		},
//   	}),
//   })
//
type LocalCacheMode string

const (
	// Caches Git metadata for primary and secondary sources.
	LocalCacheMode_SOURCE LocalCacheMode = "SOURCE"
	// Caches existing Docker layers.
	LocalCacheMode_DOCKER_LAYER LocalCacheMode = "DOCKER_LAYER"
	// Caches directories you specify in the buildspec file.
	LocalCacheMode_CUSTOM LocalCacheMode = "CUSTOM"
)

