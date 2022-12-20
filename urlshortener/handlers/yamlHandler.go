package handlers

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func YamlHandler(b []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []PathsUrl
	err := yaml.Unmarshal(b, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type PathsUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
