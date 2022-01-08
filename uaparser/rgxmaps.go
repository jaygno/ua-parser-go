package uaparser

import "regexp"

var rgxMaps map[string][]*BrowserParser = map[string][]*BrowserParser{
	"browser": []*BrowserParser{
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(?:crmo|crios)\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       CHROME,
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)edg(?:e|ios|a)?\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       EDGE,
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(opera mini)\/([-\w\.]+)`),
				regexp.MustCompile(`(?i)(opera [mobiletab]{3,6})\b.+version\/([-\w\.]+)`),
				regexp.MustCompile(`(?i)(opera)(?:.+version\/|[\/ ]+)([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)opios[\/ ]+([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       OPERA + " Mini",
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bopr\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       OPERA,
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(kindle)\/([\w\.]+)`),
				regexp.MustCompile(`(?i)(lunascape|maxthon|netfront|jasmine|blazer)[\/ ]?([\w\.]*)`),
				regexp.MustCompile(`(?i)(avant |iemobile|slim)(?:browser)?[\/ ]?([\w\.]*)`),
				regexp.MustCompile(`(?i)(ba?idubrowser)[\/ ]?([\w\.]+)`),
				regexp.MustCompile(`(?i)(?:ms|\()(ie) ([\w\.]+)`),
				regexp.MustCompile(`(?i)(flock|rockmelt|midori|epiphany|silk|skyfire|ovibrowser|bolt|iron|vivaldi|iridium|phantomjs|bowser|quark|qupzilla|falkon|rekonq|puffin|brave|whale|qqbrowserlite|qq)\/([-\w\.]+)`),
				regexp.MustCompile(`(?i)(weibo)__([\d\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:\buc? ?browser|(?:juc.+)ucweb)[\/ ]?([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "UC" + BROWSER,
		},

		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bqbcore\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "WeChat(Win) Desktop",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)micromessenger\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "WeChat",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)konqueror\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "Konqueror",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)trident.+rv[: ]([\w\.]{1,9})\b.+like gecko`),
			},
			VersionIdx: 1,
			Name:       IE,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)yabrowser\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "Yandex",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(avast|avg)\/([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1, // TODO CHECK HERE
			NameParser: func(b []byte) string {
				return string(b) + " Secure " + BROWSER
			},
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bfocus\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       FIREFOX + " Focus",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bopt\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       OPERA + " Touch",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)coc_coc\w+\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "Coc Coc",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)dolfin\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "Dolphin",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)coast\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       OPERA + " Coast",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)miuibrowser\/([\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       "MIUI " + BROWSER,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)fxios\/([-\w\.]+)`),
			},
			VersionIdx: 1,
			Name:       FIREFOX,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bqihu|(qi?ho?o?|360)browser`),
			},
			VersionIdx: 10,
			Name:       "360 " + BROWSER,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(oculus|samsung|sailfish)browser\/([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
			NameParser: func(b []byte) string {
				return string(b) + " " + BROWSER
			},
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(comodo_dragon)\/([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(electron)\/([\w\.]+) safari`),
				regexp.MustCompile(`(?i)(tesla)(?: qtcarbrowser|\/(20\d\d\.[-\w\.]+))`),
				regexp.MustCompile(`(?i)m?(qqbrowser|baiduboxapp|2345Explorer)[\/ ]?([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(metasr)[\/ ]?([\w\.]+)`),
				regexp.MustCompile(`(?i)(lbbrowser)`),
			},
			VersionIdx: 10, // TODO fix here no version exist
			NameIndex:  1,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				//regexp.MustCompile(`(?i)((?:fban\/fbios|fb_iab\/fb4a)(?!.+fbav)|;fbav\/([\w\.]+);)`),
				regexp.MustCompile(`(?i)(fb_iab\/fb4a;fbav\/([\w\.]+);)`),
				regexp.MustCompile(`(?i)(fban\/fbios;fbav\/([\w\.]+);)`),
				regexp.MustCompile(`(?i)(fban\/fbios;)`),
			},
			VersionIdx: 2,
			Name:       FACEBOOK,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)safari (line)\/([\w\.]+)`),
				regexp.MustCompile(`(?i)\b(line)\/([\w\.]+)\/iab`),
				regexp.MustCompile(`(?i)(chromium|instagram)[\/ ]([-\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bgsa\/([\w\.]+) .*safari\/`),
			},
			VersionIdx: 1,
			Name:       "GSA",
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)headlesschrome(?:\/([\w\.]+)| )`),
			},
			VersionIdx: 1,
			Name:       CHROME + " Headless",
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i) wv\).+(chrome)\/([\w\.]+)`),
			},
			VersionIdx: 2,
			Name:       CHROME + " WebView",
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+ version\/([\w\.]+)\b.+(?:mobile safari|safari)`),
			},
			VersionIdx: 1,
			Name:       "Android " + BROWSER,
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(chrome|omniweb|arora|[tizenoka]{5} ?browser)\/v?([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)version\/([\w\.]+) .*mobile\/\w+ (safari)`),
			},
			VersionIdx: 1,
			Name:       "Mobile Safari",
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)version\/([\w\.]+) .*(mobile ?safari|safari)`),
			},
			VersionIdx: 1,
			NameIndex:  2,
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)webkit.+?(mobile ?safari|safari)(\/[\w\.]+)`),
			},
			VersionIdx:    2, // TODO FIX get version
			VersionParser: SafariMap,
			NameIndex:     1,
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(webkit|khtml)\/([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(navigator|netscape\d?)\/([-\w\.]+)`),
			},
			VersionIdx: 2,
			Name:       "Netscape",
		}, &BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)mobile vr; rv:([\w\.]+)\).+firefox`),
			},
			VersionIdx: 1,
			Name:       FIREFOX + " Reality",
		},
		&BrowserParser{
			Expr: []*regexp.Regexp{
				regexp.MustCompile(`(?i)ekiohf.+(flow)\/([\w\.]+)`),
				regexp.MustCompile(`(?i)(swiftfox)`),

				regexp.MustCompile(`(?i)(icedragon|iceweasel|camino|chimera|fennec|maemo browser|minimo|conkeror|klar)[\/ ]?([\w\.\+]+)`),
				regexp.MustCompile(`(?i)(seamonkey|k-meleon|icecat|iceape|firebird|phoenix|palemoon|basilisk|waterfox)\/([-\w\.]+)$`),
				regexp.MustCompile(`(?i)(firefox)\/([\w\.]+)`),
				regexp.MustCompile(`(?i)(mozilla)\/([\w\.]+) .+rv\:.+gecko\/\d+`),
				regexp.MustCompile(`(?i)(polaris|lynx|dillo|icab|doris|amaya|w3m|netsurf|sleipnir|obigo|mosaic|(?:go|ice|up)[\. ]?browser)[-\/ ]?v?([\w\.]+)`),
				regexp.MustCompile(`(?i)(links) \(([\w\.]+)`),
			},
			VersionIdx: 2,
			NameIndex:  1,
		},
	},
}
