package mounts // import "github.com/fdurand/moby/volume/mounts"

func (p *linuxParser) HasResource(m *MountPoint, absolutePath string) bool {
	return false
}
