/*
Напишите метод клонирования объекта указанного типа.
*/

package main

import "fmt"

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	clone := &Config{}
	clone.Version = cfg.Version
	clone.Plugins = make([]string, len(cfg.Plugins))
	copy(clone.Plugins, cfg.Plugins)
	clone.Stat = make(map[string]int)
	for k, v := range cfg.Stat {
		clone.Stat[k] = v
	}
	return clone
}

func (cfg *Config) Clone0() *Config {
	clone := &Config{
		Version: cfg.Version,
		Stat:    make(map[string]int),
	}
	for _, v := range cfg.Plugins {
		clone.Plugins = append(clone.Plugins, v)
	}
	for key, v := range cfg.Stat {
		clone.Stat[key] = v
	}
	return clone
}

func main() {
	cfg := &Config{
		Version: "1.0.0",
		Plugins: []string{"plugin1", "plugin2"},
		Stat:    map[string]int{"stat1": 1, "stat2": 2},
	}
	clone := cfg.Clone()
	fmt.Println(clone)
}
