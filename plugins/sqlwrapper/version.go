// Code generated by release workflow. DO NOT EDIT.

package sqlwrapper

// Version is the current release version of the otelsql instrumentation.
func Version() string {
	return "0.3.4"
}

// SemVersion is the semantic version to be supplied to tracer/meter creation.
func SemVersion() string {
	return "semver:" + Version()
}
