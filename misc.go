package appconfig

import "strings"

func epAD(ep string) string {
	var val string
	tmp := strings.Split(ep, `'`)
	for _, t := range tmp {
		x := strings.Split(t, `:`)
		if len(x) > 1 {
			val = x[len(x)-1]
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
