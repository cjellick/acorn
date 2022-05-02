package v1

#SidecarSpec: {
	#ContainerBaseSpec
	init: bool | *false
}

#AcornBuildSpec: {
	context:   string | *"."
	acornfile: string | *"acorn.cue"
	params: [string]: _
}

#BuildSpec: {
	baseImage:  string | *""
	context:    string | *"."
	dockerfile: string | *"Dockerfile"
	target:     string | *""
	contextDirs: [string]: string
	args: [string]:        string
}

#AliasSpec: {
	name: string
}

#ContainerSpec: {
	#ContainerBaseSpec
	aliases: [...#AliasSpec]
	sidecars: [string]: #SidecarSpec
}

#JobSpec: {
	#ContainerBaseSpec
	schedule: string | *""
	sidecars: [string]: #SidecarSpec
}

#EnvSecretValue: {
	key:       string | *""
	name:      string
	optional?: bool
}

#EnvVarSpec: {
	name:    string
	value?:  string
	secret?: #EnvSecretValue
}

#ContainerBaseSpec: {
	image?: string
	build?: #BuildSpec
	entrypoint: [...string]
	command: [...string]
	environment: [...#EnvVarSpec]
	workingDir:  string | *""
	interactive: bool | *false
	ports: [...#PortSpec]
	files: [string]: #FileSpec
	dirs: [string]:  #VolumeMountSpec
}

#VolumeMountSpec: {
	{
		secret: {
			name:      string
			optional?: bool
		}
	} |
	{
		volume:  string
		subPath: string | *""
	} |
	{
		contextDir: string
	}
}

#FileSecretSpec: {
	name:      string
	key:       string
	optional?: bool
}

#FileSpec: {
	{
		content: string
	} | {
		secret: #FileSecretSpec
	}
}

#ImageSpec: {
	image:  string | *""
	build?: #BuildSpec
}

#AccessMode: "readWriteMany" | "readWriteOnce" | "readOnlyMany" | "readWriteOncePod"

#VolumeSpec: {
	class:       string | *""
	size:        int | *10
	accessModes: [#AccessMode, ...#AccessMode] | *["readWriteOnce"]
}

#SecretSpec: {
	type:      string
	optional?: bool
	params?: [string]: _
	data: [string]:    (string | bytes)
}

#VolumeBinding: {
	volume:        string
	volumeRequest: string
}

#SecretBinding: {
	secret:        string
	secretRequest: string
}

#AcornSpec: {
	image?: string
	build?: #AcornBuildSpec
	ports: [...#PortSpec]
	volumes: [...#VolumeBinding]
	secrets: [...#SecretBinding]
	params: [string]: _
}

#AppSpec: {
	containers: [string]: #ContainerSpec
	jobs: [string]:       #JobSpec
	images: [string]:     #ImageSpec
	volumes: [string]:    #VolumeSpec
	secrets: [string]:    #SecretSpec
	acorns: [string]:     #AcornSpec
}

#PortSpec: {
	publish:       bool | *false
	port:          int
	containerPort: int | *port
	protocol:      *"tcp" | "udp" | "http" | "https"
}
