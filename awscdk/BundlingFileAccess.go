package awscdk


// The access mechanism used to make source files available to the bundling container and to return the bundling output back to the host.
//
// Example:
//   go.NewGoFunction(this, jsii.String("GoFunction"), &GoFunctionProps{
//   	Entry: jsii.String("app/cmd/api"),
//   	Bundling: &BundlingOptions{
//   		BundlingFileAccess: awscdk.BundlingFileAccess_VOLUME_COPY,
//   	},
//   })
//
type BundlingFileAccess string

const (
	// Creates temporary volumes and containers to copy files from the host to the bundling container and back.
	//
	// This is slower, but works also in more complex situations with remote or shared docker sockets.
	BundlingFileAccess_VOLUME_COPY BundlingFileAccess = "VOLUME_COPY"
	// The source and output folders will be mounted as bind mount from the host system This is faster and simpler, but less portable than `VOLUME_COPY`.
	BundlingFileAccess_BIND_MOUNT BundlingFileAccess = "BIND_MOUNT"
)

