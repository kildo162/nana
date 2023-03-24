package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func ParseVersion(s string) (v Version, err error) {
	parts := strings.Split(strings.TrimSpace(s), ".")
	if len(parts) != 3 {
		err = fmt.Errorf("can not parse string %s to version", s)
		return
	}
	v.Major, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}
	v.Minor, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}
	v.Patch, err = strconv.Atoi(parts[2])
	if err != nil {
		return
	}
	return
}

func (v *Version) NextPatch() {
	v.Patch = v.Patch + 1
}

func (v *Version) NextMinor() {
	v.Patch = 0
	v.Minor = v.Minor + 1
}

func (v *Version) NextMajor() {
	v.Patch = 0
	v.Minor = 0
	v.Major = v.Major + 1
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) MinorBranch() string {
	return fmt.Sprintf("%d.%d.x", v.Major, v.Minor)
}

func (v *Version) MajorBranch() string {
	return fmt.Sprintf("%d.x.x", v.Major)
}

func GetFileVersion() (c *Data, err error) {
	c, err = readConf("versions.yml")
	return c, err
}

func readConf(filename string) (*Data, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Data{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}

func UpdateVersion(moduleName string, version Version) error {
	data, err := GetFileVersion()
	if err != nil {
		fmt.Println("Cannot read versions.yml - " + err.Error())
		return err
	}

	for i, module := range data.Modules {
		if module.Name == moduleName {
			data.Modules[i].Version = version.String()
		}
	}

	err = writeConf("versions.yml", data)
	if err != nil {
		fmt.Println("Cannot write versions.yml - " + err.Error())
		return err
	}
	return nil
}

func writeConf(filename string, data *Data) error {
	buf, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("in file %q: %w", filename, err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.Write(buf)
	if err2 != nil {
		return err2
	}

	return nil
}
