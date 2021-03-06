package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/drone-plugins/drone-cache/cache"
	"github.com/drone-plugins/drone-cache/storage"
)

type Plugin struct {
	Filename string
	Path     string
	Mode     string
	Mount    []string

	Storage storage.Storage
}

const (
	RestoreMode = "restore"
	RebuildMode = "rebuild"
)

// Exec runs the plugin
func (p *Plugin) Exec() error {
	c, err := cache.New(p.Storage)

	if err != nil {
		return err
	}

	path := p.Path + p.Filename

	if p.Mode == RebuildMode {
		log.Infof("Rebuilding cache at %s", path)
		err = c.Rebuild(p.Mount, path)

		if err == nil {
			log.Infof("Cache rebuilt")
		}

		return err
	}

	log.Infof("Restoring cache at %s", path)
	err = c.Restore(path)

	if err == nil {
		log.Info("Cache restored")
	}

	return err
}
