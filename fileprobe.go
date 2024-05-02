package fileprobe

import (
	"os"
)

const _path = "/tmp/live"

type handler struct {
	filePath string
}

func NewHandler() *handler {
	return &handler{
		filePath: _path,
	}
}

func (h *handler) SetPath(path string) {
	h.filePath = path
}

func (h *handler) GetPath() string {
	return h.filePath
}

func (h *handler) Create() error {
	_, err := os.Create(h.filePath)
	return err
}

func (h *handler) Remove() error {
	return os.Remove(h.filePath)
}

func (h *handler) Exists() bool {
	info, err := os.Stat(h.filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
