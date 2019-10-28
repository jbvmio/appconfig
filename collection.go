package appconfig

import (
	"regexp"
)

// Collection represents a collection of Data.
type Collection []Data

// AssignADs assigns the AppDomain for all underlying data if applicable, otherwise will assign the default value given.
// If the value given is comma delimited, the left most value is will be used.
// If an array is given, then the first element is used.
func (c Collection) AssignADs(defaultValue ...string) {
	var dv string
	if len(defaultValue) > 0 {
		dv = defaultValue[0]
	}
	for i := 0; i < len(c); i++ {
		c[i].AssignAD(adSplit(dv))
	}
}

// Keys returns all the keys found in the Collection.
func (c Collection) Keys() (keys []string) {
	dupe := make(map[string]bool)
	for _, d := range c {
		if !dupe[d.Key] {
			dupe[d.Key] = true
			keys = append(keys, d.Key)
		}
	}
	return
}

// Values returns all the values found in the Collection.
func (c Collection) Values() (values []string) {
	dupe := make(map[string]bool)
	for _, d := range c {
		if !dupe[d.Value] {
			dupe[d.Value] = true
			values = append(values, d.Value)
		}
	}
	return
}

// Pkgs returns all the pkgs found in the Collection.
func (c Collection) Pkgs() (pkgs []string) {
	dupe := make(map[string]bool)
	for _, d := range c {
		if !dupe[d.Pkg] {
			dupe[d.Pkg] = true
			pkgs = append(pkgs, d.Pkg)
		}
	}
	return
}

// ADs returns all the AppDomains found in the Collection.
func (c Collection) ADs() (ads []string) {
	dupe := make(map[string]bool)
	for _, d := range c {
		if !dupe[d.AppDomain] {
			dupe[d.AppDomain] = true
			ads = append(ads, d.AppDomain)
		}
	}
	return
}

// Types returns all the types found in the Collection.
func (c Collection) Types() (types []string) {
	dupe := make(map[string]bool)
	for _, d := range c {
		if !dupe[d.T] {
			dupe[d.T] = true
			types = append(types, d.T)
		}
	}
	return
}

// HasType returns true if the given DataType is present in the Collection.
func (c Collection) HasType(dt DataType) bool {
	for _, d := range c {
		if d.DataType() == dt {
			return true
		}
	}
	return false
}

// FromType returns a sub Collection containing only the specified DataType.
func (c Collection) FromType(dt DataType) Collection {
	var data []Data
	for _, d := range c {
		if d.DataType() == dt {
			data = append(data, d)
		}
	}
	return data
}

// FromPkg returns a sub Collection containing only the data that contains a reference to the given pkg.
func (c Collection) FromPkg(pkg string) Collection {
	var data []Data
	for _, d := range c {
		if d.HasPkg(pkg) {
			data = append(data, d)
		}
	}
	return data
}

// FromPkgRegexp returns a sub Collection containing only the data whose pkg matches the given regexp.
func (c Collection) FromPkgRegexp(regex *regexp.Regexp) Collection {
	var data []Data
	for _, d := range c {
		if regex.MatchString(d.Pkg) {
			data = append(data, d)
		}
	}
	return data
}

// FromKey returns a sub Collection containing only the data that contains a reference to the given key.
func (c Collection) FromKey(key string) Collection {
	var data []Data
	for _, d := range c {
		if d.HasKey(key) {
			data = append(data, d)
		}
	}
	return data
}

// FromKeyRegexp returns a sub Collection containing only the data whose key matches the given regexp.
func (c Collection) FromKeyRegexp(regex *regexp.Regexp) Collection {
	var data []Data
	for _, d := range c {
		if regex.MatchString(d.Key) {
			data = append(data, d)
		}
	}
	return data
}

// FromAD returns a sub Collection containing only the data that contains a reference to the given AppDomain.
func (c Collection) FromAD(ad string) Collection {
	var data []Data
	for _, d := range c {
		if d.HasAD(ad) {
			data = append(data, d)
		}
	}
	return data
}

// FromADRegexp returns a sub Collection containing only the data if the AppDomain matches the given regexp.
func (c Collection) FromADRegexp(regex *regexp.Regexp) Collection {
	var data []Data
	for _, d := range c {
		if regex.MatchString(d.AppDomain) {
			data = append(data, d)
		}
	}
	return data
}

// Get returns all the values found in the Collection if the given key matches.
func (c Collection) Get(key string) (values []string) {
	for _, d := range c {
		if d.HasKey(key) {
			values = append(values, d.Value)
		}
	}
	return
}

// GetFromPkg returns all the values found in the Collection for the given pkg and key.
func (c Collection) GetFromPkg(pkg, key string) (values []string) {
	for _, d := range c {
		if d.HasPkg(pkg) && d.HasKey(key) {
			values = append(values, d.Value)
		}
	}
	return
}
