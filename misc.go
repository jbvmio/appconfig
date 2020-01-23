package appconfig

import "strings"

func epAD(ep string) string {
	var val string
	switch {
	case !strings.Contains(ep, `'`):
		return ""
	default:
		tmp := strings.Split(ep, `'`)
		for _, t := range tmp {
			x := strings.Split(t, `:`)
			if len(x) > 1 {
				val = x[len(x)-1]
			}
		}
	}
	return val
}

func adSplit(ad string) string {
	ads := strings.Split(ad, `,`)
	if len(ads) > 0 {
		return ads[0]
	}
	return ad
}

func filterUnique(vals []string) []string {
	var tmp []string
	dupe := make(map[string]bool)
	for _, v := range vals {
		if !dupe[v] {
			dupe[v] = true
			tmp = append(tmp, v)
		}
	}
	return tmp
}
