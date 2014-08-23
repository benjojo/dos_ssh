package main

import (
	"encoding/hex"
)

func VESAtoVT100(code byte) string {
	// Super hax lol
	template := "\x1B["
	out := hex.EncodeToString([]byte{code})
	if out[:1] == "0" {
		template = template + "40"
	} else if out[:1] == "1" {
		template = template + "44"
	} else if out[:1] == "2" {
		template = template + "42"
	} else if out[:1] == "3" {
		template = template + "46"
	} else if out[:1] == "4" {
		template = template + "41"
	} else if out[:1] == "5" {
		template = template + "45"
	} else if out[:1] == "6" {
		template = template + "43"
	} else if out[:1] == "7" {
		template = template + "47"
	} else if out[:1] == "8" {
		template = template + "40"
	} else if out[:1] == "9" {
		template = template + "44"
	} else if out[:1] == "A" {
		template = template + "42"
	} else if out[:1] == "B" {
		template = template + "46"
	} else if out[:1] == "C" {
		template = template + "41"
	} else if out[:1] == "D" {
		template = template + "45"
	} else if out[:1] == "E" {
		template = template + "43"
	}

	if out[1:] == "0" {
		template = template + ";30"
	} else if out[1:] == "1" {
		template = template + ";34"
	} else if out[1:] == "2" {
		template = template + ";32"
	} else if out[1:] == "3" {
		template = template + ";36"
	} else if out[1:] == "4" {
		template = template + ";31"
	} else if out[1:] == "5" {
		template = template + ";35"
	} else if out[1:] == "6" {
		template = template + ";33"
	} else if out[1:] == "7" {
		template = template + ";37"
	} else if out[1:] == "8" {
		template = template + ";30"
	} else if out[1:] == "9" {
		template = template + ";34"
	} else if out[1:] == "A" {
		template = template + ";32"
	} else if out[1:] == "B" {
		template = template + ";36"
	} else if out[1:] == "C" {
		template = template + ";31"
	} else if out[1:] == "D" {
		template = template + ";35"
	} else if out[1:] == "E" {
		template = template + ";33"
	}
	template = template + "m"
	return template
	/*
			Set Display Attributes

		Set Attribute Mode	<ESC>[{attr1};...;{attrn}m
		Sets multiple display attribute settings. The following lists standard attributes:
		0	Reset all attributes
		1	Bright
		2	Dim
		4	Underscore
		5	Blink
		7	Reverse
		8	Hidden

			Foreground Colours
		30	Black
		31	Red
		32	Green
		33	Yellow
		34	Blue
		35	Magenta
		36	Cyan
		37	White

			Background Colours
		40	Black
		41	Red
		42	Green
		43	Yellow
		44	Blue
		45	Magenta
		46	Cyan
		47	White
	*/
}

func CorrectBadChars(in byte) string {
	in2 := in - 1
	if uint8(in2) == 1 {
		return "☻"
	}
	if uint8(in2) == 2 {
		return "♥"
	}
	if uint8(in2) == 3 {
		return "♦"
	}
	if uint8(in2) == 4 {
		return "♣"
	}
	if uint8(in2) == 5 {
		return "♠"
	}
	if uint8(in2) == 6 {
		return "•"
	}
	if uint8(in2) == 7 {
		return "◘"
	}
	if uint8(in2) == 8 {
		return "○"
	}
	if uint8(in2) == 9 {
		return "◙"
	}
	if uint8(in2) == 10 {
		return "♂"
	}
	if uint8(in2) == 11 {
		return "♀"
	}
	if uint8(in2) == 12 {
		return "♪"
	}
	if uint8(in2) == 13 {
		return "♬"
	}
	if uint8(in2) == 14 {
		return "☼"
	}
	if uint8(in2) == 15 {
		return "►"
	}
	if uint8(in2) == 16 {
		return "◄"
	}
	if uint8(in2) == 17 {
		return "↕"
	}
	if uint8(in2) == 18 {
		return "‼"
	}
	if uint8(in2) == 19 {
		return "¶"
	}
	if uint8(in2) == 20 {
		return "§"
	}
	if uint8(in2) == 21 {
		return "▬"
	}
	if uint8(in2) == 22 {
		return "↨"
	}
	if uint8(in2) == 23 {
		return "↑"
	}
	if uint8(in2) == 24 {
		return "↓"
	}
	if uint8(in2) == 25 {
		return "→"
	}
	if uint8(in2) == 26 {
		return "←"
	}
	if uint8(in2) == 27 {
		return "∟"
	}
	if uint8(in2) == 28 {
		return "↔"
	}
	if uint8(in2) == 29 {
		return "▲"
	}
	if uint8(in2) == 30 {
		return "▼"
	}
	if uint8(in2) == 126 {
		return "⌂"
	}
	if uint8(in2) == 127 {
		return "Ç"
	}
	if uint8(in2) == 128 {
		return "ü"
	}
	if uint8(in2) == 129 {
		return "é"
	}
	if uint8(in2) == 130 {
		return "â"
	}
	if uint8(in2) == 131 {
		return "ä"
	}
	if uint8(in2) == 132 {
		return "à"
	}
	if uint8(in2) == 133 {
		return "å"
	}
	if uint8(in2) == 134 {
		return "ç"
	}
	if uint8(in2) == 135 {
		return "ê"
	}
	if uint8(in2) == 136 {
		return "ë"
	}
	if uint8(in2) == 137 {
		return "è"
	}
	if uint8(in2) == 138 {
		return "ï"
	}
	if uint8(in2) == 139 {
		return "î"
	}
	if uint8(in2) == 140 {
		return "ì"
	}
	if uint8(in2) == 141 {
		return "Ä"
	}
	if uint8(in2) == 142 {
		return "Å"
	}
	if uint8(in2) == 143 {
		return "É"
	}
	if uint8(in2) == 144 {
		return "æ"
	}
	if uint8(in2) == 145 {
		return "Æ"
	}
	if uint8(in2) == 146 {
		return "ô"
	}
	if uint8(in2) == 147 {
		return "ö"
	}
	if uint8(in2) == 148 {
		return "ò"
	}
	if uint8(in2) == 149 {
		return "û"
	}
	if uint8(in2) == 150 {
		return "ù"
	}
	if uint8(in2) == 151 {
		return "ÿ"
	}
	if uint8(in2) == 152 {
		return "Ö"
	}
	if uint8(in2) == 153 {
		return "Ü"
	}
	if uint8(in2) == 154 {
		return "¢"
	}
	if uint8(in2) == 155 {
		return "£"
	}
	if uint8(in2) == 156 {
		return "¥"
	}
	if uint8(in2) == 157 {
		return "₧"
	}
	if uint8(in2) == 158 {
		return "ƒ"
	}
	if uint8(in2) == 159 {
		return "á"
	}
	if uint8(in2) == 160 {
		return "í"
	}
	if uint8(in2) == 161 {
		return "ó"
	}
	if uint8(in2) == 162 {
		return "ú"
	}
	if uint8(in2) == 163 {
		return "ñ"
	}
	if uint8(in2) == 164 {
		return "Ñ"
	}
	if uint8(in2) == 165 {
		return "ª"
	}
	if uint8(in2) == 166 {
		return "º"
	}
	if uint8(in2) == 167 {
		return "¿"
	}
	if uint8(in2) == 168 {
		return "⌐"
	}
	if uint8(in2) == 169 {
		return "¬"
	}
	if uint8(in2) == 170 {
		return "½"
	}
	if uint8(in2) == 171 {
		return "¼"
	}
	if uint8(in2) == 172 {
		return "¡"
	}
	if uint8(in2) == 173 {
		return "«"
	}
	if uint8(in2) == 174 {
		return "»"
	}
	if uint8(in2) == 175 {
		return "░"
	}
	if uint8(in2) == 176 {
		return "▒"
	}
	if uint8(in2) == 177 {
		return "▓"
	}
	if uint8(in2) == 178 {
		return "│"
	}
	if uint8(in2) == 179 {
		return "┤"
	}
	if uint8(in2) == 180 {
		return "╡"
	}
	if uint8(in2) == 181 {
		return "╢"
	}
	if uint8(in2) == 182 {
		return "╖"
	}
	if uint8(in2) == 183 {
		return "╕"
	}
	if uint8(in2) == 184 {
		return "╣"
	}
	if uint8(in2) == 185 {
		return "║"
	}
	if uint8(in2) == 186 {
		return "╗"
	}
	if uint8(in2) == 187 {
		return "╝"
	}
	if uint8(in2) == 188 {
		return "╜"
	}
	if uint8(in2) == 189 {
		return "╛"
	}
	if uint8(in2) == 190 {
		return "┐"
	}
	if uint8(in2) == 191 {
		return "└"
	}
	if uint8(in2) == 192 {
		return "┴"
	}
	if uint8(in2) == 193 {
		return "┬"
	}
	if uint8(in2) == 194 {
		return "├"
	}
	if uint8(in2) == 195 {
		return "─"
	}
	if uint8(in2) == 196 {
		return "┼"
	}
	if uint8(in2) == 197 {
		return "╞"
	}
	if uint8(in2) == 198 {
		return "╟"
	}
	if uint8(in2) == 199 {
		return "╚"
	}
	if uint8(in2) == 200 {
		return "╔"
	}
	if uint8(in2) == 201 {
		return "╩"
	}
	if uint8(in2) == 202 {
		return "╦"
	}
	if uint8(in2) == 203 {
		return "╠"
	}
	if uint8(in2) == 204 {
		return "═"
	}
	if uint8(in2) == 205 {
		return "╬"
	}
	if uint8(in2) == 206 {
		return "╧"
	}
	if uint8(in2) == 207 {
		return "╨"
	}
	if uint8(in2) == 208 {
		return "╤"
	}
	if uint8(in2) == 209 {
		return "╥"
	}
	if uint8(in2) == 210 {
		return "╙"
	}
	if uint8(in2) == 211 {
		return "╘"
	}
	if uint8(in2) == 212 {
		return "╒"
	}
	if uint8(in2) == 213 {
		return "╓"
	}
	if uint8(in2) == 214 {
		return "╫"
	}
	if uint8(in2) == 215 {
		return "╪"
	}
	if uint8(in2) == 216 {
		return "┘"
	}
	if uint8(in2) == 217 {
		return "┌"
	}
	if uint8(in2) == 218 {
		return "█"
	}
	if uint8(in2) == 219 {
		return "▄"
	}
	if uint8(in2) == 220 {
		return "▌"
	}
	if uint8(in2) == 221 {
		return "▐"
	}
	if uint8(in2) == 222 {
		return "▀"
	}
	if uint8(in2) == 223 {
		return "α"
	}
	if uint8(in2) == 224 {
		return "ß"
	}
	if uint8(in2) == 225 {
		return "Γ"
	}
	if uint8(in2) == 226 {
		return "π"
	}
	if uint8(in2) == 227 {
		return "Σ"
	}
	if uint8(in2) == 228 {
		return "σ"
	}
	if uint8(in2) == 229 {
		return "µ"
	}
	if uint8(in2) == 230 {
		return "τ"
	}
	if uint8(in2) == 231 {
		return "Φ"
	}
	if uint8(in2) == 232 {
		return "Θ"
	}
	if uint8(in2) == 233 {
		return "Ω"
	}
	if uint8(in2) == 234 {
		return "δ"
	}
	if uint8(in2) == 235 {
		return "∞"
	}
	if uint8(in2) == 236 {
		return "φ"
	}
	if uint8(in2) == 237 {
		return "ε"
	}
	if uint8(in2) == 238 {
		return "∩"
	}
	if uint8(in2) == 239 {
		return "≡"
	}
	if uint8(in2) == 240 {
		return "±"
	}
	if uint8(in2) == 241 {
		return "≥"
	}
	if uint8(in2) == 242 {
		return "≤"
	}
	if uint8(in2) == 243 {
		return "⌠"
	}
	if uint8(in2) == 244 {
		return "⌡"
	}
	if uint8(in2) == 245 {
		return "÷"
	}
	if uint8(in2) == 246 {
		return "≈"
	}
	if uint8(in2) == 247 {
		return "°"
	}
	if uint8(in2) == 248 {
		return "∙"
	}
	if uint8(in2) == 249 {
		return "·"
	}
	if uint8(in2) == 250 {
		return "√"
	}
	if uint8(in2) == 251 {
		return "ⁿ"
	}
	if uint8(in2) == 252 {
		return "²"
	}
	if uint8(in2) == 253 {
		return "■"
	}
	return string(in)
}
