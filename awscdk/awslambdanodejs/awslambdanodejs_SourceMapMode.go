package awslambdanodejs


// SourceMap mode for esbuild.
//
// Example:
//   nodejs.NewNodejsFunction(this, jsii.String("my-handler"), &nodejsFunctionProps{
//   	bundling: &bundlingOptions{
//   		minify: jsii.Boolean(true),
//   		 // minify code, defaults to false
//   		sourceMap: jsii.Boolean(true),
//   		 // include source map, defaults to false
//   		sourceMapMode: nodejs.sourceMapMode_INLINE,
//   		 // defaults to SourceMapMode.DEFAULT
//   		sourcesContent: jsii.Boolean(false),
//   		 // do not include original source into source map, defaults to true
//   		target: jsii.String("es2020"),
//   		 // target environment for the generated JavaScript code
//   		loader: map[string]*string{
//   			 // Use the 'dataurl' loader for '.png' files
//   			".png": jsii.String("dataurl"),
//   		},
//   		define: map[string]*string{
//   			 // Replace strings during build time
//   			"process.env.API_KEY": JSON.stringify(jsii.String("xxx-xxxx-xxx")),
//   			"process.env.PRODUCTION": JSON.stringify(jsii.Boolean(true)),
//   			"process.env.NUMBER": JSON.stringify(jsii.Number(123)),
//   		},
//   		logLevel: nodejs.logLevel_SILENT,
//   		 // defaults to LogLevel.WARNING
//   		keepNames: jsii.Boolean(true),
//   		 // defaults to false
//   		tsconfig: jsii.String("custom-tsconfig.json"),
//   		 // use custom-tsconfig.json instead of default,
//   		metafile: jsii.Boolean(true),
//   		 // include meta file, defaults to false
//   		banner: jsii.String("/* comments */"),
//   		 // requires esbuild >= 0.9.0, defaults to none
//   		footer: jsii.String("/* comments */"),
//   		 // requires esbuild >= 0.9.0, defaults to none
//   		charset: nodejs.charset_UTF8,
//   		 // do not escape non-ASCII characters, defaults to Charset.ASCII
//   		format: nodejs.outputFormat_ESM,
//   		 // ECMAScript module output format, defaults to OutputFormat.CJS (OutputFormat.ESM requires Node.js 14.x)
//   		mainFields: []*string{
//   			jsii.String("module"),
//   			jsii.String("main"),
//   		},
//   		 // prefer ECMAScript versions of dependencies
//   		inject: []*string{
//   			jsii.String("./my-shim.js"),
//   			jsii.String("./other-shim.js"),
//   		},
//   		 // allows to automatically replace a global variable with an import from another file
//   		esbuildArgs: map[string]interface{}{
//   			 // Pass additional arguments to esbuild
//   			"--log-limit": jsii.String("0"),
//   			"--splitting": jsii.Boolean(true),
//   		},
//   	},
//   })
//
// See: https://esbuild.github.io/api/#sourcemap
//
type SourceMapMode string

const (
	// Default sourceMap mode - will generate a .js.map file alongside any generated .js file and add a special //# sourceMappingURL= comment to the bottom of the .js file pointing to the .js.map file.
	SourceMapMode_DEFAULT SourceMapMode = "DEFAULT"
	// External sourceMap mode - If you want to omit the special //# sourceMappingURL= comment from the generated .js file but you still want to generate the .js.map files.
	SourceMapMode_EXTERNAL SourceMapMode = "EXTERNAL"
	// Inline sourceMap mode - If you want to insert the entire source map into the .js file instead of generating a separate .js.map file.
	SourceMapMode_INLINE SourceMapMode = "INLINE"
	// Both sourceMap mode - If you want to have the effect of both inline and external simultaneously.
	SourceMapMode_BOTH SourceMapMode = "BOTH"
)

