package version

// Buildstamp version number defined by the compiler:
//
//	-ldflags "-X github.com/flowline-io/flowkit/version.Buildstamp=value_to_assign_to_buildstamp"
//
// Reported to clients in response to {hi} message.
// For instance, to define the buildstamp as a timestamp of when the server was built add a
// flag to compiler command line:
//
//	-ldflags "-X github.com/flowline-io/flowkit/version.Buildstamp=`date -u '+%Y-%m-%dT%H:%M:%SZ'`"
var Buildstamp = "undef"

// Buildtags set it to git tag:
//
//	-ldflags "-X github.com/flowline-io/flowkit/version.Buildtags=`git describe --tags`"
var Buildtags = "v1.0"
