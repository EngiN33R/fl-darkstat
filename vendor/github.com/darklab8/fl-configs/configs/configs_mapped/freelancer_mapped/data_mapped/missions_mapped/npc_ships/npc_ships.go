package npc_ships

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
)

const (
	FILENAME = "npcships.ini"
)

type NPCShipArch struct {
	semantic.Model
	Nickname *semantic.String
	Level    *semantic.String
	NpcClass []*semantic.String
}

type Config struct {
	*iniload.IniLoader

	NpcShips           []*NPCShipArch
	NpcShipsByNickname map[string]*NPCShipArch
}

func Read(input_file *iniload.IniLoader) *Config {
	frelconfig := &Config{
		IniLoader:          input_file,
		NpcShipsByNickname: make(map[string]*NPCShipArch),
	}
	if sections, ok := frelconfig.SectionMap["[NPCShipArch]"]; ok {
		for _, section := range sections {
			npc_ship_arch := &NPCShipArch{}
			npc_ship_arch.Map(section)
			npc_ship_arch.Nickname = semantic.NewString(section, "nickname", semantic.WithLowercaseS(), semantic.WithoutSpacesS())
			npc_ship_arch.Level = semantic.NewString(section, "level", semantic.WithLowercaseS(), semantic.WithoutSpacesS())

			if npc_class_param, ok := section.ParamMap["npc_class"]; ok {
				for index_order, _ := range npc_class_param[0].Values {
					npc_ship_arch.NpcClass = append(npc_ship_arch.NpcClass,
						semantic.NewString(section, "npc_class", semantic.OptsS(semantic.Order(index_order)), semantic.WithLowercaseS(), semantic.WithoutSpacesS()))
				}

			}
			frelconfig.NpcShips = append(frelconfig.NpcShips, npc_ship_arch)
			frelconfig.NpcShipsByNickname[npc_ship_arch.Nickname.Get()] = npc_ship_arch
		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}
